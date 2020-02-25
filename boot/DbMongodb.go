package boot

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var MongoDB *MongodbInterface

type MongodbInterface struct {
	Client *mongo.Client
}

var OpenMongo = func() {
	MongoDB = &MongodbInterface{}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("链接mongodb失败%+v", err)
	}
	MongoDB.Client = client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("链接 mongodb服务器错误", err)
		return
	}
}

func (thisDb *MongodbInterface) Start() {
	var once sync.Once
	once.Do(OpenMongo)
	////集合
	//collection := client.Database("nxdev").Collection("users")
	//cursor:= collection.FindOne(context.Background(),bson.M{"id":3751926})
	//type u struct {
	//	Id int `json:"id"`
	//	RealName string `json:"realname"`
	//	NickName string `json:"nickname"`
	//}
	//user := u{}
	//cursor.Decode(&user)
	//fmt.Println(user)
}
