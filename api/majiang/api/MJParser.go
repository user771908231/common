package api

import (
	"casino_common/api/majiang"
	"fmt"
	"reflect"
)

//跑得快解析器
type MJParser interface {
	//出牌的时候比较大小
	CanPeng(interface{}, interface{}) (bool, error)
	CanGang(interface{}, interface{}) (bool, error)
	CanChi(interface{}, interface{}) (bool, error)
	CanBu(interface{}, interface{}) (bool, error)
	//通过一副牌的id解析牌型
	Parse(pids []int32) (interface{}, error)
	//洗一副扑克牌出来
	XiPai() interface{}                     //洗牌
	Hu(...interface{}) (interface{}, error) //胡牌的方式...
}

//麻将的骨架，通用的方法都在这里
type MJParserCore struct {
}

//是否能碰牌
func (p *MJParserCore) CanPeng(userGameData interface{}, pengPai interface{}) (bool, error) {
	fmt.Printf("得到的userGameData.type%v \n", reflect.TypeOf(userGameData))

	gameData := userGameData.(MJUserGameData)
	handPais := gameData.GetHandPais()
	pengPai2 := pengPai.(*majiang.MJPAI)
	paiCount := 0

	for _, p := range handPais {
		fmt.Printf("得到的碰牌的信息:%v", p)
		if p.Value == pengPai2.Value && p.Flower == pengPai2.Flower {
			paiCount++
		}
	}
	return paiCount >= 2, nil
}
