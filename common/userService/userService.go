package userService

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"gopkg.in/mgo.v2"
	"casino_common/utils/db"
	"casino_common/common/log"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/common/Error"
	"casino_common/common/model"
	"casino_common/common/consts/tableName"
	"casino_common/proto/ddproto"
	"casino_common/common/sys"
	"github.com/golang/protobuf/proto"
)

var USER_REDIS_KEY_AGENT_SESSION = "agent_session"        //用户session的可以

//得到
func GetKey(pre string, userId uint32) string {
	userIdStr, _ := numUtils.Uint2String(userId)
	result := strings.Join([]string{pre, userIdStr}, "_")
	return result
}

/**
	1,create 一个user
	2,保存mongo
	3,缓存到redis
 */
func NewUserAndSave(openId, wxNickName, headUrl string, sex int32, city string) (*ddproto.User, error) {

	//1,创建user获得自增主键
	id, err := db.GetNextSeq(tableName.DBT_T_USER)
	if err != nil {
		return nil, err
	}
	//构造user
	nuser := &model.T_user{}
	nuser.Mid = bson.NewObjectId()
	nuser.Id = uint32(id)
	nuser.Sex = sex
	nuser.City = city
	nuser.Diamond = NEW_USER_DIAMOND_REWARD                //新用户注册的时候,默认的钻石数量
	nuser.NickName = wxNickName
	nuser.OpenId = openId
	nuser.HeadUrl = headUrl

	//2保存数据到数据库
	err = db.InsertMgoData(tableName.DBT_T_USER, nuser)
	if err != nil {
		log.E("保存用户的时候失败 error【%v】", err.Error())
		return nil, err
	}

	result, _ := Tuser2Ruser(nuser)
	return result, nil
}

func GetRedisUserKey(id uint32) string {
	return GetKey(tableName.DBT_T_USER, id)
}

func ClearUserSeesion(userId uint32) {
	redisUtils.Del(GetRedisUserSeesionKey(userId))
}


//取session的rediskey
func GetRedisUserSeesionKey(userid uint32) string {
	return GetKey(USER_REDIS_KEY_AGENT_SESSION, userid)
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
		// 获取连接 connection
		tuser := &model.T_user{}
		db.Query(func(d *mgo.Database) {
			d.C(tableName.DBT_T_USER).Find(bson.M{"id": id}).One(tuser)
		})

		log.T("在mongo中查询到了user[%v],现在开始缓存", tuser)
		//把从数据获得的结果填充到redis的model中
		buser, _ = Tuser2Ruser(tuser)
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
		buser.Diamond = proto.Int64(20)
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

//初始化redis中用户的金额的值
func InitUserMoney2Redis(user *ddproto.User) {
	SetUserMoney(user.GetId(), USER_COIN_REDIS_KEY, user.GetCoin())
	SetUserMoney(user.GetId(), USER_DIAMOND_REDIS_KEY, user.GetDiamond())
	SetUserMoney(user.GetId(), USER_DIAMOND2_REDIS_KEY, user.GetDiamond2())
}


//把redis中的数据刷新到数据库
func FlashUser2Mongo(userId uint32) error {
	u := GetUserById(userId)
	UpsertRUser2Mongo(u)
	return nil
}

func UpsertRUser2Mongo(u *ddproto.User) {
	//ddproto.User转化为  model.User
	tuser, _ := Ruser2Tuser(u)        //
	UpsertTUser2Mongo(tuser)
}

//保存用户到mongo
func UpsertTUser2Mongo(tuser *model.T_user) {
	//得到数据库连接池
	if tuser.Mid == "" {
		db.InsertMgoData(tableName.DBT_T_USER, tuser)
	} else {
		db.UpdateMgoData(tableName.DBT_T_USER, tuser)
	}
}

/**
	mongo中User模型转化为 redis中的user模型
 */
func Tuser2Ruser(tu *model.T_user) (*ddproto.User, error) {

	//检测suer是否存在
	if tu.Id <= 0 {
		return nil, Error.NewFailError("user不存在")
	}

	//返回转换之后的结果
	result := &ddproto.User{}
	if tu.Mid.Hex() != "" {
		hesStr := tu.Mid.Hex()
		result.Mid = &hesStr
		//log.T("获得t_user.mid %v",hesStr)
	}

	result.Id = &tu.Id
	result.NickName = &tu.NickName
	result.Coin = &tu.Coin
	result.Diamond = &tu.Diamond
	result.Diamond2 = &tu.Diamond2
	result.UnionId = &tu.UnionId
	result.OpenId = &tu.OpenId
	result.HeadUrl = &tu.HeadUrl
	result.Sex = &tu.Sex
	result.City = &tu.City
	return result, nil
}

/**
	redis中的user模型转化为mongdo的User模型
	把Redis_user 转化为mongo_t_user的时候喂自动为其分配objectId,方存储
 */

func Ruser2Tuser(ru *ddproto.User) (*model.T_user, error) {
	result := &model.T_user{}

	if ru.Mid != nil {
		result.Mid = bson.ObjectIdHex(ru.GetMid())
	} else {
		result.Mid = bson.NewObjectId()
	}

	result.Id = ru.GetId()
	result.NickName = ru.GetNickName()
	result.Coin = ru.GetCoin()
	result.HeadUrl = ru.GetHeadUrl()
	result.OpenId = ru.GetOpenId()
	result.Diamond = ru.GetDiamond()
	result.Sex = ru.GetSex()
	result.City = ru.GetCity()

	return result, nil
}

/**

 */
func CheckUserIdRightful(userId uint32) bool {
	u := GetUserById(userId)
	if u == nil {
		return false
	} else {
		return true
	}
}

func GetUserByOpenId(openId  string) *ddproto.User {
	log.T("通过openId[%v]查询用户是否存在...", openId)
	//2,从数据库中查询
	ruser := &ddproto.User{}
	tuser := &model.T_user{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER).Find(bson.M{"openid": openId}).One(tuser)
	})

	if tuser == nil {
		log.T("在mongo中没有查询到user[%v].", openId)
		ruser = nil
	} else {
		log.T("在mongo中查询到了user.openId[%v],现在开始缓存", tuser.OpenId)
		//把从数据获得的结果填充到redis的model中
		ruser, _ = Tuser2Ruser(tuser)
		if ruser != nil {
			SaveUser2Redis(ruser)
			InitUserMoney2Redis(ruser)
		}
	}

	//判断用户是否存在,如果不存在,则返回空
	return ruser
}

