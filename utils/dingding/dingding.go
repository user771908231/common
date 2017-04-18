package utils

import (
	"io/ioutil"
	"bytes"
	"net/http"
	"casino_common/common/log"
	"encoding/json"
	"fmt"
	"sync"
)

var DingMsger *DingMessage

var (
	//获取token
	corpid           string = "dingb02f35b77b75d3c1"
	corpsecret       string = "yohnVsGxdP_DdIuA_fIO33_QEzOfZHJ_vc2b-19KhqshrNbLPCwMBO9cFGG6vxSj"
	body_type        string = "application/json;charset=utf-8"
	URL_GETTOKEN     string = "https://oapi.dingtalk.com/gettoken?corpid=" + corpid + "&corpsecret=" + corpsecret
	URL_SENDMSG      string = "https://oapi.dingtalk.com/chat/send?access_token="
	EC_TOKEN_TIMEOUT int32  = 30014 //token过期的错误吗
	EC_TOKEN         int32  = 40014
	REFRESHCOUNTMAX  int32  = 10 //充实token的次数
)

func init() {
	fmt.Println("开始初始化ddmsger...")
	DingMsger = &DingMessage{
		chatid:  "chat7548ac2ea5f759f5b260a37013263409",
		msgtype: "text",
		msgbody: make(chan *sendBody, 100),
	}
	DingMsger.refrehToken() //刷新tocker
	go DingMsger.start()    //开始等待消息
}

type dingRet struct {
	Errmsg       string `json:"errmsg"`
	Access_token string `json:"access_token"`
	Errcode      int32 `json:"errcode"`
}

//发送消息时候的回复
type dingSendRet struct {
	Errmsg  string `json:"errmsg"`
	Errcode int32 `json:"errcode"`
}

//发送消息的结构
type sendBody struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Text    c `json:"text"`
}

type c struct {
	Content string `json:"content"`
}

func NewMsg(chantId, msgtyp, content string) *sendBody {
	ret := &sendBody{
		Chatid:  chantId,
		Msgtype: msgtyp,
		Text: c{
			Content: content,
		},
	}
	return ret
}

//发送丁丁消息
type DingMessage struct {
	token             string
	chatid            string
	msgtype           string
	msgbody           chan *sendBody
	refreshTokenCount int32
	sync.Mutex
}

func (m *DingMessage) Send(msg string) {
	m.msgbody <- NewMsg(m.chatid, m.msgtype, msg)
}

func (m *DingMessage) send(postParams []byte) {
	var url string = URL_SENDMSG + m.token
	req := bytes.NewBuffer(postParams)
	resp, _ := http.Post(url, body_type, req)
	body2, _ := ioutil.ReadAll(resp.Body)

	//
	dr := &dingSendRet{}
	json.Unmarshal(body2, dr)
	fmt.Printf("ddmsger 发送信息 得到的结果 %+v \n", dr)
	if (dr.Errcode == EC_TOKEN_TIMEOUT || dr.Errcode == EC_TOKEN) &&
		m.refreshTokenCount < REFRESHCOUNTMAX {
		//需要重新获取token
		m.refrehToken()
		m.send(postParams)
		m.refreshTokenCount ++
	} else {
		m.refreshTokenCount = 0
	}
}

//开始接受消息
func (m *DingMessage) start() {
	fmt.Println("ddmsger 开始接受消息...")
	for i := range m.msgbody { // ch关闭时，for循环会自动结束
		b, _ := json.Marshal(i)
		m.send(b)
	}
}

func (m *DingMessage) refrehToken() {
	dr := &dingRet{}
	r, _ := http.Get(URL_GETTOKEN)
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &dr)
	if err != nil {
		log.Printf("解析json的时候err:%v", err)
	}
	m.token = dr.Access_token //给你token赋值
	log.Printf("刷新之后的token :%+v", dr)
}
