package phzData

import (
	"testing"
)

func TestInitPaiByIndex(t *testing.T) {
	indexs := []int{18, 17, 19, 20}
	pais := []*PHZPoker{}
	for _, i := range indexs {
		pai := InitPaiByIndex(int32(i))
		pais = append(pais, pai)
	}
	t.Logf("pais %v", Cards2String(pais))
	t.Logf("pais count %v", CountHandPais(pais))

	for _, pai := range pais {
		t.Logf("pai[%v] paiValue[%v]", pai.GetLogDes(), pai.GetValue())
	}
}
