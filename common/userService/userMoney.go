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
	"casino_common/common/cfg"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
)

//更新用户的钻石之后,在放回用户当前的余额,更新用户钻石需要同事更新redis和mongo的数据

//更新用户的数据
func UpdateRedisUserMoney(userId uint32) error {
	user := GetUserById(userId)
	if user == nil {
		log.E("更新用户的diamond失败,用户[%v]为空", userId)
		return errors.New("增加money失败,没有找到用户")
	}

	//修改并且更新用户数据
	SyncReidsUserMoney(user)
	return nil
}

func GetUserMoney(userId uint32, money string) int64 {
	balance := redisUtils.GetInt64(redisUtils.K(money, userId))
	return balance
}

func SetUserMoney(userId uint32, money string, diamond int64) {
	redisUtils.SetInt64(redisUtils.K(money, userId), diamond)
}

//获取用户的钻石
func GetUserDiamond(userId uint32) int64 {
	return GetUserMoney(userId, cfg.RKEY_USER_DIAMOND)
}

//获取用户的钻石
func GetUserDiamond2(userId uint32) int64 {
	return GetUserMoney(userId, cfg.RKEY_USER_DIAMOND2)
}

func GetUserRoomCard(userId uint32) int64 {
	return GetUserMoney(userId, cfg.RKEY_USER_ROOMCARD)
}

func GetUserCoin(userId uint32) int64 {
	return GetUserMoney(userId, cfg.RKEY_USER_COIN)
}

//获取用户奖券
func GetUserTicket(userId uint32) int32 {
	user := new(ddproto.User)
	db.C(tableName.DBT_T_USER).Find(bson.M{"id": userId}, user)
	return user.GetTicket()
}

//获取用户红包
func GetUserBonus(userId uint32) float64 {
	user := new(ddproto.User)
	db.C(tableName.DBT_T_USER).Find(bson.M{"id": userId}, user)
	return float64(int64(user.GetBonus()*100)) / 100
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
	log.T("为用户[%v]增加[%v][%v]", userid, key, d)
	//1,增加余额
	remain := redisUtils.INCRBY(redisUtils.K(key, userid), d)
	//2,更新redis和数据库中的数据
	err := UpdateRedisUserMoney(userid) //增加用户货币之后
	//3,返回值
	return remain, err
}

//减少用户的货币
func decrUser(userid uint32, key string, d int64) (int64, error) {
	remain := redisUtils.DECRBY(redisUtils.K(key, userid), d)
	if remain < 0 {
		old, _ := incrUser(userid, key, d)
		errMsg := fmt.Sprintf("用户[%v]的key[%v][%v]不足,减少的时候失败", userid, key, old)
		log.E(errMsg)
		return remain, errors.New(errMsg)
	} else {
		//更新redis和mongo中的数据
		UpdateRedisUserMoney(userid)
		return remain, nil
	}
}

//增加用户的钻石
func INCRUserDiamond(userid uint32, d int64) (int64, error) {
	return incrUser(userid, cfg.RKEY_USER_DIAMOND, d)
}

//减少用户的砖石
func DECRUserDiamond(userid uint32, d int64) (int64, error) {
	count := GetUserDiamond(userid)
	if count-d < 0 {
		return count, errors.New("余额不足，减少钻石失败！")
	}
	return decrUser(userid, cfg.RKEY_USER_DIAMOND, d)
}

//增加用户的房卡
func INCRUserRoomcard(userId uint32, d int64) (int64, error) {
	return incrUser(userId, cfg.RKEY_USER_ROOMCARD, d)
}

//减少用户的房卡
func DECRUserRoomcard(userId uint32, d int64) (int64, error) {
	count := GetUserRoomCard(userId)
	if count-d < 0 {
		return count, errors.New("余额不足，减少房卡失败！")
	}
	return decrUser(userId, cfg.RKEY_USER_ROOMCARD, d)
}

//增加用户的朋友桌钻石
func INCRUserDiamond2(userid uint32, d int64) (int64, error) {
	return incrUser(userid, cfg.RKEY_USER_DIAMOND2, d)
}

//减少用户的朋友桌钻石
func DECRUserDiamond2(userid uint32, d int64) (int64, error) {
	count := GetUserDiamond2(userid)
	if count-d < 0 {
		return count, errors.New("余额不足，减少朋友桌钻石失败！")
	}
	return decrUser(userid, cfg.RKEY_USER_DIAMOND2, d)
}

//增加用户的金币
func INCRUserCOIN(userid uint32, d int64) (int64, error) {
	return incrUser(userid, cfg.RKEY_USER_COIN, d)
}

//减少用户的金币
func DECRUserCOIN(userid uint32, d int64) (int64, error) {
	count := GetUserCoin(userid)
	if count-d < 0 {
		return count, errors.New("余额不足，减少用户金币失败！")
	}
	return decrUser(userid, cfg.RKEY_USER_COIN, d)
}

//减少用户的金币,当玩家金币不足的时候，设置玩家的金币为0
func DECRUserCOINv2(userId uint32, d int64) (int64, error) {
	//todo
	return 0, nil
}

//增加用户奖券
func INCRUserTicket(userid uint32, d int32) (int32, error) {
	ticket := GetUserTicket(userid)
	ticket_new := ticket + d
	user := GetUserById(userid)
	user.Ticket = proto.Int32(ticket_new)
	UpdateUser2Mgo(user)

	return ticket_new, nil
}

//减少用户奖券
func DECUserTicket(userid uint32, d int32) (int32, error) {
	ticket := GetUserTicket(userid)
	ticket_new := ticket - d
	if ticket_new < 0 {
		return ticket, errors.New("奖券余额不足")
	}
	user := GetUserById(userid)
	user.Ticket = proto.Int32(ticket_new)
	UpdateUser2Mgo(user)

	return ticket_new, nil
}

//增加用户红包
func INCRUserBonus(userid uint32, d float64) (float64, error) {
	bonus := GetUserBonus(userid)
	bonus_new := bonus + d
	//保留有效小数
	bonus_new = float64(int64(bonus_new*100)) / 100

	user := GetUserById(userid)
	user.Bonus = proto.Float64(bonus_new)
	UpdateUser2Mgo(user)

	return bonus_new, nil
}

//减少用户红包
func DECUserBonus(userid uint32, d float64) (float64, error) {
	bonus := GetUserBonus(userid)
	bonus_new := bonus - d
	//保留有效小数
	bonus_new = float64(int64(bonus_new*100)) / 100

	if bonus_new < 0 {
		return bonus, errors.New("红包余额不足")
	}
	user := GetUserById(userid)
	user.Bonus = proto.Float64(bonus_new)
	UpdateUser2Mgo(user)

	return bonus_new, nil
}
