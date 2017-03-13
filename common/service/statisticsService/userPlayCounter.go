package statisticsService

/**
	每一个游戏玩家可以new一个UserPlayCounter,只需要在牌局结束的时候，调用方法就行了...
 */
type UserPlayCounter struct {
	gamId     int32  //游戏id
	userId    uint32 //玩家id
	roomType  int32  //房间类型
	roomLevel int32  //房间等级
}

//一个针对user的局数统计器
func NewUserPlayCounter(gameId int32, roomType int32, roomLevel int32) *UserPlayCounter {
	return &UserPlayCounter{
		gamId:     gameId,
		roomType:  roomType,
		roomLevel: roomLevel,
	}
}

//玩家的局数统计
func (c *UserPlayCounter) IncPlayCount(win bool) (totalPlayCount int32, winPlayCount int32, e error) {
	return c.incPlayCount(c.userId, c.gamId, c.roomType, c.roomLevel, win)
}

//todo 实现计数器的方法
func (c *UserPlayCounter) incPlayCount(userId uint32, gameId int32, roomType int32, roomLevel int32, win bool) (totalPlayCount int32, winPlayCount int32, e error) {
	return 0, 0, nil
}
