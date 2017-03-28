package handlers

import (
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/gate"
	"gopkg.in/mgo.v2/bson"
	"github.com/golang/protobuf/proto"
	"casino_common/proto/funcsInit"
	"casino_common/utils/redisUtils"
	"casino_common/common/cfg"
	"time"
)

//获取二维码
func HandlerQrLoginGetCode(args []interface{}) {
	//m := args[0].(*ddproto.CommonReqQrLogin)
	a := args[1].(gate.Agent)
	new_code := bson.NewObjectId().Hex()
	msg := ddproto.CommonAckQrLogin{
		Code: proto.String(new_code),
		Url: proto.String("http://wx.tondeen.com/weixin/game/qrlogin?code=" + new_code),
	}
	a.WriteMsg(&msg)
}

//返回微信info
func HandlerQrLoginGetWxInfo(args []interface{}) {
	m := args[0].(*ddproto.CommonReqQrWxInfo)
	a := args[1].(gate.Agent)
	msg := &ddproto.CommonAckQrWxInfo{
		Header: commonNewPorot.NewHeader(),
	}
	code := m.GetCode()

	key := cfg.RKEY_QR_CODE+ "_" +code

	redis_wxinfo := redisUtils.GetObj(key, &ddproto.WeixinInfo{})
	if redis_wxinfo == nil {
		*msg.Header.Code = -1
		*msg.Header.Error = "wx_info not found."
		a.WriteMsg(msg)
		return
	}

	wx_info := redis_wxinfo.(*ddproto.WeixinInfo)
	*msg.Header.Code = 1
	*msg.Header.Error = "获取wx_info成功！"
	msg.Wxinfo = wx_info
	a.WriteMsg(msg)

	//30秒后删除redis中保存的数据
	time.AfterFunc(30 * time.Second, func() {
		redisUtils.Del(key)
	})
}
