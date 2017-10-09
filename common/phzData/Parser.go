package phzData

import (
	"casino_common/common/Error"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"github.com/name5566/leaf/util"
	"reflect"
)

type PHZParser interface {
	CanTi(...interface{}) (interface{}, error)   //根据桌面的牌判断手牌里三张的牌是否可以跑
	GetTiPai(interface{}) interface{}            //获取手牌里原始可以提的牌
	CanPeng(...interface{}) (interface{}, error) //根据桌面的牌判断手牌里是否有可以碰的牌
	CanChi(...interface{}) (interface{}, error)  //根据桌面的牌判断手牌里是否有可以吃的牌型
	//TryHu(...interface{}) (interface{}, error)   //根据桌面的牌判断是否可以胡牌
	CheckHu(interface{}) (bool, error) //校验能否天胡、地胡
}

type ParserCore struct {
}

type HuInfo struct {
	YiJuHua     []*YJH      //胡牌时手牌里的一句话
	YiJiaoPai   []*YJH      //胡牌时手里的一缴牌
	DuiZis      []*DuiZi    //胡牌时手里的对子
	KanPais     []*YiKanPai //胡牌时手里的坎牌
	TiPais      []*YiTi     //胡牌时手里的提牌
	HuXi        int32       //胡息数
	WinUser     uint32      //赢家
	LoseUser    []uint32    //输家
	IsZimo      bool        //是否自摸
	IsDianPao   bool        //是否点炮
	DianPaoUser uint32      //点炮的玩家
	Fan         int32       //胡牌时的番数
	WinScore    int64       //得分
	HuPai       *PHZPoker   //胡的哪张牌
	HuType      int32       //胡牌类型
}

//一句话或者一缴牌
type YJH struct {
	Pais []*PHZPoker
	HuXi int32
}

//对子、将牌
type DuiZi struct {
	Pais []*PHZPoker
}

//一砍牌
type YiKanPai struct {
	Pais []*PHZPoker
	HuXi int32
}

//一提牌
type YiTi struct {
	Pais []*PHZPoker
	HuXi int32
}

func CountHandPais(pais []*PHZPoker) []int {
	counts := make([]int, 21) //牌值1——10:小字  11——20:大字
	if pais == nil || len(pais) == 0 {
		return nil
	}
	for _, p := range pais {
		if p == nil {
			continue
		}
		if p.GetPaiIndexByValue() == 0 {
			log.W("CountHandPais occur error...p.value==0")
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
				ti.Pais = append(ti.Pais, tiPai)
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
	handPokers := util.DeepClone(userGameData.([]*PHZPoker)).([]*PHZPoker)
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
				biResult, err := GetBiPais(tempChi, tempPokers, chiPoker)
				if biResult != nil && len(biResult) > 0 {
					chiResult = append(chiResult, tempChi)
				}
				if biResult == nil && err == nil {
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
	handPokers = DelPaiFromPokersByID(handPokers, chi.Pais[1])
	handPokers = DelPaiFromPokersByID(handPokers, chi.Pais[2])
	//从删除后的手牌里找是否还有跟checkPai牌值相同的牌
	//如果没有了牌值相同的牌，则不用比牌，返回比牌数据为空
	if pais := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pais) == 0 {
		log.T("获取比牌时，手上没有牌值相同的牌，不用比牌")
		return nil, nil
	}
	//如果还有1张, 将这张牌删除然后从删除后的牌里找checkPai的吃牌，如果有则比牌成功，否则比牌失败
	if pais := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pais) == 1 {
		log.T("获取比牌时，手上有1张牌值相同的牌，需要比牌")
		handPokers = DelPaiFromPokersByID(handPokers, pais[0])
		tempResult := GetYiJuHua1(handPokers, pais[0])
		//比牌是一句话
		if tempResult != nil && len(tempResult) > 0 {
			for _, t := range tempResult {
				log.T("删除吃牌后，从手牌找到的比牌结果有：[%v]", Cards2String(t.Pais))
			}
			result = append(result, tempResult...)
		}
		log.T("result=[%v]", result)
		if result != nil && len(result) > 0 {
			return result, nil
		} else {
			return nil, Error.NewError(-1, "需要比牌，没有找到比牌结果")
		}
	}
	//如果删除后还有2张
	if pais := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pais) == 2 {
		log.T("获取比牌时，手上有2张牌值相同的牌，需要比牌")
		handPokers = DelPaiFromPokersByID(handPokers, pais[0])
		//比牌是一缴牌
		tempResult := GetYiJuHua1(handPokers, pais[0])
		if tempResult != nil && len(tempResult) > 0 {
			for _, c := range tempResult {
				if len(GetPaisByValue2(c.Pais, pais[0].GetValue())) == 2 {
					log.T("删除手牌后，从手牌里找到的一缴牌比牌结果有：[%v]", Cards2String(c.Pais))
					result = append(result, c)
				}
			}
		}
		//比牌是一句话
		handPokers = DelPaiFromPokersByID(handPokers, pais[1])
		tempResult = GetYiJuHua1(handPokers, pais[1])
		if tempResult != nil && len(tempResult) > 1 {
			for _, c := range tempResult {
				log.T("删除手牌后，从手牌里找到的一句话比牌结果有：[%v]", Cards2String(c.Pais))
			}
			result = append(result, tempResult...)
		}
		if result != nil && len(result) > 0 {
			log.T("result=[%v]", result)
			return result, nil
		} else {
			return nil, Error.NewError(-1, "需要比牌，没有找到比牌结果")
		}
	}
	return nil, Error.NewError(-1, "没有比牌结果")
}

