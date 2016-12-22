package taskService

import (
	"casino_common/proto/ddproto"
	"sync"
)

var allTask []ddproto.TTask
var once sync.Once

//做初始化
func init() {
	once.Do(initAllTask)
}

//生成任务
/**
	目前的任务只有跟场次有关系....
	1,任务都是通用的，不需要记录玩家没有完成的任务，只需要记录玩家 完成的任务....
	2,这里采用单列的模式...
*/
func initAllTask() {
	//读取任务配置(配置文件，数据库等)

}

func GetAllTask() []ddproto.TTask {
	return allTask
}

func GetUserTask(userId uint32) []*ddproto.TTask {
	//更具用户的已经完成的任务和所有的任务来区分

	//ddproto.TUserTask{}	//这个是用户已经完成的任务
	return nil
}
