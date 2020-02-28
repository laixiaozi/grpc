package utility

import (
	"17jzh.com/user-service/config"
	"time"
)

func NowStr() string {
	return time.Now().Format(config.TIME_LAYOUT_DEFAULT)
}
