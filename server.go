package main

import (
	"17jzh.com/user-service/Services"
	"17jzh.com/user-service/boot"
	"17jzh.com/user-service/config"
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
	"google.golang.org/grpc"
	"net"
)

func init() {
	boot.MysqlDb.Start()
	boot.MongoDB.Start()
	boot.RedisDb.Start()
	utility.Debug("启动grpc服务...")
}
func main() {
	lis, err := net.Listen("tcp", config.SERVER_PORT)
	if err != nil {
		utility.Abort("无法启动服务监听", err)
	}
	s := grpc.NewServer()
	//注册服务
	pbs.RegisterUserServiceServer(s, &Services.UserServiceServer{})
	if err := s.Serve(lis); err != nil {
		utility.Abort("启动Rpc服务失败", err)
	}
	utility.Debug("启动user grpc服务..")
}