func DelPaiFromPokers(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
	for i, p := range handPokers {
		if p == nil {
			continue
		}
		if p.GetValue() == poker.GetValue() {
			handPokers = append(handPokers[:i], handPokers[i+1:]...)
			break
		}
	}
	return handPokers
}

func DelPaiFromPokersByID(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
	for i, p := range handPokers {
		if p.GetId() == poker.GetId() {
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
		if poker == nil {
			continue
		}
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
		if poker == nil {
			continue
		}
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

func TryHu(userGameData interface{}, pai interface{}, isOutCard bool) (interface{}, error) {
	pais := util.DeepClone(userGameData.([]*PHZPoker)).([]*PHZPoker)
	var checkPokers []*PHZPoker
	huInfo := &HuInfo{}
	var srcKan []*YiKanPai

	//找出牌组里原始坎牌
	if len(pais) >= 3 {
		count := CountHandPais(pais)
		for paiValue, paiCount := range count {
			if paiCount == 3 {
				kan := &YiKanPai{}
				pokers := GetPaisByValue2(pais, int32(paiValue))
				kan.Pais = append(kan.Pais, pokers...)
				if paiValue <= 10 {
					kan.HuXi = 3
				} else {
					kan.HuXi = 6
				}
				for _, p := range pokers {
					pais = DelPaiFromPokersByID(pais, p)
				}
				log.T("check胡牌时玩家手牌原始的坎牌是：[%v]", Cards2String(kan.Pais))
				srcKan = append(srcKan, kan)
			}
		}
	}
	var checkPai *PHZPoker
	v := reflect.ValueOf(pai)
	if !v.IsValid() || v.IsNil() {
		checkPokers = util.DeepClone(pais).([]*PHZPoker)
	} else {
		checkPai = pai.(*PHZPoker)
		checkPokers = append(util.DeepClone(pais).([]*PHZPoker), checkPai)
	}
	//如果找完原始的坎牌，手里没有牌了，则牌型上是可以胡了
	if checkPokers == nil || len(checkPokers) == 0 {
		huInfo.KanPais = append(huInfo.KanPais, srcKan...)
		log.T("check玩家是否可以胡时，玩家手里只有坎牌，可以胡牌...")
		return huInfo, nil
	}
	log.T("TryHu的牌是：[%v]", Cards2String(checkPokers))
	huInfo = GetHuInfoCase1(checkPokers, checkPai, isOutCard)
	if huInfo != nil {
		if srcKan != nil || len(srcKan) > 0 {
			huInfo.KanPais = append(huInfo.KanPais, srcKan...)
		}
		return huInfo, nil
	}
	huInfo = GetHuInfoCase2(checkPokers, checkPai, isOutCard)
	if huInfo != nil {
		if srcKan != nil || len(srcKan) > 0 {
			huInfo.KanPais = append(huInfo.KanPais, srcKan...)
		}
		return huInfo, nil
	}
	huInfo = GetHuInfoCase3(checkPokers, checkPai, isOutCard)
	if huInfo != nil {
		if srcKan != nil || len(srcKan) > 0 {
			huInfo.KanPais = append(huInfo.KanPais, srcKan...)
		}
		return huInfo, nil
	}
	huInfo = GetHuInfoCase4(checkPokers, checkPai, isOutCard)
	if huInfo != nil {
		if srcKan != nil || len(srcKan) > 0 {
			huInfo.KanPais = append(huInfo.KanPais, srcKan...)
		}
		return huInfo, nil
	}
	huInfo = GetHuInfoCase5(checkPokers, checkPai, isOutCard)
	if huInfo != nil {
		if srcKan != nil || len(srcKan) > 0 {
			huInfo.KanPais = append(huInfo.KanPais, srcKan...)
		}
		return huInfo, nil
	}
	huInfo = GetHuInfoCase6(checkPokers, checkPai, isOutCard)
	if huInfo != nil {
		if srcKan != nil || len(srcKan) > 0 {
			huInfo.KanPais = append(huInfo.KanPais, srcKan...)
		}
		return huInfo, nil
	}
	return nil, nil
}

func GetHuInfoCase1(pais []*PHZPoker, checkPai *PHZPoker, isOut bool) *HuInfo {
	checkPokers := util.DeepClone(pais).([]*PHZPoker)
	huInfo := &HuInfo{}

	for {
		//3、找一句话， 一句话函数内部已经按照胡息优先级找牌
		yijuhua, handPokers := GetYiJuHuaForHu(checkPokers)
		if yijuhua != nil && len(yijuhua) > 0 {
			/*for _, yjh := range yijuhua {
				log.T("check胡牌时找到的一句话是：[%v]", Cards2String(yjh.Pais))
			}*/
			huInfo.YiJuHua = append(huInfo.YiJuHua, yijuhua...)
		}
		checkPokers = handPokers

		//4、找一缴牌
		yijiaopai, handPokers := GetYiJiaoPaiForHu(checkPokers)
		if yijiaopai != nil && len(yijiaopai) > 0 {
			/*for _, yjp := range yijiaopai {
				log.T("check胡牌时找到的一缴牌是：[%v]", Cards2String(yjp.Pais))
			}*/
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yijiaopai...)
		}
		checkPokers = handPokers

		//2、找成坎的牌
		kanData, handPokers := GetKanPaisForHu(checkPokers)
		if kanData != nil && len(kanData) > 0 {
			for _, kan := range kanData {
				//log.T("check胡牌时找到的坎牌是：[%v]", Cards2String(kan.Pais))
				if checkPai != nil {
					if kan.Pais[0].GetValue() == checkPai.GetValue() && isOut {
						if kan.Pais[0].GetValue() <= 10 {
							kan.HuXi = 1
						} else {
							kan.HuXi = 3
						}
					}
				}
			}
			huInfo.KanPais = append(huInfo.KanPais, kanData...)
		}
		checkPokers = handPokers

		//5、找对子将牌
		duiZiData, handPokers := GetDuiZiForHu(checkPokers)
		if duiZiData != nil && len(duiZiData) > 0 {
			/*for _, duizi := range duiZiData {
				log.T("check胡牌时找到的对子是：[%v]", Cards2String(duizi.Pais))
			}*/
			huInfo.DuiZis = append(huInfo.DuiZis, duiZiData...)
		}
		checkPokers = handPokers
		log.T("checkHu 后剩余牌是：[%v]   [%v]", Cards2String(checkPokers), checkPokers)
		//log.T("len(checkPokers)= [%v]", len(checkPokers))
		if checkPokers == nil || len(checkPokers) == 0 {
			if duiZiData == nil || len(duiZiData) == 0 {
				return huInfo
			}
			if duiZiData != nil && len(duiZiData) == 1 {
				return huInfo
			}
			return nil
		} else if (kanData == nil || len(kanData) == 0) && (yijuhua == nil || len(yijuhua) == 0) &&
			(yijiaopai == nil || len(yijiaopai) == 0) && (duiZiData == nil || len(duiZiData) > 1) && len(checkPokers) > 0 {
			return nil
		}
	}
}

func GetHuInfoCase2(pais []*PHZPoker, checkPai *PHZPoker, isOut bool) *HuInfo {
	checkPokers := util.DeepClone(pais).([]*PHZPoker)
	huInfo := &HuInfo{}

	for {
		//3、找一句话， 一句话函数内部已经按照胡息优先级找牌
		yijuhua, handPokers := GetYiJuHuaForHu(checkPokers)
		if yijuhua != nil && len(yijuhua) > 0 {
			/*for _, yjh := range yijuhua {
				log.T("check胡牌时找到的一句话是：[%v]", Cards2String(yjh.Pais))
			}*/
			huInfo.YiJuHua = append(huInfo.YiJuHua, yijuhua...)
		}
		checkPokers = handPokers

		//2、找成坎的牌
		kanData, handPokers := GetKanPaisForHu(checkPokers)
		if kanData != nil && len(kanData) > 0 {
			for _, kan := range kanData {
				//log.T("check胡牌时找到的坎牌是：[%v]", Cards2String(kan.Pais))
				if checkPai != nil {
					if kan.Pais[0].GetValue() == checkPai.GetValue() && isOut {
						if kan.Pais[0].GetValue() <= 10 {
							kan.HuXi = 1
						} else {
							kan.HuXi = 3
						}
					}
				}
			}
			huInfo.KanPais = append(huInfo.KanPais, kanData...)
		}
		checkPokers = handPokers

		//4、找一缴牌
		yijiaopai, handPokers := GetYiJiaoPaiForHu(checkPokers)
		if yijiaopai != nil && len(yijiaopai) > 0 {
			/*for _, yjp := range yijiaopai {
				log.T("check胡牌时找到的一缴牌是：[%v]", Cards2String(yjp.Pais))
			}*/
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yijiaopai...)
		}
		checkPokers = handPokers

		//5、找对子将牌
		duiZiData, handPokers := GetDuiZiForHu(checkPokers)
		if duiZiData != nil && len(duiZiData) > 0 {
			/*for _, duizi := range duiZiData {
				log.T("check胡牌时找到的对子是：[%v]", Cards2String(duizi.Pais))
			}*/
			huInfo.DuiZis = append(huInfo.DuiZis, duiZiData...)
		}
		checkPokers = handPokers
		log.T("checkHu 后剩余牌是：[%v]   [%v]", Cards2String(checkPokers), checkPokers)
		if checkPokers == nil || len(checkPokers) == 0 {
			if duiZiData == nil || len(duiZiData) == 0 {
				return huInfo
			}
			if duiZiData != nil && len(duiZiData) == 1 {
				return huInfo
			}
			return nil
		} else if (kanData == nil || len(kanData) == 0) && (yijuhua == nil || len(yijuhua) == 0) &&
			(yijiaopai == nil || len(yijiaopai) == 0) && (duiZiData == nil || len(duiZiData) > 1) && len(checkPokers) > 0 {
			return nil
		}
	}
}

func GetHuInfoCase3(pais []*PHZPoker, checkPai *PHZPoker, isOut bool) *HuInfo {
	checkPokers := util.DeepClone(pais).([]*PHZPoker)
	huInfo := &HuInfo{}

	for {
		//4、找一缴牌
		yijiaopai, handPokers := GetYiJiaoPaiForHu(checkPokers)
		if yijiaopai != nil && len(yijiaopai) > 0 {
			/*for _, yjp := range yijiaopai {
				log.T("check胡牌时找到的一缴牌是：[%v]", Cards2String(yjp.Pais))
			}*/
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yijiaopai...)
		}
		checkPokers = handPokers
		//3、找一句话， 一句话函数内部已经按照胡息优先级找牌
		yijuhua, handPokers := GetYiJuHuaForHu(checkPokers)
		if yijuhua != nil && len(yijuhua) > 0 {
			/*for _, yjh := range yijuhua {
				log.T("check胡牌时找到的一句话是：[%v]", Cards2String(yjh.Pais))
			}*/
			huInfo.YiJuHua = append(huInfo.YiJuHua, yijuhua...)
		}
		checkPokers = handPokers
		//2、找成坎的牌
		kanData, handPokers := GetKanPaisForHu(checkPokers)
		if kanData != nil && len(kanData) > 0 {
			for _, kan := range kanData {
				//log.T("check胡牌时找到的坎牌是：[%v]", Cards2String(kan.Pais))
				if checkPai != nil {
					if kan.Pais[0].GetValue() == checkPai.GetValue() && isOut {
						if kan.Pais[0].GetValue() <= 10 {
							kan.HuXi = 1
						} else {
							kan.HuXi = 3
						}
					}
				}
			}
			huInfo.KanPais = append(huInfo.KanPais, kanData...)
		}
		checkPokers = handPokers

		//5、找对子将牌
		duiZiData, handPokers := GetDuiZiForHu(checkPokers)
		if duiZiData != nil && len(duiZiData) > 0 {
			/*for _, duizi := range duiZiData {
				log.T("check胡牌时找到的对子是：[%v]", Cards2String(duizi.Pais))
			}*/
			huInfo.DuiZis = append(huInfo.DuiZis, duiZiData...)
		}
		checkPokers = handPokers
		log.T("checkHu 后剩余牌是：[%v]   [%v]", Cards2String(checkPokers), checkPokers)
		if checkPokers == nil || len(checkPokers) == 0 {
			if duiZiData == nil || len(duiZiData) == 0 {
				return huInfo
			}
			if duiZiData != nil && len(duiZiData) == 1 {
				return huInfo
			}
			return nil
		} else if (kanData == nil || len(kanData) == 0) && (yijuhua == nil || len(yijuhua) == 0) &&
			(yijiaopai == nil || len(yijiaopai) == 0) && (duiZiData == nil || len(duiZiData) > 1) && len(checkPokers) > 0 {
			return nil
		}
	}
}

