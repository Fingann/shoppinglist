package database

//ElasticReader will perform read actions against the database
type ElasticReader interface {
	Search(id string) (interface{}, error)
	Exists(id interface{}) (bool, error)
}

//ElasticWriter will perform write actions against the database
type ElasticWriter interface {
	Index(entity interface{}) (interface{}, error)
}

//ElasticRepository repository interface
type ElasticRepository interface {
	ElasticReader
	ElasticWriter
}
