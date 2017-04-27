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

func TestGetRoomPassV2(t *testing.T) {
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(12))
	t.Log(GetRoomPassV2(12))
	t.Log(GetRoomPassV2(11))
	t.Log(GetRoomPassV2(39))
	t.Log(GetRoomPassV2(99))
	t.Log(GetRoomPassV2(1))
	t.Log(GetRoomPassV2(2))
	t.Log(GetRoomPassV2(3))
}
