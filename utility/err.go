package utility

/*错误系统*/

func Err(){
	defer func() {
		if err := recover(); err != nil {
			Debug("系统异常终端", err)
		}
	}()
}
