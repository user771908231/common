package reward

//奖励的上下文
type Context struct {
	UserId     uint32 //获得奖励的用户
	GameId     int32  //游戏
	RewardType int32  //奖励类型
}
