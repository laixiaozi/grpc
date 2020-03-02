package main

import (
	"17jzh.com/user-service/config"
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
	"context"
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
	t := 1 //并发数量
	args := os.Args
	if len(args) > 1 {
		t, _ = strconv.Atoi(args[1])
	}
	if t > config.MysqlConfig["maxOpen"].(int) {
		t = config.MysqlConfig["maxOpen"].(int)
	}

	f := func(uid int32, i int) {
		RpcClient()
		ch <- i
	}

	for i := 0; i < t; i++ {
		var uid int = i
		go f(int32(uid), i)
	}

	for i := 0; i < t; i++ {
		<-ch
		//utility.Debug("ch : ", <-ch)  //显示执行的顺序
	}
	utility.Debug("the end..")
}




//测试接口
func RpcClient() {
	conn, _ := GetRpcConn()
	defer conn.Close()
	c := pbs.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//var usmod *pbs.UsersMod
	//
	//if r, err := c.GetUserById(ctx, &pbs.UserId{Id: 7494755}); err != nil {
	//	utility.Debug("gprc 请求失败" , err)
	//} else {
	//	usmod = r
	//}

	uname := &pbs.UserName{Name:"张"}
	if r, err := c.SearchUserByName(ctx, uname); err != nil {
		utility.Debug("gprc 请求失败" , err)
	} else {
		utility.Debug(r)
	}
}


func CreateUser() {
	conn, _ := GetRpcConn()
	defer conn.Close()
	c := pbs.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	usmod := GetUserMod(1)
	if r, err := c.CreateUser(ctx, &usmod); err != nil {
		utility.Debug("gprc 请求失败" , err)
	} else {
		utility.Debug(r.Id)
	}
}

//获取用户id
func GetUserById(uid int64) {
	conn, _ := GetRpcConn()
	defer conn.Close()
	c := pbs.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	if r, err := c.GetUserById(ctx, &pbs.UserId{Id: uid}); err != nil {
		utility.Debug("gprc 请求失败" , err)
	} else {
		utility.Debug(r.Realname, r.Id)
	}
}

// 获取 rpc请求 句柄
func GetRpcConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		utility.Debug("链接错误", err)
		return nil, err
	}
	return conn, nil
}

func GetUserMod(id int64) pbs.UsersMod {
	//随机数
	return pbs.UsersMod{
		Id:         id,
		Member:      utility.GetRandNum(100000, 900000),
		Realname:    "测试真是姓名",
		Headimg:     "",
		Headimg2:    "",
		Mobile:      "13341342324",
		RoleId:      1,
		Cid:         1,
		IsVip:       0,
		Status:      1,
		EduType:     2,
		EduYear:     2020,
		Exp:         211,
		LoginAt:     time.Now().Format("2006-01-02 15:04:-5"),
		DeviceId:    "test",
		ClientType:  9,
		UpdatedAt:   "",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:-5"),
		DeletedAt:   "",
		SchoolId:    1,
		ClassroomId: 1,
		PlateformId: 1,
		IsNotify:    1,
		Openid:      "",
		Nickname:    "",
	}
}
