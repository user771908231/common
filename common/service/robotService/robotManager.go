package robotService

import (
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/model/userDao"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"casino_common/utils/rand"
	"casino_common/utils/redisUtils"
	"fmt"
	"github.com/golang/protobuf/proto"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ROBOT_AVAILABLE_TRUE  string = "true"
	ROBOT_AVAILABLE_FALSE string = "false"
)

type RobotsMgrApi interface {
	ExpropriationRobotByCoin(coin int64) *Robot
	ReleaseRobots(id uint32)
	ExpropriationRobotByCoin2(c1 int64, c2 int64) *Robot
}

//数据库
type RobotInfo struct {
	UserId    float64
	NickName  string
	HeaderUrl string
	Sex       float64
	City      string
	Available bool
	RegTime   time.Time
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

func (rm *RobotsManager) GetRobotAvailableRedisKey(robotId uint32) string {
	return fmt.Sprintf("%v_%v_%v", consts.RKEY_ROBOT_AVAILABLE, rm.gameId, robotId)
}

//初始化一个管理器
func (rm *RobotsManager) Oninit() {
	var users []*ddproto.User
	users = userDao.FindUsersByKV("robottype", int32(rm.gameId))
	//转换
	for _, u := range users {
		r := NewRobots(u)
		if rm.gameId == ddproto.CommonEnumGame_GID_DDZ {
			//这里没有数据恢复的游戏服务器需要特殊处理
			r.available = true
			rm.robotsAbleCount++
			continue
		}

		//同步redis的征用状态
		available := redisUtils.Get(rm.GetRobotAvailableRedisKey(u.GetId()))
		switch available {
		case "", ROBOT_AVAILABLE_TRUE:
			r.available = true
			rm.robotsAbleCount++
		case ROBOT_AVAILABLE_FALSE:
			r.available = false
		}
		rm.robots = append(rm.robots, r)
	}
	//rm.robotsAbleCount = int32(len(rm.robots)) //初始化机器人的数量
}

//获取所有机器人
func (rm *RobotsManager) GetRobots() []*Robot {
	return rm.robots
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

//创建一组机器人的方法
func (rm *RobotsManager) NewRobotsAndSave(num, minCoin, maxCoin int32) {
	log.T("请求创建%v个机器人 金币下限[%v] 金币上限[%v]", num, minCoin, maxCoin)
	for i := 0; num > 0; num-- {
		log.T("创建第%v个机器人", i)
		uid, err := userService.GetNewUserIdByIndex(int32(i))
		if err != nil {
			log.E("创建第%v个机器人时获取一个新的玩家id失败 已终止创建，err:%v", i, err)
			break
		}

		log.T("开始创建一个uid为%v的机器人", uid)
		rm.NewRobotAndSave(uid, minCoin, maxCoin)
		i++
	}
}

//新创建一个机器人，并保存到数据库 [minCoin, maxCoin)
func (rm *RobotsManager) NewRobotAndSave(uid uint32, minCoin, maxCoin int32) *Robot {
	defer Error.ErrorRecovery("NewRobotAndSave")
	//从数据库中获取一个可用的玩家信息
	var user *ddproto.User = nil
	var err error = nil
	robotInfos := []*RobotInfo{}
	var robotInfo *RobotInfo = nil

	ids, _ := numUtils.Uint2String(uid)
	nickName := "游客" + ids
	headUrl := ""
	sex := rand.Rand(0, 1)
	city := ""

	//查询数据库 获取机器人信息
	err = db.C(tableName.DBT_T_ROBOT_INFO).FindAll(bson.M{"available": "true"}, &robotInfos)
	if err != nil {
		log.E("新建一个机器人的时候 查询机器人库信息错误 err: %v", err)
	} else {
		index := int(rand.Rand(0, int32(len(robotInfos))))
		log.T("随机取机器人库存信息 可用库存[%v] index:[%v]", len(robotInfos), index)
		for i, info := range robotInfos {
			if i >= index && info != nil {
				robotInfo = info
				break
			}
		}
		if robotInfo == nil {
			for _, info := range robotInfos {
				if info != nil {
					robotInfo = info
				}
			}
		}
	}
	if robotInfo != nil {
		if robotInfo.NickName != "" {
			nickName = robotInfo.NickName
		}

		sex = int32(robotInfo.Sex)
		city = robotInfo.City

		if robotInfo.HeaderUrl != "" {
			headUrl = robotInfo.HeaderUrl
		}

		//这里已经用了
	}

	log.T("开始创建一个机器人uid:%v nickName:%v headUrl:%v sex:%v city:%v", uid, nickName, headUrl, sex, city)
	//channel : "robot"
	user, err = userService.NewUserAndSave(uid, "", fmt.Sprintf("%v", uid), nickName, headUrl, sex, city, "robot", "localhost")
	if err != nil || user == nil {
		log.E("创建一个机器人uid:%v nickName:%v headUrl:%v sex:%v city:%v 出错err:%v", uid, nickName, headUrl, sex, city, err)
		return nil
	}

	//更新状态
	if robotInfo != nil {
		err = db.C(tableName.DBT_T_ROBOT_INFO).Update(bson.M{"available": "true", "nickname": robotInfo.NickName, "sex": robotInfo.Sex, "headerurl": robotInfo.HeaderUrl}, bson.M{"$set": bson.M{"available": "false", "userid": uid, "regtime": time.Now()}})
		if err != nil {
			log.E("更新机器人[%v]数据库信息的时候失败err: %v", uid, err)
		}
	}

	//2,注册机器人
	user.RobotType = proto.Int32(int32(rm.gameId))
	randRobotCoin := rand.Rand(minCoin, maxCoin)

	//取随机数的整
	if randRobotCoin > 100 {
		randRobotCoin = int32(randRobotCoin/100) * 100
	} else {
		//默认值
		randRobotCoin = 5000
	}

	c, _ := userService.INCRUserCOIN(user.GetId(), int64(randRobotCoin), "创建新机器人初始化")
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

	redisUtils.Set(rm.GetRobotAvailableRedisKey(r.GetId()), ROBOT_AVAILABLE_TRUE)
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
			atomic.AddInt32(&rm.robotsAbleCount, -1) //可以使用的机器人数量-1
			redisUtils.Set(rm.GetRobotAvailableRedisKey(r.GetId()), ROBOT_AVAILABLE_FALSE)
			return r
		}
	}
	return nil
}

