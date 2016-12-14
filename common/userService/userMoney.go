package userService

import (
	"time"
	"errors"
	"casino_common/utils/db"
	"casino_common/common/log"
	"casino_common/utils/redisUtils"
	"casino_common/common/Error"
	"casino_common/common/model"
	"casino_common/common/consts/tableName"
)

var NEW_USER_DIAMOND_REWARD int64 = 20                //新用户登陆的时候,默认的砖石数量

var USER_COIN_REDIS_KEY = "user_coin_redis_key"        //金币
var USER_DIAMOND_REDIS_KEY = "user_diamond_redis_key"        //钻石，金币场
var USER_DIAMOND2_REDIS_KEY = "user_diamond2_redis_key"        //钻石朋友桌

//更新用户的钻石之后,在放回用户当前的余额,更新用户钻石需要同事更新redis和mongo的数据

//更新用户的数据
func UpdateUserMoney(userId uint32) {
	user := GetUserById(userId)
	if user == nil {
		log.E("更新用户的diamond失败,用户为空")
		return
	}

	//修改并且更新用户数据
	*user.Coin = GetUserMoney(userId, USER_COIN_REDIS_KEY)
	*user.Diamond = GetUserMoney(userId, USER_DIAMOND_REDIS_KEY)
	*user.Diamond2 = GetUserMoney(userId, USER_DIAMOND2_REDIS_KEY)
	SaveUser2Redis(user)
}

func GetUserMoney(userId uint32, money string) int64 {
	balance := redisUtils.GetInt64(GetKey(money, userId))
	return balance
}

func SetUserMoney(userId uint32, money string, diamond int64) {
	redisUtils.SetInt64(GetKey(money, userId), diamond)
}

//获取用户的钻石
func GetUserDiamond(userId uint32) int64 {
	return GetUserMoney(userId, USER_DIAMOND_REDIS_KEY)
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

//---------------------------------------增加用户货币的通用方法-----------------------------------------
//增加用户的货币
func incrUser(userid uint32, key string, d int64) (int64, error) {
	log.T("为用户[%v]增加钻石[%v]", userid, d)
	//1,增加余额
	remain := redisUtils.INCRBY(GetKey(key, userid), d)
	//2,更新redis和数据库中的数据
	UpdateUserMoney(userid)
	//3,返回值
	return remain, nil
}

//减少用户的货币
func decrUser(userid uint32, key string, d int64) (int64, error) {
	remain := redisUtils.DECRBY(GetKey(key, userid), d)
	if remain < 0 {
		log.E("用户[%v]的余额diamond不足,减少的时候失败")
		incrUser(userid, key, d)
		return remain, errors.New("用户余额不足")
	} else {
		//更新redis和mongo中的数据
		UpdateUserMoney(userid)
		return remain, nil
	}
}

//增加用户的钻石
func INCRUserDiamond(userid uint32, d int64) (int64, error) {
	return incrUser(userid, USER_DIAMOND_REDIS_KEY, d)
}

//减少用户的砖石
func DECRUserDiamond(userid uint32, d int64) (int64, error) {
	return decrUser(userid, USER_DIAMOND_REDIS_KEY, d)
}

//增加用户的朋友桌钻石
func INCRUserDiamond2(userid uint32, d int64) (int64, error) {
	return incrUser(userid, USER_DIAMOND2_REDIS_KEY, d)
}

//减少用户的朋友桌钻石
func DECRUserDiamond2(userid uint32, d int64) (int64, error) {
	return decrUser(userid, USER_DIAMOND2_REDIS_KEY, d)
}

//增加用户的金币
func INCRUserCOIN(userid uint32, d int64) (int64, error) {
	return incrUser(userid, USER_COIN_REDIS_KEY, d)
}

//减少用户的金币
func DECRUserCOIN(userid uint32, d int64) (int64, error) {
	return decrUser(userid, USER_COIN_REDIS_KEY, d)
}





