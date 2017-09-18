package phzData

import (
	"casino_common/common/log"
	"casino_common/proto/ddproto"
)

type PHZParser interface {
	CanTi(...interface{}) (interface{}, error)   //根据桌面的牌判断手牌里三张的牌是否可以跑
	GetTiPai(interface{}) interface{}            //获取手牌里原始可以提的牌
	CanPeng(...interface{}) (interface{}, error) //根据桌面的牌判断手牌里是否有可以碰的牌
	CanChi(...interface{}) (interface{}, error)  //根据桌面的牌判断手牌里是否有可以吃的牌型
	TryHu(...interface{}) (interface{}, error)   //根据桌面的牌判断是否可以胡牌
	CheckHu(interface{}) (bool, error)           //校验能否天胡、地胡
}

type ParserCore struct {
}

type HuInfo struct {
	WinUser    uint32 //
	HuXi       int32  //胡息数
	LoseUser   uint32
	IsZimo     bool
	GetUserId  uint32
	SendUserId uint32
	HuDesc     string
	Fan        int32
	Score      int64
	Pai        *PHZPoker
	HuType     int32
	PaiType    []int32 //牌的类型
}

func (p *ParserCore) CountHandPais(pais []*PHZPoker) []int {
	counts := make([]int, 21) //牌值1——10:小字  11——20:大字
	for _, p := range pais {
		if p.GetPaiIndexByValue() == 0 {
			log.T("CountHandPais occur error...p.value==0")
			return nil
		}
		counts[p.GetPaiIndexByValue()]++
	}
	log.T("测试用...counts=%v", counts)
	return counts
}

func (p *ParserCore) GetTiPai(data interface{}) interface{} {
	tiInfo := []*TiPai{}
	handPokers := data.([]*PHZPoker)
	paiCounts := p.CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			if paiCount == 4 {
				ti := &TiPai{}
				ti.TiType = int32(ddproto.PhzEnumTiType_PHZ_TITYPE_HAVE_FOUR)
				if paiValue <= 10 {
					log.T("提小字，胡息数为9")
					ti.HuXi = 9 //提小字9胡息
				} else {
					log.T("提大字，胡息数为12")
					ti.HuXi = 12 //提大字12胡息
				}
				//将提的牌放进提牌结构里
				for _, pai := range handPokers {
					if pai.GetValue() == int32(paiValue) {
						ti.Pais = append(ti.Pais, pai)
					}
				}
				tiInfo = append(tiInfo, ti)
			}
		}
		return tiInfo
	}
	return nil
}

func (p *ParserCore) CanTi(data interface{}, tiData interface{}) (interface{}, error) {
	ti := &TiPai{}
	handPokers := data.([]*PHZPoker)
	tiPai := tiData.(*PHZPoker)
	paiCounts := p.CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			if paiCount == 3 && paiValue == int(tiPai.GetValue()) {
				if paiValue <= 10 {
					log.T("跑小字，胡息数9")
					ti.HuXi = 9 //提小字9胡息
				} else {
					log.T("跑大字，胡息数12")
					ti.HuXi = 12 //提小字12胡息
				}
				for _, pai := range handPokers {
					if pai.GetValue() == int32(paiValue) {
						ti.Pais = append(ti.Pais, pai)
					}
				}
				return ti, nil
			}
		}
	}
	return nil, nil
}

func (p *ParserCore) CanPeng(userGameData interface{}, pengData interface{}) (interface{}, error) {
	handPokers := userGameData.([]*PHZPoker)
	pengPai := pengData.(*PHZPoker)
	pengReslut := &PengPai{}
	paiCounts := p.CountHandPais(handPokers)
	pengReslut.Pais = append(pengReslut.Pais, pengPai) //将桌面内的碰牌也存进去
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			log.T("测试用...paiValue=[%v]  paiCount=[%v] pengPai.value=[%v]", paiValue, paiCount, pengPai.GetValue())
			if paiValue == int(pengPai.GetValue()) && paiCount == 2 {
				log.T("可以碰牌...碰的牌是：[%v]  pengPai.value=[%v]   paiValue= [%v]", Card2String(pengPai), pengPai.GetValue(), paiValue)
				for _, pai := range handPokers {
					if pai.GetValue() == int32(paiValue) {
						pengReslut.Pais = append(pengReslut.Pais, pai)
						if paiValue <= 10 {
							log.T("碰小字，胡息数为：[%v]", 1)
							pengReslut.HuXi = 1 //碰小字1胡息
						} else {
							log.T("碰大字，胡息数为：[%v]", 3)
							pengReslut.HuXi = 3 //碰大字3胡息
						}
					}
				}
				return pengReslut, nil
			}
		}
	}
	return nil, nil
}

