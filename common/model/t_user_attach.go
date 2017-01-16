package model

type T_user_attach struct {
	UserId              uint32 //用户id
	GameId              int32  //游戏编号 暂不使用
	LastDrawLotteryTime string //上一次转盘奖励领取时间
	LastSignTime        string //上一次签到的时间
	SignTotalDays       int32  //总共签到天数 
	SignContinuousDays  int32  //连续签到天数
	AllowanceTimes      int32  //今日补助领取次数
	LastAllowanceTime   string //上一次补助领取时间
}