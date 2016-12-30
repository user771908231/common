package service

import (
	"time"
	"casino_common/utils/jobUtils"
	"github.com/name5566/leaf/gate"
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
)

//检测器接口
type PayChecker interface {
	OnInit()
	Check()
	IsOK() bool
	IsTimeOut() bool
}

//状态检测器
type WxPayChecker struct {
	InitStatus int32         //初始状态
	OverStatus int32         //支付的结果状态，1，成功，2，用户取消等等
	TimeOut    time.Duration //超时的时间
	agent      gate.Agent    //链接
	tradeNo    string
}

//redis中的支付状态改变了之后，表示支付完成
func (c *WxPayChecker) IsOK() bool {
	return true
}

func (c *WxPayChecker) IsTimeOut() bool {
	return true
}

//初始化检测器的数据
func (c *WxPayChecker) OnInit() {

}

func (c *WxPayChecker) Res(p proto.Message) {
	agent := c.agent
	if agent != nil {
		agent.WriteMsg(p)
	}
}

//开始检测
func (c *WxPayChecker) Check() {
	//每秒钟执行一次
	jobUtils.DoAsynJob(time.Second * 1, func() bool {
		detal := GetDetailsByTradeNo(c.tradeNo)
		if detal.GetStatus() == ddproto.PayEnumTradeStatus_PAY_S_SUCC {
			//得到需要返回的信息
			ack := &ddproto.WxpayAckSyncChecker{
				WxpayId:      proto.Int32(detal.GetPayModelId()),
				Diamond:      proto.Int64(detal.GetDiamond()),
				ChangeDiamond:proto.Int64(detal.GetChangeDiamond()),
				Coin:         proto.Int64(detal.GetCoin()),
				ChangeCoin:   proto.Int64(detal.GetChangeCoin())}
			//恢复app需要的数据
			c.Res(ack)

			//
			DelDetails(c.tradeNo) //保存到数据库之后删除//
			return true
		}

		if c.IsTimeOut() {
			return true
		}

		return false
	})
}

func NewWxPayChecker(tradeNo string, a gate.Agent) *WxPayChecker {
	c := &WxPayChecker{tradeNo:tradeNo, agent:a}
	return c
}

//初始化一个检测
func InitChecker(tradeNo string, a gate.Agent) {
	//new 一个checker
	c := NewWxPayChecker(tradeNo, a)
	c.OnInit()
	c.Check()
}
