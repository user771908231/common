package phzData

import (
	"casino_common/common/Error"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/util"
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

func CountHandPais(pais []*PHZPoker) []int {
	counts := make([]int, 21) //牌值1——10:小字  11——20:大字
	for _, p := range pais {
		if p.GetPaiIndexByValue() == 0 {
			log.T("CountHandPais occur error...p.value==0")
			return nil
		}
		counts[p.GetPaiIndexByValue()]++
	}
	return counts
}

func (p *ParserCore) GetTiPai(data interface{}) interface{} {
	tiInfo := []*TiPai{}
	handPokers := data.([]*PHZPoker)
	paiCounts := CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			//log.T("测试用...paiValue=[%v]		paiCount=[%v]", paiValue, paiCount)
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
	if tiInfo != nil || len(tiInfo) > 0 {
		return tiInfo
	} else {
		return nil
	}
}

func (p *ParserCore) CanTi(data interface{}, tiData interface{}) (interface{}, error) {
	ti := &TiPai{}
	//var ti *TiPai
	handPokers := data.([]*PHZPoker)
	tiPai := tiData.(*PHZPoker)
	paiCounts := CountHandPais(handPokers)
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			//log.T("测试用...paiValue=[%v]		paiCount=[%v]   checkPaiValue=[%v]", paiValue, paiCount, tiPai.GetValue())
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
			}
		}
	}
	if len(ti.Pais) > 0 {
		return ti, nil
	} else {
		return nil, nil
	}
}

func (p *ParserCore) CanPeng(userGameData interface{}, pengData interface{}) (interface{}, error) {
	handPokers := userGameData.([]*PHZPoker)
	pengPai := pengData.(*PHZPoker)
	pengReslut := &PengPai{}
	paiCounts := CountHandPais(handPokers)
	pengReslut.Pais = append(pengReslut.Pais, pengPai) //将桌面内的碰牌也存进去
	if paiCounts != nil {
		for paiValue, paiCount := range paiCounts {
			//log.T("测试用...paiValue=[%v]		paiCount=[%v]	checkPaiValue=[%v]", paiValue, paiCount, pengPai.GetValue())
			if paiValue == int(pengPai.GetValue()) && paiCount == 2 {
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
	paiCounts := CountHandPais(handPokers)
	for paiValue, paiCount := range paiCounts {
		//手上有一砍牌时，只能提或者跑，不能吃
		if int32(paiValue) == chiPoker.GetValue() && paiCount == 3 {
			return nil, nil
		}
		if paiCount == 3 {
			tempPokers := util.DeepClone(handPokers).([]*PHZPoker)
			for _, p := range tempPokers {
				if p.GetValue() == int32(paiValue) {
					handPokers = DelPaiFromPokers(handPokers, p)
				}
			}
		}
	}
	sameValueArray := GetPaisByValue2(handPokers, chiPoker.GetValue())
	//手牌里没有牌值与chiPoker一样的牌
	if sameValueArray == nil || len(sameValueArray) == 0 {
		log.T("找吃牌的组合时，手牌里没有牌值一样的牌")
		ret := GetYiJuHua1(handPokers, chiPoker)
		if ret != nil || len(ret) > 0 {
			chiResult = append(chiResult, ret...)
			return chiResult, nil
		}
	}
	//手牌里只有一张牌值与chiPoker一样的牌，需要比牌
	if len(sameValueArray) >= 1 {
		log.T("找可以吃的牌的组合时，手牌里有[%v]张牌值一样的牌，需要比牌", len(sameValueArray))
		tempPokers := util.DeepClone(handPokers).([]*PHZPoker)
		//将牌值相同的一张牌从手牌副本里删除，然后用chiPoker去找可以吃的组合
		tempResult := GetYiJuHua1(handPokers, chiPoker)
		if tempResult != nil || len(tempResult) > 0 {
			log.T("开始找每一种可以吃的牌型的比牌结果...")
			for _, tempChi := range tempResult {
				//找到每一种吃牌组合的比牌结果
				log.T("找[%v]的比牌结果...", Cards2String(tempChi.Pais))
				biResult, _ := GetBiPais(tempChi, tempPokers, chiPoker)
				if biResult != nil && len(biResult) > 0 {
					chiResult = append(chiResult, tempChi)
				}
			}
		}
		if chiResult != nil && len(chiResult) > 0 {
			return chiResult, nil
		}
	}
	return nil, nil
}

//校验每一组吃牌是否可以吃:从手牌里删除被吃的亮张牌，从剩余牌里找是否能够成一句话、一缴牌
func GetBiPais(chi *ChiPai, Pokers []*PHZPoker, checkPai *PHZPoker) ([]*ChiPai, error) {
	//删除手牌里的坎牌
	handPokers := util.DeepClone(Pokers).([]*PHZPoker)
	paiCounts := CountHandPais(handPokers)
	for paiValue, paiCount := range paiCounts {
		if paiCount == 3 {
			tempPokers := util.DeepClone(handPokers).([]*PHZPoker)
			for _, pai := range tempPokers {
				if pai.GetValue() == int32(paiValue) {
					handPokers = DelPaiFromPokers(handPokers, pai)
				}
			}
		}
	}
	var result []*ChiPai
	//删除被吃的牌
	handPokers = DelPaiFromPokers(handPokers, chi.Pais[1])
	handPokers = DelPaiFromPokers(handPokers, chi.Pais[2])
	//从删除后的手牌里找是否还有跟checkPai牌值相同的牌
	//如果没有了牌值相同的牌，则不用比牌，返回比牌数据为空
	if pais := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pais) == 0 {
		log.T("获取比牌时，手上没有牌值相同的牌，不用比牌")
		return nil, nil
	}
	//如果还有1张, 将这张牌删除然后从删除后的牌里找checkPai的吃牌，如果有则比牌成功，否则比牌失败
	if pais := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pais) == 1 {
		log.T("获取比牌时，手上有1张牌值相同的牌，需要比牌")
		handPokers = DelPaiFromPokers(handPokers, checkPai)
		tempResult := GetYiJuHua1(handPokers, pais[0])
		//比牌是一句话
		if tempResult != nil && len(tempResult) > 0 {
			result = append(result, tempResult...)
		}
		log.T("result=[%v]", result)
		if result != nil && len(result) > 0 {
			return result, nil
		} else {
			return nil, nil
		}
	}
	//如果删除后还有2张
	if pais := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pais) == 2 {
		log.T("获取比牌时，手上有2张牌值相同的牌，需要比牌")
		handPokers = DelPaiFromPokers(handPokers, checkPai)
		//比牌是一缴牌
		tempResult := GetYiJuHua1(handPokers, pais[0])
		if tempResult != nil && len(tempResult) > 0 {
			for _, c := range tempResult {
				if len(GetPaisByValue2(c.Pais, checkPai.GetValue())) == 2 {
					result = append(result, c)
				}
			}
		}
		//比牌是一句话
		handPokers = DelPaiFromPokers(handPokers, checkPai)
		tempResult = GetYiJuHua1(handPokers, pais[1])
		if tempResult != nil && len(tempResult) > 1 {
			result = append(result, tempResult...)
		}
		if result != nil && len(result) > 0 {
			log.T("result=[%v]", result)
			return result, nil
		} else {
			return nil, nil
		}
	}
	return nil, Error.NewError(-1, "没有比牌结果")
}

