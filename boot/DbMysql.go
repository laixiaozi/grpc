package boot

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)


var MysqlDb *MysqlInterface

/*mysql数据库链接*/
type MysqlInterface struct {
	DB      *sql.DB
	OpenErr error
	Ctx     context.Context
}

const DB_URL string = "root:123456@tcp(localhost:3306)/jzh-dev?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"

/*初始化mysql链接数据库*/
var OpenMysql= func() {
	MysqlDb = &MysqlInterface{}
	db, err := sql.Open("mysql", DB_URL)
	if err != nil {
		fmt.Println("打开数据库错误", err)
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


