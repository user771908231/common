package api

type MJRoomMgr interface {
	OnInit() error                                  //初始化room管理器
	GetRoom(...interface{}) (MJRoom, error)         //通过room类型和level等级获得一个room
	GetDeskBySession(userId uint32) (MJDesk, error) //通过玩家的session找到一个desk
}
