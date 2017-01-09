package userService

import (
	"casino_common/utils/db"
	"casino_common/common/log"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/common/sys"
	"github.com/golang/protobuf/proto"
	"casino_common/common/model/userDao"
)

var USER_REDIS_KEY_AGENT_SESSION = "agent_session" //用户session的key
var USER_REDIS_KEY = tableName.DBT_T_USER          //用户的key

/**
	1,create 一个user
	2,保存mongo
	3,缓存到redis
 */
func NewUserAndSave(unionId, openId, wxNickName, headUrl string, sex int32, city string) (*ddproto.User, error) {
	log.T("创建新用户，并且保存到mgo")
	//1,创建user获得自增主键
	id, err := db.GetNextSeq(tableName.DBT_T_USER)
	if err != nil {
		return nil, err
	}

	//构造user
	user := new(ddproto.User)
	user.Id = proto.Uint32(uint32(id))
	user.Sex = proto.Int32(sex)
	user.City = proto.String(city)
	user.Diamond = proto.Int64(NEW_USER_DIAMOND_REWARD) //新用户注册的时候,默认的钻石数量
	user.NickName = proto.String(wxNickName)
	user.UnionId = proto.String(unionId)
	user.OpenId = proto.String(openId)
	user.HeadUrl = proto.String(headUrl)

	//初始化默认值

	//2保存数据到数据库
	err = userDao.SaveUser2Mgo(user)
	if err != nil {
		log.E("保存用户到mongo的时候失败 error【%v】", err.Error())
		return nil, err
	}

	//3,保存到redis
	SaveUser2Redis(user)
	return user, nil
}

func GetRedisUserKey(id uint32) string {
	return redisUtils.K(USER_REDIS_KEY, id)
}

func ClearUserSeesion(userId uint32) {
	redisUtils.Del(GetRedisUserSeesionKey(userId))
}

//取session的rediskey
func GetRedisUserSeesionKey(userid uint32) string {
	return redisUtils.K(USER_REDIS_KEY_AGENT_SESSION, userid)
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
	}

	//3，从reids中取到的数据为空，那么从数据库中读取
	if result == nil {
		log.E("redis中没有找到user[%v],需要在mongo中查询,并且缓存在redis中。", id)
		tuser := userDao.FindUserById(id)
		log.T("在mongo中查询到了user[%v],现在开始缓存", tuser)
		if buser != nil {
			SaveUser2Redis(buser)
			InitUserMoney2Redis(buser)
		}
	}

	//4,如果找不到用户并且是开发这模式
	if buser == nil &&  sys.GAMEENV.DEVMODE {
		log.T("现在是开发者模式sys.GAMEENV.DEVMODE[%v],生成一个对应id[%v]的user，并且缓存到redis中...", sys.GAMEENV.DEVMODE, id)
		buser = &ddproto.User{}
		buser.Id = proto.Uint32(id)
		nickName, _ := numUtils.Uint2String(id)
		buser.NickName = proto.String(nickName)
		buser.Coin = proto.Int64(0)
		buser.Diamond = proto.Int64(20)
		buser.Diamond2 = proto.Int64(0)
		buser.RoomCard = proto.Int64(0)
		buser.LastSignTime = proto.String("")
		buser.LastDrawLotteryTime = proto.String("")
		//保存到redis
		SaveUser2Redis(buser)
		InitUserMoney2Redis(buser)
	}

	//判断用户是否存在,如果不存在,则返回空
	return buser
}

/**
	将用户model保存在redis中
 */
func SaveUser2Redis(u *ddproto.User) {
	redisUtils.SetObj(GetRedisUserKey(u.GetId()), u)
}

//把用户保存到mgo
func UpdateUser2Mgo(u *ddproto.User) {
	userDao.UpdateUser2Mgo(u)
}

//初始化redis中用户的金额的值
func InitUserMoney2Redis(user *ddproto.User) {
	SetUserMoney(user.GetId(), USER_COIN_REDIS_KEY, user.GetCoin())
	SetUserMoney(user.GetId(), USER_DIAMOND_REDIS_KEY, user.GetDiamond())
	SetUserMoney(user.GetId(), USER_DIAMOND2_REDIS_KEY, user.GetDiamond2())
}

func GetUserByOpenId(openId string) *ddproto.User {
	log.T("通过openId[%v]查询用户是否存在...", openId)
	//2,从数据库中查询
	user := userDao.FindUserByUnionId(openId)
	if user != nil {
		log.T("在mongo中查询到了user.openId[%v],现在开始缓存", user.OpenId)
		//把从数据获得的结果填充到redis的model中
		SaveUser2Redis(user)
		InitUserMoney2Redis(user)
	}
	//判断用户是否存在,如果不存在,则返回空
	return user
}
