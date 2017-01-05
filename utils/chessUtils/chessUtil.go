package chessUtils

import (
	"casino_common/utils/rand"
	"casino_common/utils/numUtils"
	"fmt"
)

//随机一副扑克牌的index
func Xipai(indexStart int, paiCount int) []int32 {
	//初始化一个顺序的牌的集合
	pmap := make([]int32, paiCount)
	for i := 0; i < paiCount; i++ {
		pmap[i] = int32(i + indexStart)
	}

	//打乱牌的集合
	pResult := make([]int32, paiCount)
	for i := 0; i < paiCount; i++ {
		rand := rand.Rand(int32(0), int32(paiCount - i))
		pResult[i] = pmap[rand]
		pmap = append(pmap[:rand], pmap[rand + 1:]...)
	}
	return pResult
}

//生成房间号
func GetRoomPass(gameId int32) string {
	pre1 := rand.Rand(1, 10)
	pre2 := rand.Rand(1, 10)
	pre3 := rand.Rand(1, 10)
	pre4 := rand.Rand(1, 10)
	pre5 := rand.Rand(1, 10)
	var pre6 int32 = 0

	sum := pre1 + pre2 + pre3 + pre4 + pre5
	a := gameId - sum % 10
	if a > 0 {
		pre6 = a
	} else {
		pre6 = 10 + a
	}

	ret := pre1 * 100000 + pre2 * 10000 + pre3 * 1000 + pre4 * 100 + pre5 * 10 + pre6
	retSrt, _ := numUtils.Int2String(ret)
	return retSrt
}
