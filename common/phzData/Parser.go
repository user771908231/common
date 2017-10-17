package phzData

import (
	"casino_common/common/Error"
	"casino_common/common/log"
	"casino_common/proto/ddproto"
	"fmt"
	"github.com/name5566/leaf/util"
	"reflect"
)

const (
	TOTALPAIVALUE_COUNT = 20 //牌张
	PAIVALUE_SMALL      = 10 //小字
	CANHU_LIMIT_HUXI    = 15 //满足胡牌条件的胡息数
)

type PHZParser interface {
	CanTi(...interface{}) (interface{}, error)   //根据桌面的牌判断手牌里三张的牌是否可以跑
	GetTiPai(interface{}) interface{}            //获取手牌里原始可以提的牌
	CanPeng(...interface{}) (interface{}, error) //根据桌面的牌判断手牌里是否有可以碰的牌
	CanChi(...interface{}) (interface{}, error)  //根据桌面的牌判断手牌里是否有可以吃的牌型
	CheckHu(interface{}) (bool, error)           //校验能否天胡、地胡
}

type ParserCore struct {
}

type HuInfo struct {
	CanHu       bool        //能否胡牌
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
	HuType      []int32     //胡牌类型
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
	counts := make([]int, TOTALPAIVALUE_COUNT+1) //牌值1——10:小字  11——20:大字
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
	log.T("判断是否可以吃牌时，删除原始坎牌前的牌是：[%v]", Cards2String(handPokers))
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
	log.T("判断是否可以吃牌时，删除原始坎牌后的牌是：[%v]", Cards2String(handPokers))
	sameValueArray := GetPaisByValue2(handPokers, chiPoker.GetValue())
	//手牌里没有牌值与chiPoker一样的牌
	if sameValueArray == nil || len(sameValueArray) == 0 {
		log.T("找吃牌的组合时，手牌里没有牌值一样的牌")
		ret := GetYiJuHua1(handPokers, chiPoker)
		if ret != nil || len(ret) > 0 {
			chiResult = append(chiResult, ret...)
			for _, chipais := range chiResult {
				log.T("判断是否可以吃牌时，手牌里没有牌值一样的牌，可以吃的牌有：[%v]", Cards2String(chipais.Pais))
			}
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
				if err != nil {
					log.T("找[%v]的比牌结果时，需要比牌，但没有找到比牌结果...不能吃", Cards2String(tempChi.Pais))
				}
				if biResult != nil && len(biResult) > 0 {
					log.T("[%v]的比牌结果有：", Cards2String(tempChi.Pais))
					for _, bb := range biResult {
						log.T("[%v]", Cards2String(bb.Pais))
					}
					chiResult = append(chiResult, tempChi)
				}
				/*if biResult == nil && err == nil {
					log.T("[%v]不用比牌，可以吃", Cards2String(tempChi.Pais))
					chiResult = append(chiResult, tempChi)
				}*/
			}
		}
		if chiResult != nil && len(chiResult) > 0 {
			for _, cc := range chiResult {
				log.T("[%v]的吃牌结果有：[%v]", Card2String(chiPoker), Cards2String(cc.Pais))
			}
			return chiResult, nil
		}
	}
	log.T("不能吃牌[%v]", Card2String(chiPoker))
	return nil, nil
}