func GetHuInfoCase4(pais []*PHZPoker, checkPai *PHZPoker, isOut bool) *HuInfo {
	checkPokers := util.DeepClone(pais).([]*PHZPoker)
	huInfo := &HuInfo{}

	for {
		//4、找一缴牌
		yijiaopai, handPokers := GetYiJiaoPaiForHu(checkPokers)
		if yijiaopai != nil && len(yijiaopai) > 0 {
			/*for _, yjp := range yijiaopai {
				log.T("check胡牌时找到的一缴牌是：[%v]", Cards2String(yjp.Pais))
			}*/
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yijiaopai...)
		}
		checkPokers = handPokers

		//2、找成坎的牌
		kanData, handPokers := GetKanPaisForHu(checkPokers)
		if kanData != nil && len(kanData) > 0 {
			for _, kan := range kanData {
				//log.T("check胡牌时找到的坎牌是：[%v]", Cards2String(kan.Pais))
				if checkPai != nil {
					if kan.Pais[0].GetValue() == checkPai.GetValue() && isOut {
						if kan.Pais[0].GetValue() <= 10 {
							kan.HuXi = 1
						} else {
							kan.HuXi = 3
						}
					}
				}
			}
			huInfo.KanPais = append(huInfo.KanPais, kanData...)
		}
		checkPokers = handPokers

		//3、找一句话， 一句话函数内部已经按照胡息优先级找牌
		yijuhua, handPokers := GetYiJuHuaForHu(checkPokers)
		if yijuhua != nil && len(yijuhua) > 0 {
			/*for _, yjh := range yijuhua {
				log.T("check胡牌时找到的一句话是：[%v]", Cards2String(yjh.Pais))
			}*/
			huInfo.YiJuHua = append(huInfo.YiJuHua, yijuhua...)
		}
		checkPokers = handPokers

		//5、找对子将牌
		duiZiData, handPokers := GetDuiZiForHu(checkPokers)
		if duiZiData != nil && len(duiZiData) > 0 {
			/*for _, duizi := range duiZiData {
				log.T("check胡牌时找到的对子是：[%v]", Cards2String(duizi.Pais))
			}*/
			huInfo.DuiZis = append(huInfo.DuiZis, duiZiData...)
		}
		checkPokers = handPokers
		log.T("checkHu 后剩余牌是：[%v]   [%v]", Cards2String(checkPokers), checkPokers)
		if checkPokers == nil || len(checkPokers) == 0 {
			if duiZiData == nil || len(duiZiData) == 0 {
				return huInfo
			}
			if duiZiData != nil && len(duiZiData) == 1 {
				return huInfo
			}
			return nil
		} else if (kanData == nil || len(kanData) == 0) && (yijuhua == nil || len(yijuhua) == 0) &&
			(yijiaopai == nil || len(yijiaopai) == 0) && (duiZiData == nil || len(duiZiData) > 1) && len(checkPokers) > 0 {
			return nil
		}
	}
}

