package config

var MysqlConfig map[string]interface{} = map[string]interface{}{
	"host":     "localhost",
	"user":     "root",
	"pwd":      "123456",
	"port":     3306,
	"database": "jzh-dev",
	"option":   "charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true",
	"maxOpen":  500, //最大连接数
	"maxConn":  500, // 连接池最大打开连接数
}
