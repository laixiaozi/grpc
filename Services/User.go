package Services

import (
	"17jzh.com/user-service/models"
	"17jzh.com/user-service/pbs"
	"context"
	"fmt"
)

type UserServiceServer struct {
	pbs.UnimplementedUserServiceServer
}

//继承服务
func (s *UserServiceServer) SayHello(ctx context.Context, Point *pbs.HelloRequest) (*pbs.HelloResponse, error) {
	message := fmt.Sprintf("这是一个测试 grpc 的方法 %s", Point.Name)
	return &pbs.HelloResponse{Message: message}, nil
}

func (s *UserServiceServer) GetUserById(ctx context.Context, Point *pbs.UserId) (*pbs.UsersMod, error) {
	userId := int64(Point.Id)
	userMod := models.UserModel{}
	userMod.MysqlGetUserById(userId)
	usmod := userMod.ToPb()
	return &usmod, nil
}

//这是一个测试 的rpc接口,
//不用作逻辑业务,
//用来测试各种逻辑和功能,
//可以随意修改
func (s *UserServiceServer) Test(ctx context.Context, Point *pbs.UsersMod) (*pbs.UsersMod, error) {
	userMod := models.UserModel{}
	userMod.PbToMod(*Point)
	userMod.CreateUser()
	usmod := userMod.ToPb()
	return &usmod, nil
}