func GetHuInfoCase5(pais []*PHZPoker, checkPai *PHZPoker, isOut bool) *HuInfo {
	checkPokers := util.DeepClone(pais).([]*PHZPoker)
	huInfo := &HuInfo{}

	for {
		//2、找成坎的牌
		kanData, handPokers := GetKanPaisForHu(checkPokers)
		if kanData != nil && len(kanData) > 0 {
			for _, kan := range kanData {
				//log.T("check胡牌时找到的坎牌是：[%v]", Cards2String(kan.Pais))
				if checkPai != nil {
					if kan.Pais[0].GetValue() == checkPai.GetValue() && isOut {
						if kan.Pais[0].GetValue() <= 10 {
							kan.HuXi = 1
						} else {
							kan.HuXi = 3
						}
					}
				}
			}
			huInfo.KanPais = append(huInfo.KanPais, kanData...)
		}
		checkPokers = handPokers

		//4、找一缴牌
		yijiaopai, handPokers := GetYiJiaoPaiForHu(checkPokers)
		if yijiaopai != nil && len(yijiaopai) > 0 {
			/*for _, yjp := range yijiaopai {
				log.T("check胡牌时找到的一缴牌是：[%v]", Cards2String(yjp.Pais))
			}*/
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yijiaopai...)
		}
		checkPokers = handPokers

		//3、找一句话， 一句话函数内部已经按照胡息优先级找牌
		yijuhua, handPokers := GetYiJuHuaForHu(checkPokers)
		if yijuhua != nil && len(yijuhua) > 0 {
			/*for _, yjh := range yijuhua {
				log.T("check胡牌时找到的一句话是：[%v]", Cards2String(yjh.Pais))
			}*/
			huInfo.YiJuHua = append(huInfo.YiJuHua, yijuhua...)
		}
		checkPokers = handPokers

		//5、找对子将牌
		duiZiData, handPokers := GetDuiZiForHu(checkPokers)
		if duiZiData != nil && len(duiZiData) > 0 {
			/*for _, duizi := range duiZiData {
				log.T("check胡牌时找到的对子是：[%v]", Cards2String(duizi.Pais))
			}*/
			huInfo.DuiZis = append(huInfo.DuiZis, duiZiData...)
		}
		checkPokers = handPokers
		log.T("checkHu 后剩余牌是：[%v]   [%v]", Cards2String(checkPokers), checkPokers)
		if checkPokers == nil || len(checkPokers) == 0 {
			if duiZiData == nil || len(duiZiData) == 0 {
				return huInfo
			}
			if duiZiData != nil && len(duiZiData) == 1 {
				return huInfo
			}
			return nil
		} else if (kanData == nil || len(kanData) == 0) && (yijuhua == nil || len(yijuhua) == 0) &&
			(yijiaopai == nil || len(yijiaopai) == 0) && (duiZiData == nil || len(duiZiData) > 1) && len(checkPokers) > 0 {
			return nil
		}
	}
}

