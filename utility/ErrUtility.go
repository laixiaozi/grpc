package utility

/*错误系统*/

func CacheErr() {
	defer func() {
		if err := recover(); err != nil {
			Debug("系统异常终端", err)
		}
	}()
}
