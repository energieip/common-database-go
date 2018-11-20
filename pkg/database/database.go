package database

const (
	//RETHINKDB database
	RETHINKDB = "rethinkdb"
)

type databaseError struct {
	s string
}

func (e *databaseError) Error() string {
	return e.s
}

// NewError raise an error
func NewError(text string) error {
	return &databaseError{text}
}

// DatabaseInterface database abstraction layer
type DatabaseInterface interface {
	Initialize(DatabaseConfig) error
	CreateDB(string) error
	CreateTable(string, string) error
	InsertRecord(dbName, tableName string, data map[string]interface{}) (string, error)
	UpdateRecord(string, string, string, map[string]interface{}) error
	GetRecords(dbName, tableName string, criteria map[string]interface{}) ([]interface{}, error)
	GetRecord(dbName, tableName string, criteria map[string]interface{}) (interface{}, error)
	RecordCount(dbName, tableName string) (int, error)
	FetchAllRecords(dbName, tableName string) ([]interface{}, error)
	DeleteRecord(dbName, tableName, id string) error
	ListenTableChange(dbName, tableName string) (*DBCursor, error)
	ListenDBChange(dbName string) (*DBCursor, error)
	ListenFilterTableChange(dbName, tableName string, criteria map[string]interface{}) (*DBCursor, error)
	Close() error
}

//DatabaseConfig configuration structure
type DatabaseConfig struct {
	IP               string
	Port             string
	User             string //for authentification
	Password         string
	ServerCertificat string
	ClientCertificat string
	ClientKey        string
}

// NewNetwork instanciate the appropriate networkinterface
func NewDatabase(protocol string) (DatabaseInterface, error) {
	switch protocol {
	case RETHINKDB:
		return &RethinkbDatabase{}, nil
	default:
		return nil, NewError("Unknow databse type " + protocol)
	}
}
