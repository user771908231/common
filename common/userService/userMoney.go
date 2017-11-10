package userService

import (
	"casino_common/common/Error"
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/model"
	"casino_common/common/service/tradeLogService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/redisUtils"
	"errors"
	"fmt"
	"time"
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
	return SyncReidsUserMoney(user)
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
	return GetUserMoney(userId, consts.RKEY_USER_DIAMOND)
}

func SetUserDiamond(userId uint32, diamond int64) int64 {
	SetUserMoney(userId, consts.RKEY_USER_DIAMOND, diamond)
	return diamond
}

//获取用户的钻石
func GetUserDiamond2(userId uint32) int64 {
	return GetUserMoney(userId, consts.RKEY_USER_DIAMOND2)
}

//获取用户的房卡
func GetUserRoomCard(userId uint32) int64 {
	return GetUserMoney(userId, consts.RKEY_USER_ROOMCARD)
}

func SetUserRoomCard(userId uint32, roomcard int64) int64 {
	SetUserMoney(userId, consts.RKEY_USER_ROOMCARD, roomcard)
	return roomcard
}

//获取用户的金币
func GetUserCoin(userId uint32) int64 {
	return GetUserMoney(userId, consts.RKEY_USER_COIN)
}

func SetUserCoin(userId uint32, coin int64) int64 {
	SetUserMoney(userId, consts.RKEY_USER_COIN, coin)
	return coin
}

//获取用户奖券
func GetUserTicket(userId uint32) int32 {
	return GetUserById(userId).GetTicket()
}

//获取用户红包
func GetUserBonus(userId uint32) float64 {
	return GetUserById(userId).GetBonus()
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
	log.T("用户账户变动[增加][%v]: user[%v] 值:[%v]", key, userid, d)
	//1,增加余额
	remain, err := redisUtils.INCRBY(redisUtils.K(key, userid), d)
	if err != nil {
		return remain, err
	}
	//2,更新redis和数据库中的数据
	err = UpdateRedisUserMoney(userid) //增加用户货币之后
	//3,返回值
	log.T("用户账户变动[增加][%v]: user[%v] 值:[%v] 增加之后redis的值:%v", key, userid, d, remain)
	return remain, err
}

//减少用户的货币
func decrUser(userid uint32, key string, d int64) (int64, error) {
	log.T("用户账户变动[减少][%v]: user[%v] 值:[%v]", key, userid, d)
	remain, err := redisUtils.DECRBY(redisUtils.K(key, userid), d)
	if err != nil {
		return remain, err
	}
	if remain < 0 {
		old, _ := incrUser(userid, key, d)
		errMsg := fmt.Sprintf("用户[%v]的key[%v]值为[%v]不足[%v],减少的时候失败", userid, key, old, d)
		log.E(errMsg)
		log.T("用户账户变动[减少][%v]: user[%v] 值:[%v] 减少之后redis的值:%v,备注:不够的情况", key, userid, d, remain)

		return remain, errors.New(errMsg)
	} else {
		//更新redis和mongo中的数据
		err = UpdateRedisUserMoney(userid)
		log.T("用户账户变动[减少][%v]: user[%v] 值:[%v] 减少之后redis的值:%v", key, userid, d, remain)

		return remain, err
	}
}

//增加用户的钻石
func INCRUserDiamond(userid uint32, d int64, remark string) (int64, error) {
	new_diamond, err := incrUser(userid, consts.RKEY_USER_DIAMOND, d)
	if err == nil {
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_DIAMOND, float64(d), float64(new_diamond), remark)
	}
	return new_diamond, err
}

//减少用户的砖石
func DECRUserDiamond(userid uint32, d int64, remark string) (int64, error) {
	count := GetUserDiamond(userid)
	if count-d < 0 {
		return count, errors.New("余额不足，减少钻石失败！")
	}

	new_diamond, err := decrUser(userid, consts.RKEY_USER_DIAMOND, d)
	if err == nil {
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_DIAMOND, float64(-d), float64(new_diamond), remark)
	}
	return new_diamond, err
}

