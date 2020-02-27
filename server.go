package main

import (
	"17jzh.com/user-service/Services"
	"17jzh.com/user-service/boot"
	"17jzh.com/user-service/pbs"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	SERVER_PORT = ":50051"
)

func init() {
	log.Println("初始化grpc服务")
	boot.MysqlDb.Start()
	boot.MongoDB.Start()
}
func main() {
	lis, err := net.Listen("tcp", SERVER_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//注册服务
	pbs.RegisterUserServiceServer(s, &Services.UserServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
