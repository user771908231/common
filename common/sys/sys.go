package sys

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/db"
	"casino_common/common/log"
	"casino_common/utils/redis"
	"fmt"
	"math/rand"
	"runtime"
	"time"
	"gopkg.in/mgo.v2"
)

type MgoLoger struct {

}

func (mlog MgoLoger) Output(calldepth int, s string) error {
	if calldepth == 3 {
		log.T("[MongoDB]:%s", s)
	}
	return nil
}


func SysInit(releaseTag int32,
	prodMode bool,
	redisAddr string,
	redisName string,
	logPath string,
	logName string,
	logFileSize int32,
	logFileCount int32,
	mongoIp string,
	mongoLogIp string,
	mongoName string,
	mongoSeqTables []string) error {

	var e error
	InitRedis(redisAddr, redisName)                                                    //初始化redis
	initRandSeed()                                                                     //初始化随机数种子
	runtime.GOMAXPROCS(runtime.NumCPU())                                               //初始化cpu数量
	log.InitLogger(logPath, logName, logFileSize, logFileCount)                                                   //初始化日志
	db.InitMongoDb(mongoIp, mongoLogIp, mongoName, tableName.DB_ENSURECOUNTER_KEY, mongoSeqTables) //初始化mongo 的地址
	e = initSysConfig()                                                                //初始化系统配置
	if e != nil {
		fmt.Printf("加载sys_ocnfig 的时候错误: %v", e)
		return e
	}

	//初始化mongo日志记录
	var mgoLoger MgoLoger
	mgo.SetLogger(mgoLoger)

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
