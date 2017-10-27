package userService

import (
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/model/userDao"
	"casino_common/common/sys"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/rand"
	"casino_common/utils/redisUtils"
	"errors"
	"github.com/golang/protobuf/proto"
	"time"
)

//查询一个id是否被占用
func IsExistUserId(id uint32) bool {
	user := GetUserById(uint32(id))
	return user != nil && user.GetId() == uint32(id)
}

//获取顺序增长且不重复的玩家id
func GetNewUserId() (uint32, error) {
	log.T("开始获取一个顺序增长且未被占用的玩家id")
	id, err := db.GetNextSeq(tableName.DBT_T_USER)
	if err != nil {
		//如果获取失败 返回err
		log.E("无法从数据库获取玩家自增id err:%v", err)
		return uint32(id), err
	}

	if IsExistUserId(uint32(id)) {
		log.W("获取到的玩家id[%v]被占用，重新获取", id)
		return GetNewUserId()
	}
	log.T("获取到一个顺序增长且未被占用的玩家id[%v]并返回", id)
	return uint32(id), nil
}

//获取离散且不重复的玩家id
func GetNewUserIdByIndex(index int32) (uint32, error) {
	log.T("开始根据序号%v获取一个离散且未被占用的玩家id", index)
	id, err := db.GetNextSeq(tableName.DBT_T_USER)
	if err != nil {
		//如果获取失败 返回err
		log.E("无法从数据库获取自增id err:%v", err)
		return 0, err
	}

	uid := int32(id) + index*int32(100) + rand.Rand(0, 100)

	if IsExistUserId(uint32(uid)) {
		log.W("获取到的玩家id[%v]被占用，重新获取", uid)
		return GetNewUserIdByIndex(index)
	}
	log.T("根据序号%v获取到一个离散且未被占用的玩家id[%v]并返回", index, uid)
	return uint32(uid), nil
}

/**
1,create 一个user
2,保存mongo
3,缓存到redis
*/
func NewUserAndSave(uid uint32, unionId, openId, wxNickName, headUrl string, sex int32, city string, channel string, regIp string) (*ddproto.User, error) {
	log.T("创建新用户，并且保存到mgo")
	//1,创建user获得自增主键
	var err error = nil

	//未指定玩家id 获取一个玩家id
	newUserId := uid
	if newUserId <= 0 {
		log.T("创建user时未指定合适的玩家id 开始获取一个玩家id")
		id, err := GetNewUserId()
		if err != nil {
			return nil, err
		}
		newUserId = id
	}
	log.T("开始创建一个uid为%v的新用户", newUserId)

	//构造user
	user := new(ddproto.User)
	user.Id = proto.Uint32(newUserId)
	user.Sex = proto.Int32(sex)
	user.City = proto.String(city)
	user.Diamond = proto.Int64(0)
	user.Diamond2 = proto.Int64(0)
	user.NickName = proto.String(wxNickName)
	user.UnionId = proto.String(unionId)
	user.OpenId = proto.String(openId)
	user.HeadUrl = proto.String(headUrl)
	user.Coin = proto.Int64(sys.CONFIG_SYS.GetNewUserCoin())
	user.RoomCard = proto.Int64(sys.CONFIG_SYS.GetNewUserRoomcard())
	user.Ticket = proto.Int32(0)
	user.Bonus = proto.Float64(0)
	user.RegTime = proto.Int64(time.Now().Unix())
	user.RegChannel = proto.String("weixin")
	user.RegIp = proto.String(regIp)
	user.AgentId = proto.Uint32(0)
	user.RegChannel = proto.String(channel)
	user.NewUserAward = proto.Bool(true) //注册之后是有新手奖励的,领取奖励之后设置为false
	//初始化默认值

	//2保存数据到数据库
	err = userDao.SaveUser2Mgo(user)
	if err != nil {
		log.E("保存用户到mongo的时候失败 error【%v】", err.Error())
		return nil, err
	}

	//3,保存到redis
	SaveUser2Redis(user)
	switch channel {
	case "61", "62":
		//对于白山channelId:61 和 62：新注册用户默认拥有1张房卡
		INCRUserRoomcard(user.GetId(), 1, 0, "新用户注册") //新用户注册的时候,默认的房卡数量
	case "1", "2":
		//对于四川的用户，每人给10张房卡
		INCRUserRoomcard(user.GetId(), 20, 0, "新用户注册")
	default:
		INCRUserRoomcard(user.GetId(), sys.CONFIG_SYS.GetNewUserRoomcard(), 0, "新用户注册") //新用户注册的时候,默认的房卡数量
	}

	INCRUserCOIN(user.GetId(), sys.CONFIG_SYS.GetNewUserCoin(), "新用户注册，设置默认金币数") //新用户注册的时候，默认的金币数量
	return user, nil
}

