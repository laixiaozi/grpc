package models

import (
	"17jzh.com/user-service/boot"
	"17jzh.com/user-service/config"
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
	"context"
	"encoding/json"
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

//查找最大的用户id
func (thisuser *UserModel) GetMaxUserId() int64 {
	var maxUserId int64 = 0
	maxUserId, err := thisuser.RedisGetMaxUserId()
	if err != nil {
		utility.Debug("获取 "+config.REDIS_USER_MAXID+"失败", err.Error())
		maxUserId = thisuser.MongoGetMaxUserId()
	}
	return maxUserId
}

// 添加新用户,返回用户信息
func (thisuser *UserModel) CreateUser() int64 {
	do := func(maxUserId int64) {
		thisuser.LoginAt = utility.NowStr()
		thisuser.CreatedAt = utility.NowStr()
		thisuser.ID = maxUserId
		if err := thisuser.MysqlCreateUser(); err != nil {
			utility.Debug("mysql :创建新用户失败", err)
			return
		}
		if err := thisuser.MongoCreateUser(); err != nil {
			// todo 删除mysql的数据
			thisuser.MysqlDelUser()
			utility.Debug("mongodb 创建用户信息失败", err)
			return
		}
	}
	f := func() int64 {
		//更新最大用户id
		maxUserId, _ := thisuser.RedisMaxUserIdIncr()
		thisuser.ID = maxUserId
		thisuser.CreatedAt = utility.NowStr()
		thisuser.LoginAt = utility.NowStr()
		return maxUserId
	}
	//写入数据
	maxUserId := f()
	do(maxUserId)

	return thisuser.ID
}

// 更新用户信息
func (thisuser *UserModel) UpDateUser() (int64, error) {

	OldUser := &UserModel{ID: thisuser.ID}
	if err := OldUser.MysqlGetUserById(); err != nil {
		utility.Debug(" 获取用户当前信息失败", err)
		return -1, err
	}
	r, e := thisuser.MysqlUpdateUser()
	if e != nil {
		return r, e
	}
	if err := thisuser.MongoUpdateUser(); err != nil {
		// 回退mysql的数据
		r, e = OldUser.MysqlUpdateUser()
	}
	return r, e
}

// 查询，根据姓名
func (thisuser *UserModel) SearchUserByName(userName string) {
	thisuser.MongoSearchUserByName(userName)
}

///////////////////// mysql //////////////////////////////////
/*
 根据表名称获取用户信息
*/
func (thisuser *UserModel) MysqlGetUserById() error {
	tableName := utility.GetTableNameByUserId(thisuser.ID)
	sql := fmt.Sprintf("SELECT id ,member ,realname ,headimg ,headimg2 ,mobile, "+
		"role_id, cid, is_vip, status,edu_type ,edu_year ,exp ,login_at ,device_id, client_type  "+
		" FROM `%s` WHERE id= '%d'", tableName, thisuser.ID) //updated_at , created_at , deleted_at
	conn, err := boot.MysqlDb.DB.Conn(context.TODO())
	if err != nil {
		utility.Debug(err)
		return err
	}
	defer conn.Close()
	r, err := conn.QueryContext(context.Background(), sql)
	if err != nil || r == nil || r.Err() != nil {
		utility.Debug("获取用户信息失败", err)
		return err
	}
	defer r.Close()
	if r.Next() {
		err = r.Scan(&thisuser.ID, &thisuser.Member, &thisuser.Realname, &thisuser.Headimg,
			&thisuser.Headimg2, &thisuser.Mobile, &thisuser.RoleId, &thisuser.Cid, &thisuser.IsVip,
			&thisuser.Status, &thisuser.EduType, &thisuser.EduYear, &thisuser.Exp, &thisuser.LoginAt,
			&thisuser.DeviceId, &thisuser.ClientType) //, &thisuser.UpdatedAt, &thisuser.CreatedAt, &thisuser.DeletedAt
		if err != nil {
			utility.Debug(err)
			return err
		}
	}
	return nil
}

/*新增用户*/
func (thisuser *UserModel) MysqlCreateUser() error {
	table := utility.GetTableNameByUserId(thisuser.ID)
	sql := fmt.Sprintf("INSERT INTO `%s` (id , member ,realname ,headimg ,headimg2 ,mobile, "+
		"role_id, cid, is_vip, status,edu_type ,edu_year ,exp ,login_at ,device_id, client_type ,created_at"+
		")VALUE(? , ? ,? , ? ,? ,? ,? ,? ,? , ? , ? ,? ,? ,? ,?,?,?)", table)
	res, err := boot.MysqlDb.DB.Exec(sql, thisuser.ID, thisuser.Member, thisuser.Realname, thisuser.Headimg, thisuser.Headimg2, thisuser.Mobile,
		thisuser.RoleId, thisuser.Cid, thisuser.IsVip, thisuser.Status, thisuser.EduType, thisuser.EduYear, thisuser.Exp,
		thisuser.LoginAt, thisuser.DeviceId, thisuser.ClientType, thisuser.CreatedAt)
	if err != nil {
		utility.Debug(err.Error())
		return err
	}
	thisuser.ID, _ = res.LastInsertId()
	return nil
}

/*
 删除用户信息
*/
func (thisuser *UserModel) MysqlDelUser() (int64, error) {
	table := utility.GetTableNameByUserId(thisuser.ID)
	sql := fmt.Sprintf("DELETE FROM  `%s` WHERE id  = %d", table, thisuser.ID)
	res, err := boot.MysqlDb.DB.Exec(sql)
	if err != nil {
		utility.Debug(err.Error())
		return -1, err
	}
	return res.RowsAffected()
}

/*更新用户信息*/
func (thisuser *UserModel) MysqlUpdateUser() (int64, error) {
	table := utility.GetTableNameByUserId(thisuser.ID)
	sql := fmt.Sprintf("UPDATE  `%s`"+
		" SET  realname = ? ,"+
		"  headimg = ? ,"+
		"  headimg2 =?  ,"+
		"  mobile = ? , "+
		"  role_id = ?, "+
		"  is_vip=?,"+
		"  status = ?,"+
		"  edu_type = ? ,"+
		"  edu_year = ? ,"+
		"  exp = ? "+
		//"  login_at = ?
		"  WHERE  id  = ?", table)
	res, err := boot.MysqlDb.DB.Exec(sql, thisuser.Realname, thisuser.Headimg, thisuser.Headimg2, thisuser.Mobile,
		thisuser.RoleId, thisuser.IsVip, thisuser.Status, thisuser.EduType, thisuser.EduYear, thisuser.Exp, thisuser.ID) //time.Now(),
	if err != nil {
		utility.Debug(err.Error())
		return -1, err
	}
	return res.RowsAffected()
}

///////////////////// mongodb //////////////////////////////////

/*最大的用户id*/
func (thisuser *UserModel) MongoGetMaxUserId() int64 {
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

func (thisuser *UserModel) MongoSearchUserByName(name string) {
	filter := bson.D{
		{"realname", bson.D{
			{"$regex",  name },
		  },
		},
	}
   var skipNum int64 = 0
   var limitNUm int64 = 5
	opt := options.FindOptions{Skip: &skipNum ,Limit: &limitNUm }
	 cur , err := boot.MongoDB.Collection.Find(context.TODO(), filter , &opt)
	 if err != nil{
	 	utility.Debug("查询用户姓名失败",err)
	 }

	 for cur.Next(context.TODO()){
	 	 item := UserModel{}
	 	 cur.Decode(&item)
	 	utility.DebugInfo(item)
	 }
	utility.Debug(filter)
}

/*新增用户信息*/
func (thisuser *UserModel) MongoCreateUser() error {

	data := bson.M{
		"id":          thisuser.ID,
		"member":      thisuser.Member,
		"realname":    thisuser.Realname,
		"headimg":     thisuser.Headimg,
		"headimg2":    thisuser.Headimg2,
		"mobile":      thisuser.Mobile,
		"role_id":     thisuser.RoleId,
		"cid":         thisuser.Cid,
		"is_vip":      thisuser.IsVip,
		"status":      thisuser.Status,
		"edu_type":    thisuser.EduType,
		"edu_year":    thisuser.EduYear,
		"exp":         thisuser.Exp,
		"login_at":    thisuser.LoginAt,
		"device_id":   thisuser.DeviceId,
		"client_type": thisuser.ClientType,
		"created_at":  thisuser.CreatedAt,
	}
	_, err := boot.MongoDB.Collection.InsertOne(context.Background(), data)
	return err
}

/*新增用户信息*/
func (thisuser *UserModel) MongoUpdateUser() error {
	data := bson.D{
		{
			"$set",
			bson.D{
				{"member", thisuser.Member},
				{"realname", thisuser.Realname},
				{"headimg", thisuser.Headimg},
				{"headimg2", thisuser.Headimg2},
				{"mobile", thisuser.Mobile},
				{"role_id", thisuser.RoleId},
				{"cid", thisuser.Cid},
				{"is_vip", thisuser.IsVip},
				{"status", thisuser.Status},
				{"edu_type", thisuser.EduType},
				{"edu_year", thisuser.EduYear},
				{"exp", thisuser.Exp},
				{"login_at", thisuser.LoginAt},
				{"device_id", thisuser.DeviceId},
				{"client_type", thisuser.ClientType},
				{"created_at", thisuser.CreatedAt},
			},
		},
	}
	where := bson.D{
		{"id", thisuser.ID},
	}
	_, err := boot.MongoDB.Collection.UpdateOne(context.TODO(), where, data)
	return err
}

///////////////////// reids  //////////////////////////////////

func (thisuser *UserModel) RedisGetMaxUserId() (int64, error) {
	MaxId, err := boot.RedisDb.Client.Get(config.REDIS_USER_MAXID).Int64()
	return MaxId, err
}

func (thisuser *UserModel) RedisSetMaxUserId(maxUserId int64) error {
	string, err := boot.RedisDb.Client.Set(config.REDIS_USER_MAXID, maxUserId, -1).Result()
	utility.Debug(string)
	return err
}

func (thisuser *UserModel) RedisMaxUserIdIncr() (int64, error) {
	return boot.RedisDb.Client.Incr(config.REDIS_USER_MAXID).Result()
}

func (thisuser *UserModel) RedisLock() {
	ok, err := boot.RedisDb.Client.Set(config.REDIS_LOCK_STATUS, config.REDIS_LOCK_VALUE, config.REDIS_EXP_FOREVER).Result()
	utility.Debug("已加锁:", ok, "err:", err)
}

func (thisuser *UserModel) RedisUnLock() {
	ok, err := boot.RedisDb.Client.Set(config.REDIS_LOCK_STATUS, config.REDIS_UNLOCK_VALUE, config.REDIS_EXP_FOREVER).Result()
	utility.Debug("已解锁:", ok, "err:", err)
}

func (thisuser *UserModel) RedisIsLock() bool {
	isLock, err := boot.RedisDb.Client.Get(config.REDIS_LOCK_STATUS).Int64()
	if err != nil {
		utility.Debug("获取redis锁失败:", err)
		ok, err := boot.RedisDb.Client.Set(config.REDIS_LOCK_STATUS, config.REDIS_UNLOCK_VALUE, config.REDIS_EXP_FOREVER).Result()
		utility.Debug("新增锁定状态:", ok, "err:", err)
	}
	return isLock == int64(config.REDIS_LOCK_VALUE) //1 锁定状态 0 未锁定状态
}

/*加入用户注册信息到队列里面*/
func (thisuser *UserModel) RedisJoinList() bool {
	d, err := json.Marshal(thisuser)
	if err != nil {
		utility.Debug("用户加入队列失败", err)
		return false
	}
	n, err := boot.RedisDb.Client.LPush(config.REDIS_USER_LIST, d).Result()
	if err != nil {
		utility.Debug("用户加入队列失败", err)
		return false
	}
	utility.Debug(n)
	return false
}

/*消费到队列里面*/
func (thisuser *UserModel) RedisConsumList() {

	n, err := boot.RedisDb.Client.LLen(config.REDIS_USER_LIST).Result()
	if err != nil {
		utility.Debug("获取用户队列失败", err)
		return
	}
	utility.Debug("当前队列中的用户数量:", n)
	for i := 0; int64(i) < n; i++ {
		userdata := boot.RedisDb.Client.RPop(config.REDIS_USER_LIST).String()
		utility.Debug(userdata)
		umod := UserModel{}
		err := json.Unmarshal([]byte(userdata), umod)
		if err != nil {
			continue
		}
		//todo 写入数据
		utility.Debug("当前消费用户的id:", umod.ID)

	}
}

///////////////////// //////////////////////////////////

//生成grpc需要的数据类型
/*学校,昵称,等属性需要单独取*/
func (thisuser *UserModel) ToPb() pbs.UsersMod {

	pbuserMod := pbs.UsersMod{}
	pbuserMod.Id = thisuser.ID
	pbuserMod.Member = thisuser.Member
	pbuserMod.Realname = thisuser.Realname
	pbuserMod.Headimg = thisuser.Headimg
	pbuserMod.Headimg2 = thisuser.Headimg2
	pbuserMod.Mobile = thisuser.Mobile
	pbuserMod.RoleId = int32(thisuser.RoleId)
	pbuserMod.Cid = int32(thisuser.Cid)
	pbuserMod.IsVip = int32(thisuser.IsVip)
	pbuserMod.Status = int32(thisuser.Status)
	pbuserMod.EduType = int32(thisuser.EduType)
	pbuserMod.EduYear = int32(thisuser.EduYear)
	pbuserMod.Exp = int32(thisuser.Exp)
	return pbuserMod
}

/*学校,昵称,等属性需要单独取*/
func (thisuser *UserModel) PbToMod(pbuserMod pbs.UsersMod) {

	thisuser.ID = pbuserMod.Id
	thisuser.Member = pbuserMod.Member
	thisuser.Realname = pbuserMod.Realname
	thisuser.Headimg = pbuserMod.Headimg
	thisuser.Headimg2 = pbuserMod.Headimg2
	thisuser.Mobile = pbuserMod.Mobile
	thisuser.RoleId = int(pbuserMod.RoleId)
	thisuser.Cid = int(pbuserMod.Cid)
	thisuser.IsVip = int(pbuserMod.IsVip)
	thisuser.Status = int(pbuserMod.Status)
	thisuser.EduType = int(pbuserMod.EduType)
	thisuser.EduYear = int(pbuserMod.EduYear)
	thisuser.Exp = int(pbuserMod.Exp)

}
