package Services

import (
	"17jzh.com/user-service/models"
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
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
	tableName := utility.GetTableNameByUserId(userId)
	userMod := models.UserModel{}
	userMod.GetUserById(tableName, userId)
	usmod := userMod.ToPb()
	//以下为测试方法,不能在本业务中使用
	//userMod.MongoCreateUser()
	//userMod.GetMaxUserId()
	//以上为测试代码,不能在本业务中使用
	return &usmod, nil
}
