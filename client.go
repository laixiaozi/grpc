package main

import (
	"17jzh.com/user-service/pbs"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	fmt.Println("grpc 客户端连接")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println("链接错误", err)
		return
	}
	defer conn.Close()
	c := pbs.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if r, err := c.GetUserById(ctx, &pbs.UserId{Id: 3751926}); err != nil {
		fmt.Println("gprc 请求失败")
	} else {
		fmt.Println(r)
	}
}
