package service

import (
	"github.com/chanxuehong/wechat.v2/mch/pay"
	"github.com/chanxuehong/wechat.v2/mch/core"
	"time"
	"casino_common/common/log"
	"fmt"
	"casino_common/proto/ddproto"
	"github.com/chanxuehong/rand"
	"github.com/golang/protobuf/proto"
	"errors"
)

const (
	PACKEGE_INFO string = "Sign=WXPay"
)

func UnifiedOrder(totalFee int64, ip string, tradeNo string, appid, mchid, apikey, deviceInfo string) (*pay.UnifiedOrderResponse, error) {
	client := core.NewClient(appid, mchid, apikey, nil)
	//app支付 需要这么多就行了
	unifiedOrderRequest := new(pay.UnifiedOrderRequest)
	unifiedOrderRequest.DeviceInfo = deviceInfo
	unifiedOrderRequest.Body = "神经棋牌-游戏充值"
	unifiedOrderRequest.OutTradeNo = tradeNo //商户内部的订单号 test ：time+userId+type+rand ?
	unifiedOrderRequest.FeeType = "CNY"
	unifiedOrderRequest.TotalFee = totalFee
	unifiedOrderRequest.SpbillCreateIP = ip
	tnow := time.Now()
	unifiedOrderRequest.TimeStart = core.FormatTime(tnow)                         //订单的创建时间
	unifiedOrderRequest.TimeExpire = core.FormatTime(tnow.Add(time.Minute * 600)) //订单失效的时间
	unifiedOrderRequest.NotifyURL = "https://www.dongzhile.cn/"                   //回调的地址
	unifiedOrderRequest.TradeType = "APP"

	result, err := pay.UnifiedOrder2(client, unifiedOrderRequest)
	if err != nil {
		log.E("统一下单的时候出现错误...err:%v", err)
		return nil, err
	}

	return result, nil
}

//回复信息
func GetAppReqParamsByack(payack *pay.UnifiedOrderResponse, apikey string) (*ddproto.WxpayAckUnifiedorder, error) {
	if payack == nil {
		return nil, errors.New("没有得到统一下单的回复...")
	}

	ret := new(ddproto.WxpayAckUnifiedorder)
	//需要签名的数据
	reqmap := make(map[string]string)
	reqmap["appid"] = payack.AppId     //
	reqmap["partnerid"] = payack.MchId //
	reqmap["prepayid"] = payack.PrepayId
	reqmap["package"] = PACKEGE_INFO
	reqmap["noncestr"] = string(rand.NewHex())
	reqmap["timestamp"] = fmt.Sprint(time.Now().Unix())

	//签名
	retSign := core.Sign(reqmap, apikey, nil)

	//包装app需要的协议
	ret.PartnerId = proto.String(reqmap["partnerid"])
	ret.PrepayId = proto.String(reqmap["prepayid"])
	//ret.Package = proto.String("") //暂时返回空
	ret.NonceStr = proto.String(reqmap["noncestr"])
	ret.TimeStamp = proto.String(reqmap["timestamp"])
	ret.Sign = proto.String(retSign)
	//return
	return ret, nil
}

//得到app需要的支付信息
func GetAppWxpayReqParams(totalFee int64, ip string, tradeNo string, appid, mchid, apikey, deviceInfo string) (*ddproto.WxpayAckUnifiedorder, error) {
	ack, err := UnifiedOrder(totalFee, ip, tradeNo, appid, mchid, apikey, deviceInfo)
	if err != nil {
		log.E("统一下单的时候出现错误:%v", err.Error())
		return nil, err
	}
	//返回结果
	return GetAppReqParamsByack(ack, apikey)
}
