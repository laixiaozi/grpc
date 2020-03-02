package boot

import (
	"17jzh.com/user-service/config"
	"17jzh.com/user-service/utility"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var MysqlDb MysqlInterface

/*mysql数据库链接*/
type MysqlInterface struct {
	DB      *sql.DB
	OpenErr error
	Ctx     context.Context
	Conn    *sql.Conn
}

/*初始化mysql链接数据库*/
var OpenMysql = func() {
	MysqlDb = MysqlInterface{}
	constr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.MysqlConfig["user"], config.MysqlConfig["pwd"], config.MysqlConfig["host"],
		config.MysqlConfig["port"].(int), config.MysqlConfig["database"], config.MysqlConfig["option"])
	db, err := sql.Open("mysql", constr)
	if err != nil {
		MysqlDb.OpenErr = err
		utility.Abort("连接到mysql 失败:" + err.Error())
	}
	//defer db.Close()  //db不应被关闭,会阻止新的连接和新的回话查询
	MysqlDb.DB = db
	MysqlDb.DB.SetMaxOpenConns(config.MysqlConfig["maxConn"].(int)) //连接池中的最大打开连接数
	MysqlDb.DB.SetMaxIdleConns(config.MysqlConfig["maxOpen"].(int)) //连接池中最大连接数
	if err := MysqlDb.DB.Ping(); err != nil {
		utility.Abort("连接数据库失败", err)
	}

}

func (thisdb *MysqlInterface) Start() {
	var once sync.Once
	once.Do(OpenMysql)
}

/*
  取一个新的连接.需要手动关闭才能释放连接到连接池
  用法
   conn , err := boot.MysqlDb.DB.Conn() 或者
   conn , err := boot.MysqlDb.DB.NewConn()
    if err != nil{
      utility.Debug(err)
      return
    }
   conn.QueryContext(context.TODO() ,sql)
  defer conn.Close() //释放连接
*/
func (thisdb *MysqlInterface) NewConn(ctx context.Context) (*sql.Conn, error) {
	Conn, err := thisdb.DB.Conn(ctx)
	return Conn, err
}
