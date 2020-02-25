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

/*测试查询*/
func (thisdb *MysqlInterface) Search(sql string, args ...interface{}) error {
	rows, err := thisdb.DB.Query(sql, args[0])
	fmt.Println("搜索用户信息测试" , args[1])
	if err != nil {
		fmt.Println("查询失败", err)
		return err
	}
	defer  rows.Close()
	if err := rows.Err(); err != nil {
		fmt.Printf("获取rows错误", err)
		return err
	}
	for rows.Next() {
		//columns, err := rows.Columns()
		if err != nil {
			fmt.Println("获取字段错误", err)
			return err
		}
		if err := rows.Scan(&args[1]); err != nil {
			fmt.Printf("获取数据错误", err)
		} else {
			fmt.Println("用户名称:", args[1])
		}
		fmt.Println("-----------")
	}
	defer thisdb.DB.Close() //查询完成之后需要关闭数据库
	return nil
}