func GetRedisUserKey(id uint32) string {
	return redisUtils.K(consts.RKEY_PRE_USER, id)
}

/**
根据用户id得到User的id
1,首先从redis中查询user信息
2,如果redis中不存在,则从mongo中查询
3,如果mongo不存在,返回错误信息,客户端跳转到登陆界面

*/
func GetUserById(id uint32) *ddproto.User {

	//1,首先在 redis中去的数据
	var buser *ddproto.User = nil

	result := redisUtils.GetObj(GetRedisUserKey(id), &ddproto.User{})

	//2，从redis 中取到的数据不为空，那么直接返回
	if result != nil {
		buser = result.(*ddproto.User)
		//下三行代码 可以暂时不用，这里很印象性能
		buser.Coin = proto.Int64(GetUserCoin(id))
		buser.Diamond = proto.Int64(GetUserDiamond(id))
		buser.Diamond2 = proto.Int64(GetUserDiamond2(id))
		buser.RoomCard = proto.Int64(GetUserRoomCard(id))
		return buser
	}

	//3，从reids中取到的数据为空，那么从数据库中读取
	if result == nil {
		log.T("redis中没有找到user[%v],需要在mongo中查询,并且缓存在redis中。", id)
		buser = userDao.FindUserById(id)
		if buser != nil {
			log.T("在mongo中查询到了user[%v],现在开始缓存", buser)
			SaveUser2Redis(buser)
			InitUserMoney2Redis(buser)
		} else {
			log.E("在mongo中没有找到user[%v],玩家不存在", id)

		}
	}

	//判断用户是否存在,如果不存在,则返回空
	return buser
}

/**
根据用户userName即phonenumber得到User
1,从mongo中查询
3,如果mongo不存在,返回错误信息,客户端跳转到登陆界面

*/
func GetUserByUserName(userName string) *ddproto.User {

	buser := userDao.FindUserByKV("phonenumber", userName)

	if buser != nil {
		return GetUserById(*buser.Id)
	}
	//判断用户是否存在,如果不存在,则返回空
	return buser
}

/**
将用户model保存在redis中
*/
func SaveUser2Redis(u *ddproto.User) error {
	return redisUtils.SetObj(GetRedisUserKey(u.GetId()), u)
}

func UpdateUser2MgoById(userId uint32) {
	user := GetUserById(userId)
	UpdateUser2Mgo(user) //通过id update user
}

//把用户保存到mgo并且同事更新redis的数据
func UpdateUser2Mgo(u *ddproto.User) {
	userDao.UpdateUser2Mgo(u)
	SaveUser2Redis(u) //保存user到redis
}

//初始化redis中用户的金额的值
func InitUserMoney2Redis(user *ddproto.User) {
	SetUserMoney(user.GetId(), consts.RKEY_USER_COIN, user.GetCoin())
	SetUserMoney(user.GetId(), consts.RKEY_USER_DIAMOND, user.GetDiamond())
	SetUserMoney(user.GetId(), consts.RKEY_USER_DIAMOND2, user.GetDiamond2())
	SetUserMoney(user.GetId(), consts.RKEY_USER_ROOMCARD, user.GetRoomCard())
}

func GetUserByUnionId(unionid string) *ddproto.User {
	log.T("通过UnionIdd[%v]查询用户是否存在...", unionid)
	//2,从数据库中查询
	user := userDao.FindUserByUnionId(unionid)
	if user != nil {
		user = GetUserById(user.GetId())
	}
	//判断用户是否存在,如果不存在,则返回空
	return user
}

//更新redis user 的各种money ,同步之后会save到redis中
func SyncReidsUserMoney(user *ddproto.User) error {
	user.Coin = proto.Int64(GetUserCoin(user.GetId()))         //更新金币
	user.Diamond = proto.Int64(GetUserDiamond(user.GetId()))   //更新钻石
	user.Diamond2 = proto.Int64(GetUserDiamond2(user.GetId())) //更新钻石2
	user.RoomCard = proto.Int64(GetUserRoomCard(user.GetId())) //更新房卡
	return SaveUser2Redis(user)
}

//调用此方法 保证mongey,redisuser,mgo 的数据一致
func SyncMgoUserMoney(userId uint32) error {
	user := GetUserById(userId)
	if user == nil {
		return errors.New("user is nil.")
	}
	err := SyncReidsUserMoney(user)
	if err != nil {
		return err
	}
	return userDao.UpdateUser2Mgo(user) //保存用户到mgo
}
