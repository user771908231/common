package chessUtils

import "testing"

func TestGetRoomPass(t *testing.T) {
	a := GetRoomPass(5)
	t.Logf("a :%v-%v", a)
	b := GetRoomPass(3)
	t.Logf("b :%v-%v", b)
	c := GetRoomPass(3)
	t.Logf("c :%v-%v", c)
	d := GetRoomPass(9)
	t.Logf("d :%v-%v", d)
}
