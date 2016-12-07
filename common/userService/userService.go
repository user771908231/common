package userService

import (
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
	"gopkg.in/mgo.v2"
	"errors"
	"casino_common/utils/db"
	"casino_common/common/log"
	"casino_common/utils/numUtils"
	"casino_common/utils/redisUtils"
	"casino_common/common/Error"
	"casino_common/common/model"
	"casino_common/common/consts/tableName"
	"casino_common/proto"
	"casino_server/service/userService"
	"casino_majiang/msg/protogo"
	"github.com/name5566/leaf/gate"
)

var NEW_USER_DIAMOND_REWARD int64 = 20                //新用户登陆的时候,默认的砖石数量
var USER_DIAMOND_REDIS_KEY = "user_diamond_redis_key"


/**
	1,create 一个user
	2,保存mongo
	3,缓存到redis
 */
func NewUserAndSave(openId, wxNickName, headUrl string, sex int32, city string) (*casinoCommonProto.User, error) {

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
	idStr, _ := numUtils.Uint2String(id)
	return strings.Join([]string{tableName.DBT_T_USER, idStr}, "-")
}

func ClearUserSeesion(userId uint32) {
	redisUtils.Del(GetRedisUserSeesionKey(userId))
}


//取session的rediskey
func GetRedisUserSeesionKey(userid uint32) string {
	idStr, _ := numUtils.Uint2String(userid)
	return strings.Join([]string{"agent_session", idStr}, "_")
}
/**
	根据用户id得到User的id
	1,首先从redis中查询user信息
	2,如果redis中不存在,则从mongo中查询
	3,如果mongo不存在,返回错误信息,客户端跳转到登陆界面

 */
func GetUserById(id uint32) *casinoCommonProto.User {
	log.T("userservice.GetUserById...")
	//1,首先在 redis中去的数据
	var buser *casinoCommonProto.User = nil
	result := redisUtils.GetObj(GetRedisUserKey(id), &casinoCommonProto.User{})
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
			SetUserDiamond(id, buser.GetDiamond())
		}

	} else {
		buser = result.(*casinoCommonProto.User)
	}

	//判断用户是否存在,如果不存在,则返回空
	return buser
}

/**
	将用户model保存在redis中
 */
func SaveUser2Redis(u *casinoCommonProto.User) {
	redisUtils.SetObj(GetRedisUserKey(u.GetId()), u)
}

/**
	保存数据到redis和mongo中
 */
func SaveUser2RedisAndMongo(u *casinoCommonProto.User) {
	SaveUser2Redis(u)
	UpsertRUser2Mongo(u)
}


//把redis中的数据刷新到数据库
func FlashUser2Mongo(userId uint32) error {
	u := GetUserById(userId)
	UpsertRUser2Mongo(u)
	return nil
}

func UpsertRUser2Mongo(u *casinoCommonProto.User) {
	//casinoCommonProto.User转化为  model.User
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
func Tuser2Ruser(tu *model.T_user) (*casinoCommonProto.User, error) {
	result := &casinoCommonProto.User{}
	if tu.Mid.Hex() != "" {
		hesStr := tu.Mid.Hex()
		result.Mid = &hesStr
		//log.T("获得t_user.mid %v",hesStr)
	}

	result.Id = &tu.Id
	result.NickName = &tu.NickName
	result.Coin = &tu.Coin
	result.Diamond = &tu.Diamond
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

func Ruser2Tuser(ru *casinoCommonProto.User) (*model.T_user, error) {
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



//更新用户的钻石之后,在放回用户当前的余额,更新用户钻石需要同事更新redis和mongo的数据
func UpdateUserDiamond(userId uint32, diamond int64) {
	user := GetUserById(userId)
	if user == nil {
		log.E("更新用户的diamond失败,用户为空")
		return
	}

	//修改并且更新用户数据
	*user.Diamond = diamond
	SaveUser2RedisAndMongo(user)
}

//用户砖石的key
func UDK(userId uint32) string {
	userIdStr, _ := numUtils.Uint2String(userId)
	result := strings.Join([]string{USER_DIAMOND_REDIS_KEY, userIdStr}, "_")
	return result
}

func GetUserDiamond(userId uint32) int64 {
	balance := redisUtils.GetInt64(UDK(userId))
	return balance
}

func SetUserDiamond(userId uint32, diamond int64) {
	redisUtils.SetInt64(UDK(userId), diamond)
}


//craete钻石交易记录

func CreateDiamonDetail(userId uint32, detailsType int32, diamond int64, remainDiamond int64, memo string) error {

	//1,获得的交易记录自增主键
	id, err := db.GetNextSeq(tableName.DBT_T_USER_DIAMOND_DETAILS)
	if err != nil {
		return Error.NewError(0, err.Error())
	}

	//2,构造交易记录
	detail := &model.T_user_diamond_details{}
	detail.Id = uint32(id)
	detail.UserId = userId
	detail.Diamond = diamond
	detail.ReaminDiamond = remainDiamond
	detail.DetailsType = detailsType
	detail.DetailTime = time.Now()
	detail.Memo = memo

	//3,保存数据
	err = db.InsertMgoData(tableName.DBT_T_USER_DIAMOND_DETAILS, detail)
	if err != nil {
		log.E("保存用户交易记录的时候失败 error【%v】", err.Error())
		return Error.NewError(0, "创建交易记录失败")
	}
	return nil
}

//增加用户的金额
func INCRUserDiamond(userid uint32, d int64) (int64, error) {
	log.T("为用户[%v]增加钻石[%v]", userid, d)
	//1,增加余额
	remain := redisUtils.INCRBY(UDK(userid), d)

	//2,更新redis和数据库中的数据
	UpdateUserDiamond(userid, remain)

	//3,返回值
	return remain, nil
}


//减少用户的砖石
func DECRUserDiamond(userid uint32, d int64) (int64, error) {
	diamond := GetUserDiamond(userid)

	//检测用户的余额是否足够
	if diamond < d {
		return diamond, errors.New("用户余额不足")
	}

	//减少用户的砖石数量
	remain := redisUtils.DECRBY(UDK(userid), d)
	if remain < 0 {
		INCRUserDiamond(userid, d)
		return diamond, errors.New("用户余额不足")
	} else {
		//更新redis和mongo中的数据
		UpdateUserDiamond(userid, remain)
		return remain, nil
	}

}

func GetUserByOpenId(openId  string) *casinoCommonProto.User {
	log.T("通过openId[%v]查询用户是否存在...", openId)
	//2,从数据库中查询
	result := &casinoCommonProto.User{}
	tuser := &model.T_user{}
	db.Query(func(d *mgo.Database) {
		d.C(tableName.DBT_T_USER).Find(bson.M{"openid": openId}).One(tuser)
	})

	if tuser == nil {
		log.T("在mongo中没有查询到user[%v].", openId)
		result = nil
	} else {
		log.T("在mongo中查询到了user.openId[%v],现在开始缓存", tuser.OpenId)
		//把从数据获得的结果填充到redis的model中
		result, _ = Tuser2Ruser(tuser)
		if result != nil {
			SaveUser2Redis(result)
			SetUserDiamond(result.GetId(), result.GetDiamond())
		}
	}

	//判断用户是否存在,如果不存在,则返回空
	return result
}

//处理用户登录
func handlerGame_Login(weixin *casinoCommonProto.WeixinInfo, a gate.Agent) {
}

