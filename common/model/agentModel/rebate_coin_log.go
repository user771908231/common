package agentModel


//金币场代理返利日志
type GamerRebateCoinLog struct {
	UserId uint32       //玩家id
	UserName string     //玩家昵称
	GameId   int32      //所玩游戏
	CoinFee  int64      //房费
	Rebate   int64      //返利
	AgentId  uint32     //领取的代理
	FromAgent uint32    //来自多级返利的代理
	Level    int        //级别
}
