package phzData

import "testing"

func TestTryHu2(t *testing.T) {
	gameData := &UserGameData{}
	gameData.RoundHuXi = 15
	//indexs := []int32{69, 70, 71, 10, 44, 45, 46, 65, 66, 9, 1, 9, 17}
	//indexs := []int32{69, 70, 71, 10, 44, 45, 46, 65, 66, 67, 9, 1, 2, 3, 9, 10, 11, 17, 18, 19}
	indexs := []int32{1, 2, 5, 6, 3, 11, 19}
	pais := []*PHZPoker{}
	for _, i := range indexs {
		pai := InitPaiByIndex(i)
		if pai != nil {
			pais = append(pais, pai)
		}
	}
	//checkPai := InitPaiByIndex(int32(11))
	gameData.HandPokers = pais
	TryHu2(gameData, nil, false)
}