func DelPaiFromPokers(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
	for i, p := range handPokers {
		if p.GetValue() == poker.GetValue() {
			handPokers = append(handPokers[:i], handPokers[i+1:]...)
			break
		}
	}
	return handPokers
}

//从没有相同牌值的手牌里获取一句话、一缴牌
func GetYiJuHua1(handPokers []*PHZPoker, checkPai *PHZPoker) []*ChiPai {
	var result []*ChiPai
	var pai1, pai2 *PHZPoker
	//找贰柒拾
	if checkPai.GetValue() == 2 {
		pai1 = GetPaiByValue(handPokers, 7)
		pai2 = GetPaiByValue(handPokers, 10)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			data.HuXi = 3
			result = append(result, data)
		}
	}
	if checkPai.GetValue() == 12 {
		pai1 = GetPaiByValue(handPokers, 17)
		pai2 = GetPaiByValue(handPokers, 20)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			data.HuXi = 6
			result = append(result, data)
		}
	}
	if checkPai.GetValue() == 7 {
		pai1 = GetPaiByValue(handPokers, 2)
		pai2 = GetPaiByValue(handPokers, 10)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			data.HuXi = 3
			result = append(result, data)
		}
	}
	if checkPai.GetValue() == 17 {
		pai1 = GetPaiByValue(handPokers, 12)
		pai2 = GetPaiByValue(handPokers, 20)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			data.HuXi = 6
			result = append(result, data)
		}
	}
	if checkPai.GetValue() == 10 {
		pai1 = GetPaiByValue(handPokers, 2)
		pai2 = GetPaiByValue(handPokers, 7)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			data.HuXi = 3
			result = append(result, data)
		}
	}
	if checkPai.GetValue() == 20 {
		pai1 = GetPaiByValue(handPokers, 12)
		pai2 = GetPaiByValue(handPokers, 17)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			data.HuXi = 6
			result = append(result, data)
		}
	}

	//找一句话
	pai1 = GetPaiByValue(handPokers, checkPai.GetValue()-2)
	pai2 = GetPaiByValue(handPokers, checkPai.GetValue()-1)
	if pai1 != nil && pai2 != nil && IsSameValue(checkPai, pai2, pai1) {
		data := &ChiPai{}
		data.Pai = checkPai
		data.Pais = append(data.Pais, checkPai)
		data.Pais = append(data.Pais, pai2)
		data.Pais = append(data.Pais, pai1)
		if checkPai.GetValue() == 3 {
			data.HuXi = 3
		}
		if checkPai.GetValue() == 13 {
			data.HuXi = 6
		}
		result = append(result, data)
	}

	pai1 = GetPaiByValue(handPokers, checkPai.GetValue()-1)
	pai2 = GetPaiByValue(handPokers, checkPai.GetValue()+1)
	if pai1 != nil && pai2 != nil && IsSameValue(checkPai, pai1, pai2) {
		data := &ChiPai{}
		data.Pai = checkPai
		data.Pais = append(data.Pais, checkPai)
		data.Pais = append(data.Pais, pai2)
		data.Pais = append(data.Pais, pai1)
		if checkPai.GetValue() == 2 {
			data.HuXi = 3
		}
		if checkPai.GetValue() == 12 {
			data.HuXi = 6
		}
		result = append(result, data)
	}

	pai1 = GetPaiByValue(handPokers, checkPai.GetValue()+1)
	pai2 = GetPaiByValue(handPokers, checkPai.GetValue()+2)
	if pai1 != nil && pai2 != nil && IsSameValue(checkPai, pai1, pai2) {
		data := &ChiPai{}
		data.Pai = checkPai
		data.Pais = append(data.Pais, checkPai)
		data.Pais = append(data.Pais, pai2)
		data.Pais = append(data.Pais, pai1)
		if checkPai.GetValue() == 1 {
			data.HuXi = 3
		}
		if checkPai.GetValue() == 11 {
			data.HuXi = 6
		}
		result = append(result, data)
	}

	//找一缴牌
	//小 小 大
	if big := GetPaisByValue2(handPokers, checkPai.GetValue()-10); len(big) == 2 {
		log.T("len(pais)=[%v]", len(big))
		pai1 = GetPaisByValue2(handPokers, checkPai.GetValue()-10)[0]
		pai2 = GetPaisByValue2(handPokers, checkPai.GetValue()-10)[1]
		if pai2 != nil && pai1 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai1)
			data.Pais = append(data.Pais, pai2)
			result = append(result, data)
		}
	}

	//大 大 小
	if big := GetPaisByValue2(handPokers, checkPai.GetValue()+10); len(big) == 2 {
		log.T("len(pais)= [%v]", len(big))
		pai1 = GetPaisByValue2(handPokers, checkPai.GetValue()+10)[0]
		pai2 = GetPaisByValue2(handPokers, checkPai.GetValue()+10)[1]
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai1)
			data.Pais = append(data.Pais, pai2)
			result = append(result, data)
		}
	}

	//手里有1张牌值一样的牌
	if pai := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pai) == 1 {
		log.T("len(pai)=[%v]", len(pai))
		pai1 = GetPaiByValue(handPokers, checkPai.GetValue())
		pai2 = GetPaiByValue(handPokers, checkPai.GetValue()+10)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			result = append(result, data)
		}

		pai1 = GetPaiByValue(handPokers, checkPai.GetValue())
		pai2 = GetPaiByValue(handPokers, checkPai.GetValue()-10)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai1)
			data.Pais = append(data.Pais, pai2)
			result = append(result, data)
		}
	}
	//手里有2张牌值一样的牌
	if pai := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pai) == 2 {
		log.T("len(pai)=[%v]", len(pai))
		pai1 = GetPaisByValue2(handPokers, checkPai.GetValue())[0]
		pai2 = GetPaiByValue(handPokers, checkPai.GetValue()+10)
		if pai1 != nil && pai2 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai2)
			data.Pais = append(data.Pais, pai1)
			result = append(result, data)
		}

		pai1 = GetPaisByValue2(handPokers, checkPai.GetValue())[0]
		pai2 = GetPaiByValue(handPokers, checkPai.GetValue()-10)
		if pai2 != nil && pai1 != nil {
			data := &ChiPai{}
			data.Pai = checkPai
			data.Pais = append(data.Pais, checkPai)
			data.Pais = append(data.Pais, pai1)
			data.Pais = append(data.Pais, pai2)
			result = append(result, data)
		}
	}
	return result
}

