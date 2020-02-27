package config

const (
	REDIS_USER_MAXID string = "user:max:id"
	REDIS_DB_USER    int    = 1
)

var RedisConfig map[string]interface{} = map[string]interface{}{
	"host": "localohst",
}
