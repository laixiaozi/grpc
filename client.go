package main

import (
	"17jzh.com/user-service/config"
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"time"
)
/*
 go run client.go   100  //100个并发
*/
func main() {
	ch := make(chan int)
	t := 10   //并发数量
	args := os.Args
	if  len(args) > 1{
		t , _ = strconv.Atoi(args[1])
	}
	if t > config.MysqlConfig["maxOpen"].(int){
		t = config.MysqlConfig["maxOpen"].(int)
	}

	f := func(uid int32 ,  i int ) {
		RpcClient(uid)
		ch <- i
	}

	for i := 0; i < t; i++ {
		var uid int = 3751920 + i
		go f(int32(uid) , i)
	}

	for i := 0; i < t; i++ {
		<- ch
		//utility.Debug("ch : ", <-ch)
	}
	utility.Debug("the end..")
}


func RpcClient(uid int32){
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("链接错误", err)
		return
	}
	defer conn.Close()
	c := pbs.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if r, err := c.GetUserById(ctx, &pbs.UserId{Id: uid}); err != nil {
		utility.Debug("gprc 请求失败")
	} else {
		utility.Debug(r.Realname , r.Id)
	}
}