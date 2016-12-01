package chessUtils

import "casino_common/utils/rand"

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

