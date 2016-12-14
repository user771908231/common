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
	Location          string    //地区
	Password          string    //用户密码，现在默认没有使用这个字段，因为用的是微信登录

	Coin              int64     //金币-金币场使用，朋友桌没有金币存储，游戏结束之后，金币置空
	Diamond           int64     //钻石-金币场使用
	Diamond2          int64     //钻石-朋友桌使用==房卡

	Scores            int32     //积分

	Viplevel          string    //vip的等级
	Level             string    //等级
	Rank              string    //称号

	TimeLastSign      time.Time //最后一次签到时间
	ContinueSignCount int32     //连续签到的次数

				    //微信
	UnionId           string    //多个openId 对应一个UnionId
	OpenId            string    //openId
	NickName          string    //昵称
	HeadUrl           string    //头像地址
	Sex               int32     //性别	1，男，2女
	City              string    //城市
}

//继承baseModel
func (t *T_user) GetMid() bson.ObjectId {
	return t.Mid
}
