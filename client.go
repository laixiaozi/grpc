package main

import (
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	ch := make(chan int)
	t := 500   //创建几个并发
	f := func(uid int32 ,  i int ) {
		RpcClient(uid)
		ch <- i
	}

	for i := 0; i < t; i++ {
		go f(3751926 , i)
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
		utility.Debug(r.Realname)
	}
}