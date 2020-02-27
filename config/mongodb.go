package config

const (
	MONGODB_DBHOST             string = "localhost"
	MONGODB_DBUSER             string = "nxdev"
	MONGODB_DBPWD              string = "123456"
	MONGODB_DATABASE           string = "nxdev"
	MONGODB_DBCOLLECTION_USERS string = "users"
	MONGODB_DBPORT             int    = 27017
)

var MongodbConfig map[string]interface{} = map[string]interface{}{
	"host":       "localhost",
	"user":       "nxdev",
	"pwd":        "123456",
	"port":       27017,
	"database":   "nxdev",
	"collection": "users",
}
