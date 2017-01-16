package userAttachDao

import "testing"

func TestFindUserAttachByUserId(t *testing.T) {
	a := FindUserAttachByKV("userid", 10009)
	println(a.UserId)
}