func GetHuInfoCase6(pais []*PHZPoker, checkPai *PHZPoker, isOut bool) *HuInfo {
	checkPokers := util.DeepClone(pais).([]*PHZPoker)
	huInfo := &HuInfo{}

	for {
		//2、找成坎的牌
		kanData, handPokers := GetKanPaisForHu(checkPokers)
		if kanData != nil && len(kanData) > 0 {
			for _, kan := range kanData {
				//log.T("check胡牌时找到的坎牌是：[%v]", Cards2String(kan.Pais))
				if checkPai != nil {
					if kan.Pais[0].GetValue() == checkPai.GetValue() && isOut {
						if kan.Pais[0].GetValue() <= 10 {
							kan.HuXi = 1
						} else {
							kan.HuXi = 3
						}
					}
				}
			}
			huInfo.KanPais = append(huInfo.KanPais, kanData...)
		}
		checkPokers = handPokers

		//3、找一句话， 一句话函数内部已经按照胡息优先级找牌
		yijuhua, handPokers := GetYiJuHuaForHu(checkPokers)
		if yijuhua != nil && len(yijuhua) > 0 {
			/*for _, yjh := range yijuhua {
				log.T("check胡牌时找到的一句话是：[%v]", Cards2String(yjh.Pais))
			}*/
			huInfo.YiJuHua = append(huInfo.YiJuHua, yijuhua...)
		}
		checkPokers = handPokers

		//4、找一缴牌
		yijiaopai, handPokers := GetYiJiaoPaiForHu(checkPokers)
		if yijiaopai != nil && len(yijiaopai) > 0 {
			/*for _, yjp := range yijiaopai {
				log.T("check胡牌时找到的一缴牌是：[%v]", Cards2String(yjp.Pais))
			}*/
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yijiaopai...)
		}
		checkPokers = handPokers

		//5、找对子将牌
		duiZiData, handPokers := GetDuiZiForHu(checkPokers)
		if duiZiData != nil && len(duiZiData) > 0 {
			/*for _, duizi := range duiZiData {
				log.T("check胡牌时找到的对子是：[%v]", Cards2String(duizi.Pais))
			}*/
			huInfo.DuiZis = append(huInfo.DuiZis, duiZiData...)
		}
		checkPokers = handPokers
		log.T("checkHu 后剩余牌是：[%v]   [%v]", Cards2String(checkPokers), checkPokers)
		if checkPokers == nil || len(checkPokers) == 0 {
			if duiZiData == nil || len(duiZiData) == 0 {
				return huInfo
			}
			if duiZiData != nil && len(duiZiData) == 1 {
				return huInfo
			}
			return nil
		} else if (kanData == nil || len(kanData) == 0) && (yijuhua == nil || len(yijuhua) == 0) &&
			(yijiaopai == nil || len(yijiaopai) == 0) && (duiZiData == nil || len(duiZiData) > 1) && len(checkPokers) > 0 {
			return nil
		}
	}
}

