package chessUtils

import (
	"casino_common/utils/numUtils"
	"casino_common/utils/rand"
	mrand "math/rand"
	"time"
)

func init() {
	//洗牌随机种子初始化
	shuffle_rand_pointer = mrand.New(mrand.NewSource(time.Now().UnixNano()))
}

var shuffle_rand_pointer *mrand.Rand

//随机一副扑克牌的index
func Xipai(indexStart int, paiCount int) []int32 {
	//初始化一个顺序的牌的集合
	pmap := make([]int32, paiCount)
	for i := 0; i < paiCount; i++ {
		pmap[i] = int32(i + indexStart)
	}

	//打乱牌的集合
	return Shuffle2(Shuffle2(pmap))
}

//将一个slice数值打乱
func Shuffle(vals []int32) []int32 {
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))
	ret := make([]int32, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

//新版洗牌算法
func Shuffle2(vals []int32) []int32 {
	ret := make([]int32, len(vals))

	for i, _ := range ret {
		rand_index := shuffle_rand_pointer.Intn(len(vals))
		ret[i] = vals[rand_index]
		vals = append(vals[:rand_index], vals[rand_index+1:]...)
	}

	return ret
}

//生成房间号
func GetRoomPassV2(gameId int32) string {
	pre1 := rand.Rand(1, 10)
	pre2 := rand.Rand(1, 10)
	pre3 := rand.Rand(1, 10)
	pre4 := rand.Rand(1, 10)
	pre5 := rand.Rand(1, 10)
	var pre6 int32 = 0

	sum := pre1 + pre2 + pre3 + pre4 + pre5
	a := gameId - sum%10
	if a >= 0 {
		pre6 = a
	} else {
		pre6 = 10 + a
	}
	//log.T("pre1[%v], pre2[%v], pre3[%v], pre4[%v], pre5[%v]", pre1, pre2, pre3, pre4, pre5)
	//log.T("gameId[%v] sum[%v] a[%v] pre6[%v]", gameId, sum, a, pre6)
	ret := pre1*100000 + pre2*10000 + pre3*1000 + pre4*100 + pre5*10 + pre6
	retSrt, _ := numUtils.Int2String(ret)
	//log.T("ret[%v] retSrt[%v]", ret, retSrt)
	return retSrt
}

//生成6位纯随机的房间号
func GetRoomPassV3(gameId int32) string {
	pre1 := rand.Rand(1, 10)
	pre2 := rand.Rand(1, 10)
	pre3 := rand.Rand(1, 10)
	pre4 := rand.Rand(1, 10)
	pre5 := rand.Rand(1, 10)
	pre6 := rand.Rand(1, 10)

	ret := pre1*100000 + pre2*10000 + pre3*1000 + pre4*100 + pre5*10 + pre6
	retSrt, _ := numUtils.Int2String(ret)
	//log.T("ret[%v] retSrt[%v]", ret, retSrt)
	return retSrt
}

func GetRoomPass(gameId int32) string {
	r := rand.Rand(1000, 10000)
	rs, _ := numUtils.Int2String(r)
	gids, _ := numUtils.Int2String(gameId)
	if gameId < 10 {
		gids = "0" + gids
	}
	s1 := string([]rune(rs)[:2])
	s2 := string([]rune(rs)[2:])
	ret := s1 + gids + s2
	return ret
}