//增加用户的房卡 参数:gid:游戏id，memo :说明
func INCRUserRoomcard(userId uint32, d int64, gid int32, remark string) (int64, error) {
	roomCardBefore := GetUserRoomCard(userId)

	//log.T("INCRUserRoomcard(userId=%d,d=%d,gid=%d,memo=%s) before:%d", userId, d, gid, remark, GetUserRoomCard(userId))
	roomCard, err := incrUser(userId, consts.RKEY_USER_ROOMCARD, d)
	//log.T("INCRUserRoomcard(userId=%d,d=%d,gid=%d,memo=%s) after:%d", userId, d, gid, remark, GetUserRoomCard(userId))
	if err != nil {
		return roomCard, err
	}

	if roomCard != roomCardBefore+d {
		err = errors.New("inc roomcard validate fail.")
		return roomCard, err
	}

	//增加扣除房卡的记录
	tradeLogService.Add(userId, ddproto.HallEnumTradeType_PROPS_FANGKA, float64(d), float64(roomCard), remark)

	return roomCard, err
}

//减少用户的房卡
//参数说明:gid:游戏id, remark:备注
func DECRUserRoomcard(userId uint32, d int64, gid int32, remark string) (int64, error) {
	roomCardBefore := GetUserRoomCard(userId)
	if roomCardBefore-d < 0 {
		return roomCardBefore, errors.New("余额不足，减少房卡失败！")
	}

	roomCard, err := decrUser(userId, consts.RKEY_USER_ROOMCARD, d)

	if roomCard != roomCardBefore-d {
		err = errors.New("dec roomcard validate fail.")
		return roomCard, err
	}

	//增加扣除房卡的记录
	tradeLogService.Add(userId, ddproto.HallEnumTradeType_PROPS_FANGKA, float64(-d), float64(roomCard), remark)

	return roomCard, err
}

//增加用户的金币
func INCRUserCOIN(userid uint32, d int64, remark string) (int64, error) {
	new_coin, err := incrUser(userid, consts.RKEY_USER_COIN, d)
	if err == nil {
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_COIN, float64(d), float64(new_coin), remark)
	}
	return new_coin, err
}

//减少用户的金币
func DECRUserCOIN(userid uint32, d int64, remark string) (int64, error) {
	count := GetUserCoin(userid)
	if count-d < 0 {
		return count, errors.New("余额不足，减少用户金币失败！")
	}
	new_coin, err := decrUser(userid, consts.RKEY_USER_COIN, d)
	if err == nil {
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_COIN, float64(-d), float64(new_coin), remark)
	}
	return new_coin, err
}

//增加用户奖券
func INCRUserTicket(userid uint32, d int32, remark string) (int32, error) {
	if user := GetUserById(userid); user != nil {
		user.Ticket = proto.Int32(user.GetTicket() + d)
		err := UpdateUser2Mgo(user)
		if err != nil {
			return user.GetTicket() - d, err
		}
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_TICKET, float64(d), float64(user.GetTicket()), remark)
		return user.GetTicket(), nil
	}
	return 0, errors.New("user not found")
}

//减少用户奖券
func DECUserTicket(userid uint32, d int32, remark string) (int32, error) {
	if user := GetUserById(userid); user != nil {
		user.Ticket = proto.Int32(user.GetTicket() - d)
		err := UpdateUser2Mgo(user)
		if err != nil {
			return user.GetTicket() + d, err
		}
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_TICKET, float64(-d), float64(user.GetTicket()), remark)
		return user.GetTicket(), nil
	}
	return 0, errors.New("user not found")
}

//增加用户红包
func INCRUserBonus(userid uint32, d float64, remark string) (float64, error) {
	if user := GetUserById(userid); user != nil {
		user.Bonus = proto.Float64(user.GetBonus() + d)
		err := UpdateUser2Mgo(user)
		if err != nil {
			return user.GetBonus() - d, err
		}
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_BONUS, float64(d), float64(user.GetBonus()), remark)
		return user.GetBonus(), nil
	}
	return 0, errors.New("user not found")
}

//减少用户红包
func DECUserBonus(userid uint32, d float64, remark string) (float64, error) {
	if user := GetUserById(userid); user != nil {
		user.Bonus = proto.Float64(user.GetBonus() - d)
		err := UpdateUser2Mgo(user)
		if err != nil {
			return user.GetBonus() + d, err
		}
		tradeLogService.Add(userid, ddproto.HallEnumTradeType_TRADE_BONUS, float64(-d), float64(user.GetBonus()), remark)
		return user.GetBonus(), nil
	}
	return 0, errors.New("user not found")
}