//从手牌里获取提牌
func GetTiPaisForHu(pais []*PHZPoker) (result []*YiTi, remainPokers []*PHZPoker) {
	handPokers := util.DeepClone(pais).([]*PHZPoker)
	counts := CountHandPais(handPokers)
	if counts == nil {
		return nil, nil
	}
	for paiValue, paiCount := range counts {
		if paiCount == 4 {
			tiData := &YiTi{}
			tiData.Pais = GetPaisByValue2(handPokers, int32(paiValue))
			if paiValue <= 10 {
				tiData.HuXi = 9
			} else {
				tiData.HuXi = 12
			}
			result = append(result, tiData)
			for i := 0; i < 4; i++ {
				pai := GetPaiByValue(handPokers, int32(paiValue))
				if pai != nil {
					handPokers = DelPaiFromPokers(handPokers, pai)
				}
			}
		}
	}
	return result, handPokers
}

//从手牌里获取所有的砍牌
func GetKanPaisForHu(pais []*PHZPoker) (result []*YiKanPai, remainPokers []*PHZPoker) {
	handPokers := util.DeepClone(pais).([]*PHZPoker)
	counts := CountHandPais(handPokers)
	if counts == nil {
		return nil, nil
	}
	for paiValue, paiCount := range counts {
		if paiCount == 3 {
			//找到坎牌后保存，并将牌删除返回剩余牌
			kanData := &YiKanPai{}
			kanData.Pais = GetPaisByValue2(handPokers, int32(paiValue))
			if paiValue <= 10 {
				kanData.HuXi = 1
			} else if paiValue > 10 {
				kanData.HuXi = 3
			}
			result = append(result, kanData)
			//删除坎牌
			for i := 0; i < 3; i++ {
				pai := GetPaiByValue(handPokers, int32(paiValue))
				if pai != nil {
					handPokers = DelPaiFromPokers(handPokers, pai)
				}
			}
		}
	}
	return result, handPokers
}

