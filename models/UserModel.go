package models

import (
	"17jzh.com/user-service/boot"
	"17jzh.com/user-service/pbs"
	"17jzh.com/user-service/utility"
	"fmt"
)

/*用户模型*/
type UserModel struct {
	ID         int    `gorm:"primary_key" json:"id"`
	Member     string `json:"member"`   //会员号
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

/*
 根据表名称获取用户信息
*/
func (thisuser *UserModel) GetUserById(tablename string, userId int64) {
	utility.Err()
	sql := fmt.Sprintf("SELECT id ,member ,realname ,headimg ,headimg2 ,mobile, "+
		"role_id, cid, is_vip, status,edu_type ,edu_year ,exp ,login_at ,device_id, client_type  "+
		" FROM `%s` WHERE id= '%d'", tablename, userId) //updated_at , created_at , deleted_at
	r, err := boot.MysqlDb.DB.Query(sql)
	if err != nil || r == nil || r.Err() != nil {
		utility.Debug("获取用户信息失败", err)
		return
	}
	defer boot.MysqlDb.DB.Close()
	defer  r.Close()
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
func (thisuser *UserModel)CreateUser()error{




}


//生成grpc需要的数据类型
func (thisuser *UserModel) ToPb() pbs.UsersMod {

	pbuserMod := pbs.UsersMod{}
	pbuserMod.Id = int64(thisuser.ID)
	//pbuserMod.SchoolId = int32(thisuser.SchoolId)
	//pbuserMod.ClassroomId = int32(thisuser.ClassroomId)
	//pbuserMod.PlateformId = int32(thisuser.PlateformId)
	//pbuserMod.IsNotify = int32(thisuser.IsNotify)
	//pbuserMod.Openid = thisuser.Openid
	//pbuserMod.Nickname = thisuser.Nickname
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
