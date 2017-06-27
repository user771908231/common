package log

import "os"

func InitLogger(logPath, logName string) {
	//配置log 文件.这里可以通过conf.json 来进行配置
	SetMaxFileCount(int32(100))
	SetMaxFileSize(int64(100), MB)
	SetLevel(ALL)
	SetConsole(true)
	SetDebug(true)
	SetMaxLogSeq(int64(500))
	if logPath == "" {
		logPath = os.Getenv("GOPATH") + "/log"
	}
	InitLoggers(logPath, logName)
}
