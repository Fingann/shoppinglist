package database

import (
	"fmt"
	"os"
	"shoppinglist/models"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/estransport"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

// SearchResults wraps the Elasticsearch search response.
//
type SearchResults struct {
	Total int            `json:"total"`
	Hits  []*models.Item `json:"hits"`
}

//ElasticDB represents an Elasticsearch db connection
type ElasticDB struct {
	context   *elasticsearch.Client
	indexName string
}

//NewElasticDB returns an ElasticDB context
func NewElasticDB(indexName string) *ElasticDB {
	cfg := elasticsearch.Config{
		Logger: &estransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	eDB := ElasticDB{
		context:   es,
		indexName: indexName,
	}

	//create Index
	err = createIndex(es, indexName)
	if err != nil {
		panic(err)
	}

	return &eDB
}

func createIndex(esClient *elasticsearch.Client, name string) error {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := esClient.Indices.Exists([]string{name})
	if err != nil {
		// Handle error
		panic(err)
	}
	//if index does not exist, create a new one with the specified mapping
	if exists.StatusCode == 404 {
		res, err := esClient.Indices.Create(name)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		fmt.Println(res)
	}

	return nil
}

//Index stores an document in the index.
func (db ElasticDB) Index(entity interface{}) (interface{}, error) {
	res, err := db.context.Index(db.indexName, esutil.NewJSONReader(&entity))
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Exists returns true when a document with id already exists in the store.
//
func (db *ElasticDB) Exists(id interface{}) (bool, error) {
	res, err := db.context.Exists(db.indexName, fmt.Sprintf("%v", id))
	if err != nil {
		return false, err
	}
	switch res.StatusCode {
	case 200:
		return true, nil
	case 404:
		return false, nil
	default:
		return false, fmt.Errorf("[%s]", res.Status())
	}
}

//Search after Something
func (db *ElasticDB) Search(query string) (interface{}, error) {
	return nil, nil
}

/*
// Search returns results matching a query, paginated by after.
func (db *ElasticDB) Search(query string, after ...string) (*SearchResults, error) {
	var results SearchResults

	res, err := db.context.API.Search(
		db.context.Search.WithIndex(s.indexName),
		db.context.Search.WithBody(es.buildQuery(query, after...)),
	)
	if err != nil {
		return &results, err
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return &results, err
		}
		return &results, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}

	type envelopeResponse struct {
		Took int
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				ID         string          `json:"_id"`
				Source     json.RawMessage `json:"_source"`
				Highlights json.RawMessage `json:"highlight"`
				Sort       []interface{}   `json:"sort"`
			}
		}
	}

	var r envelopeResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return &results, err
	}

	results.Total = r.Hits.Total.Value

	if len(r.Hits.Hits) < 1 {
		results.Hits = []*Hit{}
		return &results, nil
	}

	for _, hit := range r.Hits.Hits {
		var h Hit
		h.ID = hit.ID
		h.Sort = hit.Sort
		h.URL = strings.Join([]string{baseURL, h.ID, ""}, "/")

		if err := json.Unmarshal(hit.Source, &h); err != nil {
			return &results, err
		}

		if len(hit.Highlights) > 0 {
			if err := json.Unmarshal(hit.Highlights, &h.Highlights); err != nil {
				return &results, err
			}
		}

		results.Hits = append(results.Hits, &h)
	}

	return &results, nil
}
*/
