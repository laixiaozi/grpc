package utility

import "fmt"

/**
根绝 id 获取用户表名称
*/
const (
	tableNum    int    = 16      //分表个数
	tablePrefix string = "users_" //表前缀
)

func GetTableNameByUserId(user_id int64) string {
	remainder := user_id % int64(tableNum)
	return fmt.Sprintf("%s%d", tablePrefix, remainder)
}
