package log

import "os"

func InitLogger(logPath, logName string, logFileSize int32) {
	//配置log 文件.这里可以通过conf.json 来进行配置
	SetMaxFileCount(int32(100))
	if logFileSize <= 0 {
		SetMaxFileSize(int64(100), MB)
	}else {
		SetMaxFileSize(int64(logFileSize), MB)
	}
	SetLevel(ALL)
	SetConsole(true)
	SetDebug(true)
	SetMaxLogSeq(int64(500))
	if logPath == "" {
		logPath = os.Getenv("GOPATH") + "/log"
	}
	InitLoggers(logPath, logName)
}
