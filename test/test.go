package test

import (
	"casino_common/common/sys"
	"casino_common/utils/db"
	"casino_common/common/log"
)

func TestInit() {
	//sys.InitRedis("192.168.199.200:6379", "test")
	//db.Oninit("192.168.199.200", 27017, "test", "id")
	//log.InitLogger("", "") //初始化日志处理
	sys.InitRedis("127.0.0.1:6379", "test")
	db.Oninit("127.0.0.1:27017", "127.0.0.1:27017", "test", "id")
	log.InitLogger("", "") //初始化日志处理
}
