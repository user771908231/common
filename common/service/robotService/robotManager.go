package robotService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/model/userDao"
	"sync"
	"casino_common/common/userService"
	"github.com/golang/protobuf/proto"
	"casino_common/common/log"
	"casino_common/utils/numUtils"
	"sync/atomic"
	"casino_common/utils/rand"
	"casino_common/common/consts"
)

type RobotsMgrApi interface {
	ExpropriationRobotByCoin(coin int64) *Robot
	ReleaseRobots(id uint32)
	ExpropriationRobotByCoin2(c1 int64, c2 int64) *Robot
}

//机器人管理器
type RobotsManager struct {
	sync.Mutex
	gameId          ddproto.CommonEnumGame //游戏类型
	robots          []*Robot               //机器人
	robotsAbleCount int32                  //可以使用的机器人数量
}

//新建一个管理器
func NewRobotManager(gameId ddproto.CommonEnumGame) *RobotsManager {
	log.T("初始化NewRobotManager[%v]..", gameId)
	manager := &RobotsManager{
		gameId: gameId,
	}
	manager.Oninit()
	log.T("gameId[%v]目前机器人的数量:%v", gameId, len(manager.robots))
	return manager
}

//初始化一个管理器
func (rm *RobotsManager) Oninit() {
	var users []*ddproto.User
	users = userDao.FindUsersByKV("robottype", int32(rm.gameId))
	//转换
	for _, u := range users {
		r := NewRobots(u)
		r.available = true
		rm.robots = append(rm.robots, r)
	}
	rm.robotsAbleCount = int32(len(rm.robots)) //初始化机器人的数量
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
func (rm *RobotsManager) NewRobotAndSave() *Robot {
	//1,注册普通用户
	user, err := userService.NewUserAndSave("", "", "", "", 1, "", "robot", "localhost")
	if err != nil || user == nil {
		return nil
	}

	//2,注册机器人
	user.RobotType = proto.Int32(int32(rm.gameId))
	ids, _ := numUtils.Uint2String(user.GetId())
	user.NickName = proto.String("游客" + ids)
	c, _ := userService.INCRUserCOIN(user.GetId(), 50000)
	user.Coin = proto.Int64(c)
	//userService.SaveUser2Redis(user) //保存到redis
	userService.UpdateUser2Mgo(user) //创建机器人保存到mgo

	robot := &Robot{
		User:      user,
		available: true,
	}
	rm.addRobot(robot)
	log.T("目前机器人的数量 %v ，新增加的 %v", rm.robotsAbleCount, robot)

	//返回增加的机器人
	return robot
}

//增加一个机器人
func (rm *RobotsManager) addRobot(r *Robot) error {
	rm.Lock()
	rm.Unlock()
	rm.robots = append(rm.robots, r)
	atomic.AddInt32(&rm.robotsAbleCount, 1) //可以使用的机器人数量+1
	return nil
}

//得到一个机器人,不限制金额
func (rm *RobotsManager) ExpropriationRobot() *Robot {
	rm.Lock()
	defer rm.Unlock()

	tall := len(rm.robots)
	randIndex := rand.Rand(0, int32(tall))
	for i := 0; i < int(rm.robotsAbleCount); i++ {
		r := rm.robots[(i+int(randIndex))%tall]
		if r.available {
			r.available = false
			return r
		}
	}
	return nil
}

func (rm *RobotsManager) ExpropriationRobotByCoin2(c1 int64, c2 int64) *Robot {
	robot := rm.ExpropriationRobot()
	if c2 > 10000000 {
		c2 = 10000000
	}
	rand := rand.RandInt64(c1, c2)
	userService.SetUserMoney(robot.GetId(), consts.RKEY_USER_COIN, rand) //设置机器人玩家的金币
	robot.Coin = proto.Int64(rand)                                    //设置机器人的金额
	return robot
}

//得到一个机器人
func (rm *RobotsManager) ExpropriationRobotByCoin(coin int64) *Robot {
	rm.Lock()
	defer rm.Unlock()

	for _, r := range rm.robots {
		//log.T("机器人[%v]的coin %v,limit %v", r.GetId(), r.GetCoin(), coin)
		robotCoin := userService.GetUserCoin(r.GetId())
		if r.IsAvailable() && robotCoin >= coin {
			r.available = false
			atomic.AddInt32(&rm.robotsAbleCount, -1) //可以使用的机器人数量-1
			//打印当前可以使用的机器人，注意，这里的可以使用只表示available == true 的情况，并不是coin足够的情况
			log.T("释放征用一个之后，可以使用的机器人数量还剩下:%v", rm.robotsAbleCount)

			return r
		}
	}
	return nil
}

//通过左闭右开的金币区间得到一个机器人
func (rm *RobotsManager) ExpropriationRobotByRange(minCoin, maxCoin int64) *Robot {
	rm.Lock()
	defer rm.Unlock()

	for _, r := range rm.robots {
		robotCoin := userService.GetUserCoin(r.GetId())
		log.T("机器人[%v]的coin %v,min[%v]~max[%v]", r.GetId(), robotCoin, minCoin, maxCoin)
		if r.IsAvailable() && robotCoin >= minCoin && robotCoin < maxCoin {
			r.available = false
			atomic.AddInt32(&rm.robotsAbleCount, -1) //可以使用的机器人数量-1
			//打印当前可以使用的机器人，注意，这里的可以使用只表示available == true 的情况，并不是coin足够的情况
			log.T("释放征用一个之后，可以使用的机器人数量还剩下:%v", rm.robotsAbleCount)

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
	atomic.AddInt32(&rm.robotsAbleCount, 1) //可以使用的机器人数量+1
	log.T("释放一个机器人之后，可以使用的机器人数量还剩下:%v", rm.robotsAbleCount)

}
