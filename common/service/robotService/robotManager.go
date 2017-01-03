package robotService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/model/userDao"
	"sync"
)

//机器人管理器
type RobotsManager struct {
	sync.Mutex
	gameId ddproto.CommonEnumGame //游戏类型
	robots []*Robots              //机器人
}

//新建一个管理器
func NewRobotManager(gameId ddproto.CommonEnumGame) *RobotsManager {
	return &RobotsManager{
		gameId:gameId,
	}
	return nil
}

//初始化一个管理器
func (rm *RobotsManager) Oninit() {
	var users []*ddproto.User
	switch rm.gameId {
	case ddproto.CommonEnumGame_GID_MAHJONG: //初始化麻将
		users = userDao.FindUsersByKV("", "")
	case ddproto.CommonEnumGame_GID_DDZ: //初始化斗地主
	case ddproto.CommonEnumGame_GID_ZJH: //初始化炸金花

	}

	//转换
	for _, u := range users {
		r := NewRobots(u)
		rm.robots = append(rm.robots, r)
	}
}

//通过id得到一个机器人
func (rm *RobotsManager) getRobotById(id uint32) *Robots {
	for _, r := range rm.robots {
		if r.GetId() == id {
			return r
		}
	}
	return nil
}

//得到一个机器人
func (rm *RobotsManager) expropriationRobots() *Robots {
	rm.Lock()
	defer rm.Unlock()
	for _, r := range rm.robots {
		if r.IsAvailable() {
			return r
		}
	}
	return nil
}

//释放一个机器人,那这个机器人就可以被其他的人使用了....
func (rm *RobotsManager) releaseRobots(id uint32) {
	r := rm.getRobotById(id)
	r.available = true
}
