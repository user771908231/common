package majiang

type TMjRoomConfig struct {
	RoomName        string //房间的名称
	RoomLevel       int32  //房间的登记
	RoomBaseValue   int64  //低分
	RoomLimitCoin   int64  //准入的金额
	RoomLimitCoinUL int64  //准入的金额上限（超过这个金额之后，不允许再进入房间）
	EnterCoinFee    int64  //准入的金币限制
	GameType        int32  //玩法
}