//注意 机器人金币
//func (rm *RobotsManager) ExpropriationRobotByCoin2(c1 int64, c2 int64) *Robot {
//	robot := rm.ExpropriationRobot()
//	if c2 > 10000000 {
//		c2 = 10000000
//	}
//	rand := rand.RandInt64(c1, c2)
//	if robot == nil {
//		return nil
//	}
//	userService.SetUserMoney(robot.GetId(), consts.RKEY_USER_COIN, rand) //设置机器人玩家的金币
//	robot.Coin = proto.Int64(rand)                                       //设置机器人的金额
//	return robot
//}

//得到一个机器人
func (rm *RobotsManager) ExpropriationRobotByCoin(coin int64) *Robot {
	rm.Lock()
	defer rm.Unlock()

	canExpRobots := []*Robot{}
	for _, r := range rm.robots {
		//log.T("机器人[%v]的coin %v,limit %v", r.GetId(), r.GetCoin(), coin)
		robotCoin := userService.GetUserCoin(r.GetId())
		if r.IsAvailable() && robotCoin >= coin {
			canExpRobots = append(canExpRobots, r)
		}
	}

	index := int(rand.Rand(0, int32(len(canExpRobots))))
	log.T("请求获取金币大于[%v]的机器人 找到可用的机器人数量[%v] 随机取index[%v]的机器人", coin, len(canExpRobots), index)
	for i, r := range canExpRobots {
		if i >= index {
			r.available = false
			atomic.AddInt32(&rm.robotsAbleCount, -1) //可以使用的机器人数量-1
			redisUtils.Set(rm.GetRobotAvailableRedisKey(r.GetId()), ROBOT_AVAILABLE_FALSE)
			//打印当前可以使用的机器人，注意，这里的可以使用只表示available == true 的情况，并不是coin足够的情况
			log.T("释放征用一个机器人[%v]之后，可以使用的机器人数量还剩下:%v", r.GetId(), rm.robotsAbleCount)
			return r
		}
	}

	//for _, r := range rm.robots {
	//	//log.T("机器人[%v]的coin %v,limit %v", r.GetId(), r.GetCoin(), coin)
	//	robotCoin := userService.GetUserCoin(r.GetId())
	//	if r.IsAvailable() && robotCoin >= coin {
	//		r.available = false
	//		atomic.AddInt32(&rm.robotsAbleCount, -1) //可以使用的机器人数量-1
	//		redisUtils.Set(rm.GetRobotAvailableRedisKey(r.GetId()), ROBOT_AVAILABLE_FALSE)
	//		//打印当前可以使用的机器人，注意，这里的可以使用只表示available == true 的情况，并不是coin足够的情况
	//		log.T("释放征用一个机器人[%v]之后，可以使用的机器人数量还剩下:%v", r.GetId(), rm.robotsAbleCount)
	//		return r
	//	}
	//}
	return nil
}

//通过左闭右开的金币区间得到一个机器人 [min, max)
func (rm *RobotsManager) ExpropriationRobotByRange(minCoin, maxCoin int64) *Robot {
	rm.Lock()
	defer rm.Unlock()

	for _, r := range rm.robots {
		robotCoin := userService.GetUserCoin(r.GetId())
		//log.T("机器人[%v]的coin %v,min[%v]~max[%v]", r.GetId(), robotCoin, minCoin, maxCoin)
		if r.IsAvailable() && robotCoin >= minCoin && robotCoin < maxCoin {
			r.available = false
			atomic.AddInt32(&rm.robotsAbleCount, -1) //可以使用的机器人数量-1
			redisUtils.Set(rm.GetRobotAvailableRedisKey(r.GetId()), ROBOT_AVAILABLE_FALSE)
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
	if r != nil && !r.available {
		r.available = true
		atomic.AddInt32(&rm.robotsAbleCount, 1) //可以使用的机器人数量+1
		redisUtils.Set(rm.GetRobotAvailableRedisKey(r.GetId()), ROBOT_AVAILABLE_TRUE)
		log.T("释放机器人[%v]之后，可以使用的机器人数量还剩下:%v", id, rm.robotsAbleCount)
		return
	}
	log.T("释放机器人[%v]失败，可以使用的机器人数量还剩下:%v", id, rm.robotsAbleCount)
}

func (rm *RobotsManager) ReleaseAllRobots() (releaseCount int32) {
	log.T("开始释放所有机器人 共有[%v]个机器人", len(rm.robots))
	for i, r := range rm.robots {
		if r != nil {
			log.T("开始释放第[%v]个机器人", i)
			rm.ReleaseRobots(r.GetId())
			releaseCount++
		}
	}
	return
}
