package api

import (
	"casino_common/api/majiang"
	"fmt"
	"reflect"
	"casino_common/common/log"
)

//跑得快解析器
type MJParser interface {
	CanPeng(interface{}, interface{}) (bool, error)
	CanGang(interface{}, interface{}) (bool, error)
	CanChi(interface{}, interface{}) (bool, error)
	CanBu(interface{}, interface{}) (bool, error)
	GetJiaoInfos(interface{}) (interface{}, error) //判断是否有叫
	Parse(pids []int32) (interface{}, error)       //通过一副牌的id解析牌型
	XiPai() interface{}                            //洗牌
	Hu(...interface{}) (interface{}, error)        //胡牌的方式...
}

//麻将的骨架，通用的方法都在这里
type MJParserCore struct {
}

//统计牌 27这个谁需要考虑 东南西北发中白的情况？
func (p *MJParserCore) countHandPais(pais []*majiang.MJPAI) []int {
	counts := make([]int, 27) //0~27
	for _, p := range pais {
		counts[p.GetCountIndex() ] ++
	}
	return counts
}

//是否能碰牌
func (p *MJParserCore) CanPeng(userGameData interface{}, pengPai interface{}) (bool, error) {
	log.T("判断是否能碰:userGameData.type%v,content:%v \n", reflect.TypeOf(userGameData), userGameData)
	gameData := userGameData.(MJUserGameData) //玩家数据
	handPais := gameData.GetHandPais()        //判断的手牌
	pengPai2 := pengPai.(*majiang.MJPAI)      //判断的碰牌

	paiCount := 0
	for _, p := range handPais {
		if p.Value == pengPai2.Value && p.Flower == pengPai2.Flower {
			paiCount++
		}
	}
	return paiCount >= 2, nil
}

//是否能杠牌
/**
	判断别人打的牌是否能杠(这里都是点杠)
 */
func (p *MJParserCore) CanGang(userGameData interface{}, gangPai interface{}) (bool, error) {
	log.T("判断是否能杠:userGameData.type%v,content:%v \n", reflect.TypeOf(userGameData), userGameData)
	gameData := userGameData.(MJUserGameData) //玩家数据
	handPais := gameData.GetHandPais()        //判断的手牌
	gangPai2 := gangPai.(*majiang.MJPAI)      //判断的碰牌

	paiCount := 0
	for _, p := range handPais {
		fmt.Printf("得到的碰牌的信息:%v", p)
		if p.Value == gangPai2.Value && p.Flower == gangPai2.Flower {
			paiCount++
		}
	}
	return paiCount >= 2, nil
}

/**
	自己摸的牌，有哪些可以杠的牌
 */

var (
	GANGTYPE_MING int32 = 1
	GANGTYPE_BA   int32 = 2
	GANGTYPE_AN   int32 = 3
)

type ZimoGangCards struct {
	GangPai  *majiang.MJPAI
	GangType int32 //杠牌的类型
}

func (p *MJParserCore) ZiMoGangCards(userGameData interface{}, gangPai interface{}) (interface{}, error) {
	gameData := userGameData.(MJUserGameData) //玩家数据
	handPais := gameData.GetHandPais()        //判断的手牌
	pengPais := gameData.GetPengPais()        //得到所有的吃牌
	var ret []*ZimoGangCards
	//首先判断明杠
	counts := p.countHandPais(handPais) //统计牌
	for _, p := range handPais {
		if 4 == counts[ p.GetCountIndex()] {
			r := &ZimoGangCards{
				GangPai:  p,
				GangType: GANGTYPE_MING,
			}
			ret = append(ret, r)
		}
	}

	//再判断巴杠
	for _, p := range handPais {
		for _, peng := range pengPais {
			pengPai := peng.Pais[0] //判断碰牌的第一张和手牌是否一致
			if p.Flower == pengPai.Flower &&
				p.Value == pengPai.Value {
				r := &ZimoGangCards{
					GangPai:  p,
					GangType: GANGTYPE_BA,
				}
				ret = append(ret, r)
				break
			}
		}
	}

	//清楚重复的数据
	var ret2 []*ZimoGangCards
	isexistRet2 := func(g *ZimoGangCards) bool {
		for _, r := range ret2 {
			if r.GangPai.Flower == g.GangPai.Flower &&
				r.GangPai.Value == g.GangPai.Value {
				return true
			}
		}
		return false
	}

	for _, g := range ret {
		//首先判断ret2中是否有杠牌了
		if !isexistRet2(g) {
			ret2 = append(ret2, g)
		}
	}
	//返回最后的杠牌的数据
	return ret2, nil
}

//是否可以吃牌
func (p *MJParserCore) CanChi(interface{}, interface{}) (bool, error) {
	return false, nil

}

//得到叫牌的信息
func (p *MJParserCore) GetJiaoInfos(interface{}) (interface{}, error) {
	return nil, nil
}

//通过一副牌的id解析牌型
func (p *MJParserCore) Parse(pids []int32) (interface{}, error) {
	return nil, nil
}

//洗牌
func (p *MJParserCore) XiPai() interface{} {
	return nil
}
