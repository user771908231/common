package robotService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/model/userDao"
	"sync"
	"casino_common/common/userService"
	"github.com/golang/protobuf/proto"
	"casino_common/common/log"
)

//机器人管理器
type RobotsManager struct {
	sync.Mutex
	gameId ddproto.CommonEnumGame //游戏类型
	robots []*Robot               //机器人
}

//新建一个管理器
func NewRobotManager(gameId ddproto.CommonEnumGame) *RobotsManager {
	log.T("初始化NewRobotManager[%v]..", gameId)
	manager := &RobotsManager{
		gameId:gameId,
	}
	manager.Oninit()
	log.T("gameId[%v]目前机器人的数量:%v", gameId, len(manager.robots))
	return manager
}

//初始化一个管理器
func (rm *RobotsManager) Oninit() {
	var users []*ddproto.User

	//switch rm.gameId {
	//case ddproto.CommonEnumGame_GID_MAHJONG: //初始化麻将
	//	users = userDao.FindUsersByKV("robottype", rm.gameId)
	//case ddproto.CommonEnumGame_GID_DDZ: //初始化斗地主
	//case ddproto.CommonEnumGame_GID_ZJH: //初始化炸金花
	//
	//}

	users = userDao.FindUsersByKV("robottype", int32(rm.gameId))
	//转换
	for _, u := range users {
		r := NewRobots(u)
		r.available = true
		rm.robots = append(rm.robots, r)
	}
}

//通过id得到一个机器人
func (rm *RobotsManager) getRobotById(id uint32) *Robot {
	for _, r := range rm.robots {
		if r.GetId() == id {
			return r
		}
	}
	return nil
}

//新创建一个机器人，并保存到数据库
func (rm *RobotsManager) newRobotAndSave() *Robot {
	//1,注册普通用户
	user, err := userService.NewUserAndSave("", "", "", "", 1, "")
	if err != nil || user == nil {
		return nil
	}

	//2,注册机器人
	user.RobotType = proto.Int32(int32(rm.gameId))
	userService.SaveUser2Redis(user) //保存到redis
	userService.UpdateUser2Mgo(user) //保存到mgo

	robot := &Robot{
		User:     user,
		available:true,
	}
	log.T("新创建的机器人:%v", robot)
	rm.addRobot(robot)

	//返回增加的机器人
	return robot
}

//增加一个机器人
func (rm *RobotsManager) addRobot(r *Robot) error {
	rm.Lock()
	rm.Unlock()
	rm.robots = append(rm.robots, r)
	return nil
}

//得到一个机器人
func (rm *RobotsManager) ExpropriationRobot() *Robot {
	rm.Lock()
	defer rm.Unlock()

	for _, r := range rm.robots {
		if r.IsAvailable() {
			r.available = false
			return r
		}
	}
	return nil
}

//释放一个机器人,那这个机器人就可以被其他的人使用了....
func (rm *RobotsManager) ReleaseRobots(id uint32) {
	rm.Lock()
	defer rm.Unlock()
	r := rm.getRobotById(id)
	r.available = true
}
