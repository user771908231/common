package award

//奖励的上下文
type Context interface {
	DoAward()
}

//计时奖励
type TimeCtx struct {
	data Model
}
