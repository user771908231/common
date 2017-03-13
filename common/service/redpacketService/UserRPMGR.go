package redpacketService

/**
	每一个游戏玩家可以new一个UserPlayCounter,只需要在牌局结束的时候，调用方法就行了...
 */
type UserRPMGR struct {
	gamId     int32  //游戏id
	userId    uint32 //玩家id
	roomType  int32  //房间类型
	roomLevel int32  //房间等级
}

//一个针对user的局数统计器
func NewUserRPMGR(gameId int32, roomType int32, roomLevel int32) *UserRPMGR {
	return &UserRPMGR{
		gamId:     gameId,
		roomType:  roomType,
		roomLevel: roomLevel,
	}
}

//尝试得到红包
func (c *UserRPMGR) TryGetRP() (amount int64, e error) {
	return c.tryGetRP(c.userId, c.gamId, c.roomType, c.roomLevel)
}

//todo
func (c *UserRPMGR) TryGetRP2(winPlayCount int32) (amount int64, e error) {
	return 0, nil
}

//todo 这里去实现玩家是否能够得到红包
func (c *UserRPMGR) tryGetRP(userId uint32, gameId int32, roomType int32, roomLevel int32) (amount int64, e error) {
	return 0, nil
}
