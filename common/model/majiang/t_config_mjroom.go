package majiang

type TMjRoomConfig struct {
	RoomName        string //房间的名称
	RoomLevel       int32  //房间的等级
	RoomBaseValue   int64  //低分
	RoomLimitCoin   int64  //准入的金额
	RoomLimitCoinUL int64  //准入的金额上限（超过这个金额之后，不允许再进入房间）
	EnterCoinFee    int64  //准入的金币限制
	GameType        int32  //玩法
}

//转转麻将桌的配置信息结构
type TMjZHZHDeskConfig struct {
	UserCountLimit int32 //人数
	BaseValue      int64 //低分
	LimitCoin      int64 //准入的金额
	LimitCoinUL    int64 //准入的金额上限（超过这个金额之后，不允许再进入房间）
	EnterCoinFee   int64 //准入的金币限制
	ZhaMa          int32 //扎多少码 1为一码全中
	IsZhaMaJiaBei  bool  //扎码是否加倍
	IsBaHongZhong  bool  //是否是8红中 红中使用
	GameType       int32 //玩法
}