func GetYiJuHuaForHu(handPais []*PHZPoker) ([]*YJH, []*PHZPoker) {
	handPokers := util.DeepClone(handPais).([]*PHZPoker)
	counts := CountHandPais(handPokers)
	if counts == nil {
		return nil, nil
	}
	var result []*YJH

	//找2 7 10
	//大字
	if counts[12] > 0 && counts[17] > 0 && counts[20] > 0 {
		yijuhua := &YJH{}
		pai1 := GetPaiByValue(handPokers, 12)
		pai2 := GetPaiByValue(handPokers, 17)
		pai3 := GetPaiByValue(handPokers, 20)
		if pai1 != nil && pai2 != nil && pai3 != nil {
			yijuhua.Pais = append(yijuhua.Pais, pai1)
			handPokers = DelPaiFromPokers(handPokers, pai1)

			yijuhua.Pais = append(yijuhua.Pais, pai2)
			handPokers = DelPaiFromPokers(handPokers, pai2)

			yijuhua.Pais = append(yijuhua.Pais, pai3)
			handPokers = DelPaiFromPokers(handPokers, pai3)
			yijuhua.HuXi = 6
			result = append(result, yijuhua)
		}
	}

	//找大字的一句话
	counts = CountHandPais(handPokers)
	if counts == nil {
		return result, nil
	}
	for paiValue := 11; paiValue <= 18; paiValue++ {
		if counts[paiValue] > 0 && counts[paiValue+1] > 0 && counts[paiValue+2] > 0 {
			yijuhua := &YJH{}
			for i := 0; i < 3; i++ {
				pai := GetPaiByValue(handPokers, int32(paiValue+i))
				if pai == nil {
					break
				}
				handPokers = DelPaiFromPokers(handPokers, pai)
				yijuhua.Pais = append(yijuhua.Pais, pai)
			}
			if paiValue == 11 {
				yijuhua.HuXi = 6
			}
			result = append(result, yijuhua)
		}
	}

	//小字 2 7 10
	counts = CountHandPais(handPokers)
	if counts == nil {
		return result, nil
	}
	if counts[2] > 0 && counts[7] > 0 && counts[10] > 0 {
		yijuhua := &YJH{}
		pai1 := GetPaiByValue(handPokers, 2)
		pai2 := GetPaiByValue(handPokers, 7)
		pai3 := GetPaiByValue(handPokers, 10)
		if pai1 != nil && pai2 != nil && pai3 != nil {
			yijuhua.Pais = append(yijuhua.Pais, pai1)
			handPokers = DelPaiFromPokers(handPokers, pai1)

			yijuhua.Pais = append(yijuhua.Pais, pai2)
			handPokers = DelPaiFromPokers(handPokers, pai2)

			yijuhua.Pais = append(yijuhua.Pais, pai3)
			handPokers = DelPaiFromPokers(handPokers, pai3)
			yijuhua.HuXi = 3
			result = append(result, yijuhua)
		}
	}

	//找小字的一句话
	counts = CountHandPais(handPokers)
	if counts == nil {
		return result, nil
	}
	for paiValue := 1; paiValue <= 8; paiValue++ {
		if counts[paiValue] > 0 && counts[paiValue+1] > 0 && counts[paiValue+2] > 0 {
			yijuhua := &YJH{}
			for i := 0; i < 3; i++ {
				pai := GetPaiByValue(handPokers, int32(paiValue+i))
				if pai == nil {
					break
				}
				handPokers = DelPaiFromPokers(handPokers, pai)
				yijuhua.Pais = append(yijuhua.Pais, pai)
			}
			if paiValue == 1 {
				yijuhua.HuXi = 3
			}
			result = append(result, yijuhua)
		}
	}
	return result, handPokers
}

