package award

import "github.com/name5566/leaf/gate"

//奖励的上下文
type Context struct {
	UserId     uint32 //获得奖励的用户
	Agent      gate.Agent
	GameId     int32 //游戏
	RewardType int32 //奖励类型
}
