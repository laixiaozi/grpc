package config

const (
	REDIS_USER_MAXID string = "user:max:id"
	REDIS_DB_USER    int    = 1
	REDIS_POOL_SIZE  int    = 10 //默认每个cup的链接数量
)

var RedisConfig map[string]interface{} = map[string]interface{}{
	"host":     "localhost",
	"user":     "",
	"pwd":      "",
	"port":     6379,
	"database": REDIS_DB_USER,
}
