package boot

import (
	"17jzh.com/user-service/config"
	"17jzh.com/user-service/utility"
	"fmt"
	"sync"
)
import "github.com/go-redis/redis"

var RedisDb DbRedisInterface

type DbRedisInterface struct {
	Client *redis.Client
	Opt    redis.Options
}

var OpenRedis = func() {
	RedisDb = DbRedisInterface{}
	RedisDb.Opt.Addr = fmt.Sprintf("%s:%d", config.RedisConfig["host"].(string), config.RedisConfig["port"].(int))
	RedisDb.Opt.Password = config.RedisConfig["pwd"].(string)
	RedisDb.Opt.DB = config.RedisConfig["database"].(int)
	RedisDb.Opt.PoolSize = config.REDIS_POOL_SIZE //默认每个cup核心10个连接数 runtime.NumCPU
	RedisDb.Client = redis.NewClient(&RedisDb.Opt)
	if pong, err := RedisDb.Client.Ping().Result(); err != nil {
		utility.Abort("连接到redis失败:", err.Error(), pong)
	}
}

func (thisdb *DbRedisInterface) Start() {
	var one sync.Once
	one.Do(OpenRedis)
}
