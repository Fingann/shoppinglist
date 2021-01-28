package database

//Error contains different error codes
type Error string

const (
	//ErrUnknown error, something unexpected has happened
	ErrUnknown Error = "An unknown error has occured"
	//ErrNotFound indicates item was not found in reposetory
	ErrNotFound Error = "The item was not found in the database"
	//ErrOperationFailed indicates that the performed action failed
	ErrOperationFailed Error = "The opperation was not successfull"
)

func (e Error) Error() string {
	return string(e)
}

//Reader will perform read actions against the database
type Reader interface {
	Find(id string) (interface{}, error)
	FindAll() (interface{}, error)
}

//Writer will perform write actions against the database
type Writer interface {
	Update(entity *interface{}) (interface{}, error)
	Create(entity *interface{}) (string, error)
	Delete(id string) (interface{}, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