//从手牌里找一缴牌
func GetYiJiaoPaiForHu(pais []*PHZPoker) (result []*YJH, remainPokers []*PHZPoker) {
	handPokers := util.DeepClone(pais).([]*PHZPoker)
	counts := CountHandPais(handPokers)
	if counts == nil {
		return nil, nil
	}

	for i := 1; i <= 10; i++ {
		if counts[i] == 1 {
			if counts[i+10] == 2 {
				yijiaopai := &YJH{}
				pai1 := GetPaiByValue(handPokers, int32(i))
				if pai1 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai1)
				handPokers = DelPaiFromPokers(handPokers, pai1)

				pai2 := GetPaiByValue(handPokers, int32(i+10))
				if pai2 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai2)
				handPokers = DelPaiFromPokers(handPokers, pai2)

				pai3 := GetPaiByValue(handPokers, int32(i+10))
				if pai3 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai3)
				handPokers = DelPaiFromPokers(handPokers, pai3)
				result = append(result, yijiaopai)
			}
		}
		counts = CountHandPais(handPokers)
		if counts == nil {
			return result, nil
		}
		if counts[i] == 2 {
			if counts[i+10] == 1 {
				yijiaopai := &YJH{}
				pai1 := GetPaiByValue(handPokers, int32(i))
				if pai1 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai1)
				handPokers = DelPaiFromPokers(handPokers, pai1)

				pai2 := GetPaiByValue(handPokers, int32(i))
				if pai2 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai2)
				handPokers = DelPaiFromPokers(handPokers, pai2)

				pai3 := GetPaiByValue(handPokers, int32(i+10))
				if pai3 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai3)
				handPokers = DelPaiFromPokers(handPokers, pai3)
				result = append(result, yijiaopai)
			}
		}

		counts = CountHandPais(handPokers)
		if counts == nil {
			return result, nil
		}
		if counts[i] == 2 {
			if counts[i+10] == 2 {
				yijiaopai := &YJH{}
				pai1 := GetPaiByValue(handPokers, int32(i))
				if pai1 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai1)
				handPokers = DelPaiFromPokers(handPokers, pai1)

				pai2 := GetPaiByValue(handPokers, int32(i))
				if pai2 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai2)
				handPokers = DelPaiFromPokers(handPokers, pai2)

				pai3 := GetPaiByValue(handPokers, int32(i+10))
				if pai3 == nil {
					continue
				}
				yijiaopai.Pais = append(yijiaopai.Pais, pai3)
				handPokers = DelPaiFromPokers(handPokers, pai3)
				result = append(result, yijiaopai)
			}
		}
	}
	return result, handPokers
}

//从手牌里找对子
func GetDuiZiForHu(pais []*PHZPoker) (result []*DuiZi, remainPokers []*PHZPoker) {
	handPokers := util.DeepClone(pais).([]*PHZPoker)
	counts := CountHandPais(handPokers)
	if counts == nil {
		return nil, nil
	}
	for paiValue, paiCount := range counts {
		if paiCount == 2 {
			duizi := &DuiZi{}
			pokers := GetPaisByValue2(handPokers, int32(paiValue))
			if pokers == nil || len(pokers) < 2 {
				continue
			}
			duizi.Pais = pokers
			handPokers = DelPaiFromPokers(handPokers, pokers[0])
			handPokers = DelPaiFromPokers(handPokers, pokers[1])
			result = append(result, duizi)
		}
	}
	return result, handPokers
}
