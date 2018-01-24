package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//微信机器人
var WxRobot = SingleRobot{
	AccessToken: "https://oapi.dingtalk.com/robot/send?access_token=b54f83b6ae91c2c16a5dd5c665f3c2c33555af405456abc798c4f1c241ea2c7c",
}

type SingleRobot struct {
	AccessToken string
}

func (r *SingleRobot) SendTxt(msg string) *SendRet {
	var url string = r.AccessToken
	req := bytes.NewBuffer([]byte(fmt.Sprintf(`{"msgtype": "text","text": {"content": "%s"}}`, msg)))
	resp, _ := http.Post(url, body_type, req)
	resp_body, _ := ioutil.ReadAll(resp.Body)

	//
	dr := &SendRet{}
	json.Unmarshal(resp_body, dr)
	return dr
}
