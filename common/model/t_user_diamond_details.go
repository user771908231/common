package model

import "time"


//用户钻石的消费记录,钻石的消费记录,需要每次都保存到数据库中,
//用户的coin可以不用每次都保存,只需要在每局结束的时候生成战绩即可


var T_USER_DIAMOND_DETAILS_TYPE_CREATEDESK int32 = 2
var T_USER_DIAMOND_DETAILS_TYPE_REBUY int32 = 3
var T_USER_DIAMOND_DETAILS_TYPE_CSTH_DIAMON_REWARD int32 = 4        //锦标赛钻石奖励
var T_USER_DIAMOND_DETAILS_TYPE_DISSOVEDESK int32 = 5 //游戏还没有开始的时候解散房间需要退回钻石
var T_USER_DIAMOND_DETAILS_TYPE_CSFEE int32 = 6 //游戏还没有开始的时候解散房间需要退回钻石


type T_user_diamond_details struct {
	Id            uint32    //id
	UserId        uint32
	Diamond       int64     //收入或者支出的钻石数量
	ReaminDiamond int64     //剩余的钻石
	DetailsType   int32     //交易类型
	Memo          string    //交易备注
	DetailTime    time.Time //交易的时间
}

