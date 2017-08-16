package testUtils

import (
	"testing"
	"casino_common/test"
)

func TestGetTestXiPai(t *testing.T) {
	ret := GetTestXiPai(1)
	t.Logf("最后的结果:%v", ret)
}

func init() {
	test.TestInit()
}

func TestSetDeskPoker(t *testing.T) {
	pokers := map[uint32][]int{
		23: []int{1,2,3,4,5},
		45: []int{1,4,7},
	}

	//写入redis
	err := SetDeskPreSendPokers(1,1, pokers)
	t.Log(err)

	//第一次查询redis
	redis_poker, err := GetDeskPreSendPokers(1, 1)
	t.Log(redis_poker, err)

	//再次查询redis
	redis_poker, err = GetDeskPreSendPokers(1, 1)
	t.Log(redis_poker, err)

}