func IsSameValue(poker1 *PHZPoker, poker2 *PHZPoker, poker3 *PHZPoker) bool {
	if poker3.GetValue() <= 10 && poker2.GetValue() <= 10 && poker1.GetValue() <= 10 {
		return true
	}
	if poker3.GetValue() > 10 && poker2.GetValue() > 10 && poker1.GetValue() > 10 {
		return true
	}
	return false
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
func GetPaiByValue(handPokers []*PHZPoker, value int32) *PHZPoker {
	for _, poker := range handPokers {
		if poker.GetValue() == value {
			return poker
		}
	}
	return nil
}

//根据牌值找到一组牌里有几张相同牌值的牌
func GetPaisByValue2(handPokers []*PHZPoker, value int32) []*PHZPoker {
	retPais := []*PHZPoker{}
	for _, poker := range handPokers {
		if poker.GetValue() == value {
			retPais = append(retPais, poker)
		}
	}
	return retPais
}

//校验玩家是否天胡：砌完牌后，玩家手中有3提或5坎，称为天胡
func (p *ParserCore) CheckHu(userGameData interface{}) (bool, error) {
	tiCount := make(map[int]int32)
	kanCount := make(map[int]int32)
	handPokers := userGameData.([]*PHZPoker)
	paiCounts := CountHandPais(handPokers)
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