func (p *ParserCore) CanChi(userGameData interface{}, chiData interface{}) (interface{}, error) {
	handPokers := userGameData.([]*PHZPoker)
	chiPoker := chiData.(*PHZPoker)
	var chiResult []*ChiPai
	var pai1, pai2 *PHZPoker
	//带二七十的吃牌
	if chiPoker.GetValue() == 2 {
		//1 2 3
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 3
			chiResult = append(chiResult, chiInfo)
		}

		//2 3 4
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+2)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//2 7 10
		pai1 = p.GetPaiByValue(handPokers, 7)
		pai2 = p.GetPaiByValue(handPokers, 10)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 3
			chiResult = append(chiResult, chiInfo)
		}
	} else if chiPoker.GetValue() == 7 {
		//5 6 7
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-2)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//6 7 8
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//7 8 9
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+2)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//2 7 10
		pai1 = p.GetPaiByValue(handPokers, 2)
		pai2 = p.GetPaiByValue(handPokers, 10)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 3
			chiResult = append(chiResult, chiInfo)
		}
	} else if chiPoker.GetValue() == 10 {
		//8 9 10
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-2)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//2 7 10
		pai1 = p.GetPaiByValue(handPokers, 2)
		pai2 = p.GetPaiByValue(handPokers, 7)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 3
			chiResult = append(chiResult, chiInfo)
		}
	} else if chiPoker.GetValue() == 12 {
		//壹 贰 叁
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 6
			chiResult = append(chiResult, chiInfo)
		}

		//贰 叁 肆
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+2)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//2 7 10
		pai1 = p.GetPaiByValue(handPokers, 17)
		pai2 = p.GetPaiByValue(handPokers, 20)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 6
			chiResult = append(chiResult, chiInfo)
		}
	} else if chiPoker.GetValue() == 17 {
		//伍 陆 柒
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-2)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//陆 柒 玐
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//柒 玐 玖
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+2)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//贰 柒 拾
		pai1 = p.GetPaiByValue(handPokers, 12)
		pai2 = p.GetPaiByValue(handPokers, 20)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 6
			chiResult = append(chiResult, chiInfo)
		}
	} else if chiPoker.GetValue() == 20 {
		//玐 玖 拾
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-2)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		//贰 柒 拾
		pai1 = p.GetPaiByValue(handPokers, 2)
		pai2 = p.GetPaiByValue(handPokers, 7)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 6
			chiResult = append(chiResult, chiInfo)
		}
	} else {
		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+10)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue())
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+10)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+10)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-10)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-10)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-2)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}

		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+2)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiInfo.HuXi = 0
			chiResult = append(chiResult, chiInfo)
		}
	}

	//牌值连续的三张牌
	if chiPoker.GetValue() != 10 {
		pai1 := p.GetPaiByValue(handPokers, chiPoker.GetValue()-2)
		pai2 := p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		if pai1 != nil && pai2 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiResult = append(chiResult, chiInfo)
		}

		pai1 = p.GetPaiByValue(handPokers, chiPoker.GetValue()-1)
		pai2 = p.GetPaiByValue(handPokers, chiPoker.GetValue()+1)
		if pai2 != nil && pai1 != nil {
			chiInfo := &ChiPai{}
			chiInfo.Pai = chiPoker
			chiInfo.Pais = append(chiInfo.Pais, pai1)
			chiInfo.Pais = append(chiInfo.Pais, pai2)
			chiInfo.Pais = append(chiInfo.Pais, chiPoker)
			chiResult = append(chiResult, chiInfo)
		}
	}
	return chiResult, nil
}

func IsExitInArray(key int32, array []int32) bool {
	for _, value := range array {
		if key == value {
			return true
		}
	}
	return false
}

//根据牌值在手牌中找相应的poker
func (p *ParserCore) GetPaiByValue(handPokers []*PHZPoker, value int32) *PHZPoker {
	for _, poker := range handPokers {
		if poker.GetValue() == value {
			return poker
		}
	}
	return nil
}

//校验玩家是否天胡：砌完牌后，玩家手中有3提或5坎，称为天胡
func (p *ParserCore) CheckHu(userGameData interface{}) (bool, error) {
	tiCount := make(map[int]int32)
	kanCount := make(map[int]int32)
	handPokers := userGameData.([]*PHZPoker)
	paiCounts := p.CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			if paiCount == 4 {
				tiCount[paiValue]++
			}
			if paiCount == 3 {
				kanCount[paiValue]++
			}
		}
	}
	if len(tiCount) >= 3 {
		return true, nil
	}
	if len(kanCount) >= 5 {
		return true, nil
	}
	if len(tiCount) < 3 && len(kanCount) < 5 && (len(tiCount)+len(kanCount)) >= 5 {
		return true, nil
	}
	return false, nil
}

func (p *ParserCore) TryHu(userGameData interface{}, pai interface{}) (interface{}, error) {
	/*handPokers := userGameData.([]*PHZPoker)
	checkPai := pai.(*PHZPoker)*/
	return nil, nil
}
