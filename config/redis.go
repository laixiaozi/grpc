package config

import "time"

const (
	REDIS_USER_MAXID   string        = "user:max:id"
	REDIS_LOCK_STATUS  string        = "user:max:id:lock:status" // 0: 未锁定, 1 :锁定
	REDIS_USER_LIST    string        = "users:create:list:"      // 0: 未锁定, 1 :锁定
	REDIS_DB_USER      int           = 1
	REDIS_POOL_SIZE    int           = 10 //默认每个cup的链接数量
	REDIS_LOCK_VALUE   int           = 1
	REDIS_UNLOCK_VALUE int           = 0
	REDIS_EXP_FOREVER  time.Duration = -1
)

var RedisConfig map[string]interface{} = map[string]interface{}{
	"host":     "localhost",
	"user":     "",
	"pwd":      "",
	"port":     6379,
	"database": REDIS_DB_USER,
}
