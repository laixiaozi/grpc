package utility

import (
	"log"
)

/*调试日志*/
func Debug(argv ...interface{}) {
	//timeStr := time.Now().Format("2006/01/02 15/04/05")
	log.Println(argv)
}

func DebugInfo(argv ...interface{}) {
	log.Printf("%+v", argv)
}


func Abort(arg ...interface{}) {
	log.Fatal(arg)
}