//校验每一组吃牌是否可以吃:从手牌里删除被吃的亮张牌，从剩余牌里找是否能够成一句话、一缴牌
func GetBiPais(chi *ChiPai, Pokers []*PHZPoker, checkPai *PHZPoker) (result []*ChiPai, err error) {
	//删除手牌里的坎牌
	handPokers := util.DeepClone(Pokers).([]*PHZPoker)
	paiCounts := CountHandPais(handPokers)
	//log.T("找[%v]的比牌时，删除原始坎牌前的牌是：[%v]", Cards2String(chi.Pais), Cards2String(handPokers))
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
	//log.T("找[%v]的比牌时，删除原始坎牌后的牌是：[%v]", Cards2String(chi.Pais), Cards2String(handPokers))
	//var result []*ChiPai
	//删除被吃的牌
	handPokers = DelPaiFromPokersByID(handPokers, chi.Pais[1])
	handPokers = DelPaiFromPokersByID(handPokers, chi.Pais[2])
	//从删除后的手牌里找是否还有跟checkPai牌值相同的牌
	//如果没有了牌值相同的牌，则不用比牌，返回比牌数据为空
	if pais := GetPaisByValue2(handPokers, checkPai.GetValue()); len(pais) == 0 {
		log.T("获取[%v]的比牌时，手上没有牌值相同的牌，不用比牌", Cards2String(chi.Pais))
		result = append(result, chi)
		return result, nil
	} else {
		log.T("找[%v]的比牌结果时，手上还有[%v]张牌值相同的牌", Cards2String(chi.Pais), len(pais))
		handPokers = DelPaiFromPokersByID(handPokers, pais[0])
		tempResult := GetYiJuHua1(handPokers, pais[0])
		if tempResult != nil && len(tempResult) > 0 {
			for _, chiData := range tempResult {
				tempBiResult, err := GetBiPais(chiData, handPokers, pais[0])
				if err != nil {
					log.T("找[%v]的比牌时err! = nil", Cards2String(chiData.Pais))
					continue
				}
				if tempBiResult != nil && len(tempBiResult) > 0 {
					for _, bb := range tempBiResult {
						log.T("找[%v]的比牌时，比牌结果是：[%v]", Cards2String(chiData.Pais), Cards2String(bb.Pais))
					}
					result = append(result, tempBiResult...)
				}
				if tempBiResult != nil && len(tempBiResult) == 0 {
					log.T("找[%v]的比牌时，不用比牌")
					result = append(result, chiData)
				}
			}
			return result, nil
		}
		return nil, Error.NewError(-1, "没有比牌结果")
	}
}

