package service

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"casino_common/utils/redisUtils"
	"errors"
	"fmt"
	"casino_common/common/log"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"strings"
)

//得到同步机制的key
func GetTradeSyncKey(tradeNo string) string {
	return strings.Join([]string{"sync", tradeNo}, "_")
}

func InitTradeSyncStatus(tradeNo string) {
	redisUtils.SetInt64(GetTradeSyncKey(tradeNo), int64(ddproto.PayEnumTradeStatus_PAY_S_WATINGPAY))
}

//end_status
func EndTradeSyncStatus(tradeNo string) error {
	s := redisUtils.INCRBY(GetTradeSyncKey(tradeNo), 1)
	if s != int64(ddproto.PayEnumTradeStatus_PAY_S_SUCC) {
		return ERR_ENDTRADE_REPETITION //重复回调
	} else {
		return nil
	}
}

//得到一个支付明细
func NewAndSavePayDetails(userId uint32, mealId int32, payModelId int32, tradeNo string) (*ddproto.PayBaseDetails, error) {
	ret := &ddproto.PayBaseDetails{
		UserId:    proto.Uint32(userId),
		MealId:    proto.Int32(mealId),
		PayModelId:proto.Int32(payModelId),
		TradeNo:   proto.String(tradeNo),
		Status:    ddproto.PayEnumTradeStatus_PAY_S_UNIFIEDORDER.Enum()}

	//保存明细
	SaveDetails(ret)
	//设置一个订单准备，用于同步，避免重复微信回调时候的重复回调
	return ret, nil
}

//保存订单
func SaveDetails(d *ddproto.PayBaseDetails) {
	//保存
	redisUtils.SetObj(d.GetTradeNo(), d)
}

func SaveDetails2Mgo(d *ddproto.PayBaseDetails) {
	db.InsertMgoData(tableName.DBT_T_RECHARGE_DETAILS, d)
}

func GetDetailsByTradeNo(tradeNo string) *ddproto.PayBaseDetails {
	redata := redisUtils.GetObj(tradeNo, &ddproto.PayBaseDetails{})
	if redata != nil {
		return redata.(*ddproto.PayBaseDetails)
	}

	//没有找到数据
	return nil
}

//更新订单的状态
func UpdateDetailsStatus(tradeNo string, s ddproto.PayEnumTradeStatus) error {
	detail := GetDetailsByTradeNo(tradeNo)
	if detail == nil {
		errMsg := fmt.Sprintf("没有找到对应[%v]的订单信息.", tradeNo)
		log.E(errMsg)
		return errors.New(errMsg)
	}
	//更新状态
	detail.Status = s.Enum()
	redisUtils.SetObj(tradeNo, detail)
	return nil
}
