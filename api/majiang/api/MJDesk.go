package api

//麻将桌子的定义
type MJDesk interface {
	EnterUser(userId uint32) error
	ActOut(userId uint32, p interface{}) error //出牌的user和牌型
	ActReady(userId uint32) error
	GetDeskId() int32 //得到desk id
}

type MJDeskCore struct {
	s     *module.Skeleton
	users []MJUser
}

func NewMJDeskCore(s *module.Skeleton) *MJDeskCore {
	return &MJDeskCore{
		s: s,
	}
}
