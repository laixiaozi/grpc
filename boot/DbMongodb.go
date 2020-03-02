package boot

import (
	"17jzh.com/user-service/config"
	"17jzh.com/user-service/utility"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var MongoDB MongodbInterface

type MongodbInterface struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

var OpenMongo = func() {
	MongoDB = MongodbInterface{}
	//connectStr := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", config.MongodbConfig["user"], config.MongodbConfig["pwd"] , config.MongodbConfig["host"] ,
	connectStr := fmt.Sprintf("mongodb://%s:%d", config.MongodbConfig["host"], config.MongodbConfig["port"].(int))
	opt := options.Client().ApplyURI(connectStr)
	opt.SetMaxConnIdleTime(10 * time.Second) // 超时时间
	opt.SetMaxPoolSize(500)                  //最大连接数
	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		utility.Abort("链接mongodb失败%+v" + err.Error())
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		utility.Abort(err)
	}
	MongoDB.Client = client
	MongoDB.Collection = MongoDB.Client.Database(config.MongodbConfig["database"].(string)).Collection(config.MongodbConfig["collection"].(string))
}

func (thisDb *MongodbInterface) Start() {
	var once sync.Once
	once.Do(OpenMongo)
}
