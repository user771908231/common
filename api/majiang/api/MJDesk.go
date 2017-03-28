package api

import "github.com/name5566/leaf/module"

//麻将桌子的定义
type MJDesk interface {
	EnterUser(userId uint32) error
	ActOut(userId uint32, p interface{}) error //出牌的user和牌型
	ActReady(userId uint32) error
	GetDeskId() int32 //得到desk id
}

type MJDeskCore struct {
	s      *module.Skeleton
	deskId int32
	parser MJParser
	users  []MJUser
}

func NewMJDeskCore(s *module.Skeleton) *MJDeskCore {
	return &MJDeskCore{
		s: s,
	}
}

//胡牌的 解析器
func (d *MJDeskCore) GetParser() MJParser {
	return d.parser
}

func (d *MJDeskCore) GetDeskId() int32 {
	return d.deskId
}

func (d *MJDeskCore) GetUserById(userId uint32) MJUser {
	for _, u := range d.users {
		if u != nil && u.GetUserId() == userId {
			return u
		}
	}
	return nil
}

func (d *MJDeskCore) GetUsers() []MJUser {
	return d.users
}