func DelPaiFromPokers(handPokers []*PHZPoker, poker *PHZPoker) []*PHZPoker {
	if poker == nil {
		log.W("DelPaiFromPokers 时poker==nil")
		return handPokers
	}
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

type jiao []int

//坎的胡息在之前计算出来 传入
type CanHuInfo struct {
	canHu             bool
	jiang             int32  //将牌 如果有
	countBigErQiShi   int    //大字二七十的数量
	countSmallErQiShi int    //小字二七十的数量
	countBigYiErSan   int    //大字一二三的数量
	countSmallYiErSan int    //小字一二三的数量
	jiaos             []jiao //绞牌
	yijuhuas          []int  //一句话
	kans              []int  //一坎牌
	pengs             []int  //碰牌
	totalHuXi         int32  //总胡息
}

func (info CanHuInfo) OnInit() {
	info.canHu = false
	info.jiang = -1
	info.countBigErQiShi = 0
	info.countSmallErQiShi = 0
	info.countBigYiErSan = 0
	info.countSmallYiErSan = 0
	info.jiaos = nil
	info.yijuhuas = nil
	info.kans = nil
	info.totalHuXi = 0
}

func CanHu(huxi int32, count []int, len int, zimoPaiValue int32) (info CanHuInfo) {
	//	log.T("开始判断CanHu(huxi:[%+v], count:[%+v], len:[%+v], zimoPaiValue:[%+v])", huxi, count, len, zimoPaiValue)
	//fmt.Println(fmt.Sprintf("开始判断CanHu(huxi:[%+v], count:[%+v], len:[%+v], zimoPaiValue:[%+v])", huxi, count, len, zimoPaiValue))

	//初始化默认值
	info.OnInit()
	info.totalHuXi = huxi

	//递归完所有的牌
	if len == 0 {
		if info.totalHuXi >= CANHU_LIMIT_HUXI {
			//胡息满足条件表示胡了
			info.canHu = true
			fmt.Println(fmt.Sprintf("len == 0 && huxi >= [%v] return info:[%+v]", info, CANHU_LIMIT_HUXI))
			return info
		}
		//已经递归完所有牌 胡息不满足
		info.canHu = false
		//fmt.Println(fmt.Sprintf("len == 0 && huxi < [%v] return info:[%+v]", info, CANHU_LIMIT_HUXI))
		return info
	}

	if len%3 != 2 {
		/*这里注意胡息的优先级*/

		// 三个一样的 坎或碰
		for i := 1; i < TOTALPAIVALUE_COUNT+1; i++ {
			if count[i] >= 3 {
				count[i] -= 3

				kanPengHuxi := int32(0)
				if int(zimoPaiValue) == i {
					//自摸为坎
					if i > PAIVALUE_SMALL {
						//大字坎
						kanPengHuxi = 6
					} else {
						//小字坎
						kanPengHuxi = 3
					}
				} else {
					//其他为碰
					if i > PAIVALUE_SMALL {
						//大字碰
						kanPengHuxi = 3
					} else {
						//小字碰
						kanPengHuxi = 1
					}
				}

				info.totalHuXi += kanPengHuxi

				info = CanHu(info.totalHuXi, count, len-3, zimoPaiValue)
				if info.canHu {
					if int(zimoPaiValue) == i {
						//自摸为坎
						info.kans = append(info.kans, i)
					} else {
						//其他为碰
						info.pengs = append(info.pengs, i)
					}
					//fmt.Println(fmt.Sprintf("找到三个一样的 return info:[%+v]", info))
					return info
				}
				info.totalHuXi -= kanPengHuxi

				count[i] += 3
			}
		}

		//大字2 7 10
		if count[12] > 0 && count[17] > 0 && count[20] > 0 {
			count[12] -= 1
			count[17] -= 1
			count[20] -= 1

			info.totalHuXi += 6 //大字2 7 10

			info = CanHu(info.totalHuXi, count, len-3, zimoPaiValue)
			if info.canHu {
				//log.T("i: %v, value: %v", i, count[i])
				info.countBigErQiShi++
				//fmt.Println(fmt.Sprintf("找到大字二七十 return info:[%+v]", info))
				return info
			}

			info.totalHuXi -= 6 //大字2 7 10

			count[12] += 1
			count[17] += 1
			count[20] += 1
		}
		//大字1 2 3
		if count[11] > 0 && count[12] > 0 && count[13] > 0 {
			count[11] -= 1
			count[12] -= 1
			count[13] -= 1

			info.totalHuXi += 6 //大字1 2 3

			info = CanHu(info.totalHuXi, count, len-3, zimoPaiValue)
			if info.canHu {
				//log.T("i: %v, value: %v", i, count[i])
				info.countBigYiErSan++
				//fmt.Println(fmt.Sprintf("找到大字一二三 return info:[%+v]", info))
				return info
			}

			info.totalHuXi -= 6 //大字1 2 3

			count[11] += 1
			count[12] += 1
			count[13] += 1
		}

		//小字2 7 10
		if count[2] > 0 && count[7] > 0 && count[10] > 0 {
			count[2] -= 1
			count[7] -= 1
			count[10] -= 1

			info.totalHuXi += 3 //小字2 7 10

			info = CanHu(info.totalHuXi, count, len-3, zimoPaiValue)
			if info.canHu {
				//log.T("i: %v, value: %v", i, count[i])
				info.countSmallErQiShi++
				//fmt.Println(fmt.Sprintf("找到小字二七十 return info:[%+v]", info))
				return info
			}

			info.totalHuXi -= 3 //小字2 7 10

			count[2] += 1
			count[7] += 1
			count[10] += 1
		}
		//小字1 2 3
		if count[1] > 0 && count[2] > 0 && count[3] > 0 {
			count[1] -= 1
			count[2] -= 1
			count[3] -= 1

			info.totalHuXi += 3 //小字1 2 3

			info = CanHu(info.totalHuXi, count, len-3, zimoPaiValue)
			if info.canHu {
				//log.T("i: %v, value: %v", i, count[i])
				info.countSmallYiErSan++
				//fmt.Println(fmt.Sprintf("找到小字一二三 return info:[%+v]", info))
				return info
			}

			info.totalHuXi -= 3 //小字1 2 3

			count[1] += 1
			count[2] += 1
			count[3] += 1
		}

		//是否是一句话（顺），这里应该分开判断
		//小字 一句话
		for i := 1; i < PAIVALUE_SMALL-1; i++ {
			if count[i] > 0 && count[i+1] > 0 && count[i+2] > 0 {
				count[i] -= 1
				count[i+1] -= 1
				count[i+2] -= 1
				info = CanHu(huxi, count, len-3, zimoPaiValue)
				if info.canHu {
					//log.T("i: %v, value: %v", i, count[i])
					info.yijuhuas = append(info.yijuhuas, i)
					//fmt.Println(fmt.Sprintf("找到小字一句话 return info:[%+v]", info))
					return info
				}
				count[i] += 1
				count[i+1] += 1
				count[i+2] += 1
			}
		}
		//大字 一句话
		for i := PAIVALUE_SMALL + 1; i < TOTALPAIVALUE_COUNT-1; i++ {
			if count[i] > 0 && count[i+1] > 0 && count[i+2] > 0 {
				count[i] -= 1
				count[i+1] -= 1
				count[i+2] -= 1
				info = CanHu(huxi, count, len-3, zimoPaiValue)
				if info.canHu {
					//log.T("i: %v, value: %v", i, count[i])
					info.yijuhuas = append(info.yijuhuas, i)
					//fmt.Println(fmt.Sprintf("找到大字一句话 return info:[%+v]", info))
					return info
				}
				count[i] += 1
				count[i+1] += 1
				count[i+2] += 1
			}
		}

		//是否是一绞牌 根据小字去组合大字
		for i := 1; i < PAIVALUE_SMALL+1; i++ {
			//fmt.Println(fmt.Sprintf("是否是一绞牌[%v] 小字count[%v]:%v 大字count[%v]:%v", i, i, count[i], PAIVALUE_SMALL+i, count[PAIVALUE_SMALL+i]))
			if count[i] > 0 && count[PAIVALUE_SMALL+i] >= 2 {
				//fmt.Println("case 1")
				count[i] -= 1
				count[PAIVALUE_SMALL+i] -= 2
				info = CanHu(huxi, count, len-3, zimoPaiValue)
				if info.canHu {
					//log.T("i: %v, value: %v", i, count[i])
					jiao := jiao{i, PAIVALUE_SMALL + i, PAIVALUE_SMALL + i}
					info.jiaos = append(info.jiaos, jiao)
					//fmt.Println(fmt.Sprintf("找到一绞牌 case1 return info:[%+v]", info))
					return info
				}
				count[i] += 1
				count[PAIVALUE_SMALL+i] += 2

			} else if count[i] >= 2 && count[PAIVALUE_SMALL+i] > 0 {
				//fmt.Println("case 2")
				count[i] -= 2
				count[PAIVALUE_SMALL+i] -= 1
				info = CanHu(huxi, count, len-3, zimoPaiValue)
				if info.canHu {
					//log.T("i: %v, value: %v", i, count[i])
					jiao := jiao{i, i, PAIVALUE_SMALL + i}
					info.jiaos = append(info.jiaos, jiao)
					//fmt.Println(fmt.Sprintf("找到一绞牌 case2 return info:[%+v]", info))
					return info
				}
				count[i] += 2
				count[PAIVALUE_SMALL+i] += 1

			}
		}
	} else {
		// 说明对牌出现
		for i := 1; i < TOTALPAIVALUE_COUNT+1; i++ {
			if count[i] >= 2 {
				count[i] -= 2
				info = CanHu(huxi, count, len-2, zimoPaiValue)
				if info.canHu {
					//log.T("i: %v, value: %v", i, count[i])
					info.jiang = int32(i)
					//fmt.Println(fmt.Sprintf("找到对牌 return info:[%+v]", info))
					return info
				}
				count[i] += 2
			}
		}
	}
	info.canHu = false
	//fmt.Println(fmt.Sprintf("长度不匹配 return info:[%+v]", info))
	return info
}

func TryHu2(gameData interface{}, checkPai interface{}, isDianPao bool) (interface{}, error) {

	userGameData := util.DeepClone(gameData.(*UserGameData)).(*UserGameData)
	pais := userGameData.HandPokers

	log.T("TryHu2(handPokers:[%v] checkPai:[%v] isDianPao:[%v])", Cards2String(pais), checkPai, isDianPao)
	//fmt.Println(fmt.Sprintf("TryHu2(handPokers:[%v] checkPai:[%v] isDianPao:[%v])", Cards2String(pais), checkPai, isDianPao))

	//找出牌组里原始坎牌
	var srcKan []*YiKanPai
	totalHuXi := int32(userGameData.RoundHuXi)
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

				//累加总胡息
				totalHuXi += kan.HuXi
				for _, p := range pokers {
					pais = DelPaiFromPokersByID(pais, p)
				}
				log.T("check胡牌时玩家手牌原始的坎牌是：[%v]", Cards2String(kan.Pais))
				//fmt.Println(fmt.Sprintf("check胡牌时玩家手牌原始的坎牌是：[%v]", Cards2String(kan.Pais)))
				srcKan = append(srcKan, kan)
			}
		}
	}

	//组合checkPai
	checkPokers := util.DeepClone(pais).([]*PHZPoker)

	zimoPaiValue := int32(-1)

	if checkPai != nil {
		v := reflect.ValueOf(checkPai)
		if !v.IsNil() && v.IsValid() {
			checkPokers = append(util.DeepClone(pais).([]*PHZPoker), checkPai.(*PHZPoker))
			if !isDianPao {
				//自摸的牌值
				zimoPaiValue = checkPai.(*PHZPoker).GetValue()
			}
		}
	}

	huInfo := &HuInfo{}

	huInfo.KanPais = append(huInfo.KanPais, srcKan...)

	//如果找完原始的坎牌，手里没有牌了，则牌型上是可以胡了
	if (checkPokers == nil || len(checkPokers) == 0) && totalHuXi >= CANHU_LIMIT_HUXI {
		huInfo.CanHu = true
		log.T("check玩家是否可以胡时，玩家手里只有坎牌，可以胡牌...")
		//fmt.Println(fmt.Sprintf("check玩家是否可以胡时，玩家手里只有坎牌，可以胡牌..."))
		return huInfo, nil
	}

	log.T("CanHu的CheckPokers是:[%v]", Cards2String(checkPokers))
	//fmt.Println(fmt.Sprintf("CanHu的CheckPokers是:[%v]", Cards2String(checkPokers)))
	canHuInfo := CanHu(totalHuXi, CountHandPais(checkPokers), len(checkPokers), zimoPaiValue)

	log.T("CanHu的CheckPokers结果是:[%+v]", canHuInfo)
	//fmt.Println(fmt.Sprintf("CanHu的CheckPokers结果是:[%+v]", canHuInfo))
	huInfo.CanHu = canHuInfo.canHu
	huInfo.HuXi = canHuInfo.totalHuXi

	//能胡的话需要将结构转化
	if huInfo.CanHu {
		//将牌
		if canHuInfo.jiang > -1 {
			jiangPais := GetPaisByValue2(checkPokers, canHuInfo.jiang)
			huInfo.DuiZis = append(huInfo.DuiZis, &DuiZi{Pais: jiangPais})

			for _, delPai := range jiangPais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的将牌:[%v] 删除后的checkPokers:[%v]", Cards2String(jiangPais), Cards2String(checkPokers)))
		}

		//一坎牌
		for _, kan := range canHuInfo.kans {
			kanPais := GetPaisByValue2(checkPokers, int32(kan))
			huInfo.KanPais = append(huInfo.KanPais, &YiKanPai{Pais: kanPais})
			for _, delPai := range kanPais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的一坎牌:[%v] 删除后的checkPokers:[%v]", Cards2String(kanPais), Cards2String(checkPokers)))
		}

		//一碰牌
		for _, kan := range canHuInfo.pengs {
			kanPais := GetPaisByValue2(checkPokers, int32(kan))
			huInfo.KanPais = append(huInfo.KanPais, &YiKanPai{Pais: kanPais})
			for _, delPai := range kanPais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			fmt.Println(fmt.Sprintf("TryHu2找到的一碰牌:[%v] 删除后的checkPokers:[%v]", Cards2String(kanPais), Cards2String(checkPokers)))
		}

		//一绞牌
		for _, jiao := range canHuInfo.jiaos {
			yjh := &YJH{}
			for _, jiaoPaiValue := range jiao {
				if pai := GetPaiByValue(checkPokers, int32(jiaoPaiValue)); pai != nil {
					yjh.Pais = append(yjh.Pais, pai)
				}
			}
			huInfo.YiJiaoPai = append(huInfo.YiJiaoPai, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的一绞牌:[%v] 删除后的checkPokers:[%v]", Cards2String(yjh.Pais), Cards2String(checkPokers)))
		}

		//一句话
		for _, yijuhua := range canHuInfo.yijuhuas {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(yijuhua)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(yijuhua+1)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(yijuhua+2)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的一句话:[%v] 删除后的checkPokers:[%v]", Cards2String(yjh.Pais), Cards2String(checkPokers)))
		}
		//大字一二三
		for i := canHuInfo.countBigYiErSan; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(11)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(12)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(13)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的大字一二三:[%v] 删除后的checkPokers:[%v]", Cards2String(yjh.Pais), Cards2String(checkPokers)))
		}

		//大字二七十
		for i := canHuInfo.countBigErQiShi; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(12)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(17)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(20)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的大字二七十:[%v] 删除后的checkPokers:[%v]", Cards2String(yjh.Pais), Cards2String(checkPokers)))
		}

		//小字一二三
		for i := canHuInfo.countSmallYiErSan; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(1)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字一:[%v]", Cards2String(yjh.Pais)))
			if pai := GetPaiByValue(checkPokers, int32(2)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字二:[%v]", Cards2String(yjh.Pais)))
			if pai := GetPaiByValue(checkPokers, int32(3)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字三:[%v]", Cards2String(yjh.Pais)))

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字一二三:[%v] 删除后的checkPokers:[%v]", Cards2String(yjh.Pais), Cards2String(checkPokers)))
		}

		//小字二七十
		for i := canHuInfo.countSmallErQiShi; i > 0; i-- {
			yjh := &YJH{}
			if pai := GetPaiByValue(checkPokers, int32(2)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(7)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}
			if pai := GetPaiByValue(checkPokers, int32(10)); pai != nil {
				yjh.Pais = append(yjh.Pais, pai)
			}

			huInfo.YiJuHua = append(huInfo.YiJuHua, yjh)
			for _, delPai := range yjh.Pais {
				checkPokers = DelPaiFromPokers(checkPokers, delPai)
			}
			//fmt.Println(fmt.Sprintf("TryHu2找到的小字二七十:[%v] 删除后的checkPokers:[%v]", Cards2String(yjh.Pais), Cards2String(checkPokers)))
		}

	}

	log.T("TryHu2的huInfo结果是:[%+v]", huInfo)
	//fmt.Println(fmt.Sprintf("TryHu2的huInfo结果是:[%+v]", huInfo))
	return huInfo, nil
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
