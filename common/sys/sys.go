package sys

import (
	"fmt"
	"time"
	"casino_common/utils/redis"
	"math/rand"
	"runtime"
	"casino_common/common/log"
	"casino_common/common/db"
	"casino_common/common/consts/tableName"
)

func SysInit(releaseTag int32,
	prodMode bool,
	redisAddr string,
	redisName string,
	logPath string,
	logName string,
	mongoIp string,
	mongoName string,
	mongoSeqTables []string) error {

	var e error
	InitRedis(redisAddr, redisName)                                                    //初始化redis
	initRandSeed()                                                                     //初始化随机数种子
	runtime.GOMAXPROCS(runtime.NumCPU())                                               //初始化cpu数量
	log.InitLogger(logPath, logName)                                                   //初始化日志
	db.InitMongoDb(mongoIp, mongoName, tableName.DB_ENSURECOUNTER_KEY, mongoSeqTables) //初始化mongo 的地址
	e = initSysConfig()                                                                //初始化系统配置
	if e != nil {
		fmt.Printf("加载sys_ocnfig 的时候错误: %v", e)
		return e
	}

	initGAMEENV(prodMode) //初始化环境变量
	return nil
}

//
func InitRedis(addr, name string) error {
	fmt.Println("3，初始化redis...")
	data.InitRedis(addr, name)
	return nil
}

func initRandSeed() {
	fmt.Println("5，初始化随机数种子")
	s := time.Now().UTC().UnixNano()
	rand.Seed(s)
}
