package boot

import (
	"17jzh.com/user-service/config"
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
}


/*初始化mysql链接数据库*/
var OpenMysql = func() {
	MysqlDb = MysqlInterface{}
	constr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.MysqlConfig["user"], config.MysqlConfig["pwd"], config.MysqlConfig["host"],
		config.MysqlConfig["port"].(int), config.MysqlConfig["database"], config.MysqlConfig["option"])
	db, err := sql.Open("mysql", constr)
	fmt.Println(constr)
	if err != nil {
		MysqlDb.OpenErr = err
	}
	//defer db.Close()
	MysqlDb.DB = db
	MysqlDb.DB.SetMaxOpenConns(500) //连接池中的最大打开连接数
	MysqlDb.DB.SetMaxIdleConns(500) //连接池中最大连接数
}

func (thisdb *MysqlInterface) Start() error {
	var once sync.Once
	once.Do(OpenMysql)
	return nil
}
