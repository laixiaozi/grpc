package boot

import (
	"17jzh.com/user-service/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var MongoDB *MongodbInterface

type MongodbInterface struct {
	Client *mongo.Client
	Collection   *mongo.Collection
}

var OpenMongo = func() {
	MongoDB = &MongodbInterface{}
	connectStr := fmt.Sprintf("mongodb://%s:%d",  config.MONGODB_DBHOST , config.MONGODB_DBPORT )
	opt :=options.Client().ApplyURI(connectStr)
	opt.SetMaxConnIdleTime( 10 *time.Second)  // 超时时间
	opt.SetMaxPoolSize( 500 )   //最大连接数
	client, err := mongo.Connect(context.TODO() , opt)
	if err != nil {
		log.Fatalf("链接mongodb失败%+v", err)
	}
     if err :=	client.Ping(context.TODO() , nil); err != nil{
     	log.Fatal(err)
	 }
	MongoDB.Client = client
	MongoDB.Collection = MongoDB.Client.Database(config.MONGODB_DATABASE).Collection(config.MONGODB_DBCOLLECTION_USERS)
}

func (thisDb *MongodbInterface) Start() {
	var once sync.Once
	once.Do(OpenMongo)
}
