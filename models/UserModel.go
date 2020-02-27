package models

import (
	"17jzh.com/user-service/boot"
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/*用户模型*/
type UserModel struct {
	ID         int64  `gorm:"primary_key" json:"id"`
	Member     int64  `json:"member"`   //会员号
	Realname   string `json:"realname"` //姓名
	Headimg    string `json:"headimg"`  //用户头像
	Headimg2   string `json:"headimg2"` //用户头像
	Mobile     string `json:"mobile"`   //手机号
	RoleId     int    `json:"role_id"`  //家长角色:0-家长,1-爸爸,2-妈妈,3-老师
	Cid        int    `json:"cid"`      //首次注册渠道
	IsVip      int    `json:"is_vip"`   //是否VIP:0-非Vip,1-VIP,2-赠送
	Status     int    `json:"status"`   //用户状态:1-正常,2-封号
	EduType    int    `json:"edu_type"` //学龄段
	EduYear    int    `json:"edu_year"` //入学年份
	Exp        int    `json:"exp"`      //经验值
	LoginAt    string `json:"login_at"`
	DeviceId   string `json:"device_id"` //app 设备id
	ClientType int    `json:"client_type"`
	UpdatedAt  string `json:"updated_at"`
	CreatedAt  string `json:"created_at"`
	DeletedAt  string `json:"deleted_at"`
}

///////////////////// mysql //////////////////////////////////
/*
 根据表名称获取用户信息
*/
func (thisuser *UserModel) GetUserById(tablename string, userId int64) {
	sql := fmt.Sprintf("SELECT id ,member ,realname ,headimg ,headimg2 ,mobile, "+
		"role_id, cid, is_vip, status,edu_type ,edu_year ,exp ,login_at ,device_id, client_type  "+
		" FROM `%s` WHERE id= '%d'", tablename, userId) //updated_at , created_at , deleted_at
	r, err := boot.MysqlDb.DB.Query(sql)
	if err != nil || r == nil || r.Err() != nil {
		utility.Debug("获取用户信息失败", err)
		return
	}
	defer r.Close()
	if r.Next() {
		err = r.Scan(&thisuser.ID, &thisuser.Member, &thisuser.Realname, &thisuser.Headimg,
			&thisuser.Headimg2, &thisuser.Mobile, &thisuser.RoleId, &thisuser.Cid, &thisuser.IsVip,
			&thisuser.Status, &thisuser.EduType, &thisuser.EduYear, &thisuser.Exp, &thisuser.LoginAt,
			&thisuser.DeviceId, &thisuser.ClientType) //, &thisuser.UpdatedAt, &thisuser.CreatedAt, &thisuser.DeletedAt
		if err != nil {
			utility.Debug(err)
			return
		}
	}
}

/*新增用户*/
func (thisuser *UserModel) CreateUser(table string) error {
	sql := fmt.Sprintf("INSERT INTO `%s` (member ,realname ,headimg ,headimg2 ,mobile, "+
		"role_id, cid, is_vip, status,edu_type ,edu_year ,exp ,login_at ,device_id, client_type ,created_at"+
		")VALUE(? , ? , ? ,? ,? ,? ,? ,? , ? , ? ,? ,? ,? ,?,?,?)", table)
	res, err := boot.MysqlDb.DB.Exec(sql, thisuser.Member, thisuser.Realname, thisuser.Headimg, thisuser.Headimg2, thisuser.Mobile,
		thisuser.RoleId, thisuser.Cid, thisuser.IsVip, thisuser.Status, thisuser.EduType, thisuser.EduYear, thisuser.Exp,
		thisuser.LoginAt, thisuser.DeviceId, thisuser.ClientType, thisuser.CreatedAt)
	if err != nil {
		utility.Debug(err.Error())
		return err
	}
	thisuser.ID, _ = res.LastInsertId()
	return nil
}

///////////////////// mongodb //////////////////////////////////

/*最大的用户id*/
func (thisuser *UserModel) GetMaxUserIdByMongo() int64 {
	var maxid int64 = 0
	groupStages := bson.D{
		{"$group", bson.D{
			{"_id", "null"},
			{"max", bson.D{{"$max", "$id"}}},
		}},
	}
	opt := options.Aggregate().SetMaxTime(5 * time.Second)
	cur, err := boot.MongoDB.Collection.Aggregate(context.TODO(), mongo.Pipeline{groupStages}, opt)
	if err != nil {
		utility.Debug(err)
		return maxid
	}
	var result []bson.M
	if err := cur.All(context.TODO(), &result); err != nil {
		utility.Debug(err)
		return maxid
	}

	if len(result) > 0 {
		maxid = utility.GetInt64(result[0]["max"])
	}
	return maxid
}

///////////////////// reids  //////////////////////////////////

func (thisuser *UserModel) GetMaxUserIdByRedis() int64 {
	var maxid int64 = 0

	return maxid
}

///////////////////// //////////////////////////////////

/*新增用户信息*/
func (thisuser *UserModel) MongoCreateUser() error {

	data := bson.M{
		"id":         123456789,
		"created_at": time.Now().Format("2006-01-02"),
		"realname":   "测试写入数据啊哈哈哈",
	}
	_, err := boot.MongoDB.Collection.InsertOne(context.Background(), data)
	return err
}

//生成grpc需要的数据类型
func (thisuser *UserModel) ToPb() pbs.UsersMod {

	pbuserMod := pbs.UsersMod{}
	pbuserMod.Id = thisuser.ID
	pbuserMod.Member = thisuser.Member
	pbuserMod.RoleId = int32(thisuser.RoleId)
	pbuserMod.IsVip = int32(thisuser.IsVip)
	pbuserMod.Status = int32(thisuser.Status)
	pbuserMod.EduType = int32(thisuser.EduType)
	pbuserMod.EduYear = int32(thisuser.EduYear)
	pbuserMod.Exp = int32(thisuser.Exp)
	pbuserMod.Cid = int32(thisuser.Cid)
	pbuserMod.Realname = thisuser.Realname
	pbuserMod.Mobile = thisuser.Mobile
	pbuserMod.Headimg = thisuser.Headimg
	return pbuserMod
}
