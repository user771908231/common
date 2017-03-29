package api

import (
	"github.com/name5566/leaf/module"
	"casino_common/common/consts/tableName"
	"casino_common/utils/db"
)

//麻将桌子的定义
type MJDesk interface {
	EnterUser(...interface{}) error            //玩家进入desk，不定参数
	ActOut(userId uint32, p interface{}) error //出牌的user和牌型
	ActReady(userId uint32) error              //准备
	GetDeskId() int32                          //得到desk id
	GetPassword() string                       //得到房间号
}

type MJDeskCore struct {
	s        *module.Skeleton
	deskId   int32
	password string //房间号
	parser   MJParser
	users    []MJUser
}

func NewMJDeskCore(s *module.Skeleton) *MJDeskCore {
	desk := &MJDeskCore{
		s: s,
	}
	//main key
	desk.deskId, _ = db.GetNextSeq(tableName.DBT_MJ_DESK)
	return desk
}

//胡牌的 解析器
func (d *MJDeskCore) GetParser() MJParser {
	return d.parser
}

func (d *MJDeskCore) SetParser(p MJParser) {
	d.parser = p
}

func (d *MJDeskCore) GetDeskId() int32 {
	return d.deskId
}

func (d *MJDeskCore) GetPassword() string {
	return d.password
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
