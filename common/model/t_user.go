package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)


/**

	warn 需要注意的事情:
	model中比较ObjectId
 */

type T_user struct {
	Mid               bson.ObjectId                `json:"mid" bson:"_id"`
	Id                uint32    //id
	NickName          string    //昵称
	Location          string    //地区
	Password          string

	Coin              int64     //金币
	Diamond           int64     //钻石
	Scores            int32     //积分

	Viplevel          string    //vip的等级
	Level             string    //等级
	Rank              string    //称号

	TimeLastSign      time.Time //最后一次签到时间
	ContinueSignCount int32     //连续签到的次数

				    //微信
	OpenId            string    //openId
	HeadUrl           string    //头像
	Sex               int32
	City              string
}

func (t *T_user) GetMid() bson.ObjectId {
	return t.Mid
}
