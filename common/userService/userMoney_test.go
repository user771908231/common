package userService

import (
	"casino_common/common/sys"
	"testing"
	"casino_common/common/model/userDao"
)

func init() {
	sys.InitRedis("127.0.0.1:6379", "test")
}

func TestINCRUserCOIN(t *testing.T) {
	b, e := DECRUserDiamond(1, 2)
	if e != nil {
		t.Logf("出错%v", e)
	} else {
		t.Logf("增加之后的金额%v", b)
	}
}

func TestGetUserDiamond(t *testing.T) {
	var userId uint32 = 10341
	dia := GetUserDiamond(userId)
	c := GetUserRoomCard(userId)
	t.Logf("用户的钻石数量: %v,房卡的数量: %v", dia, c)

	muser := userDao.FindUserById(userId)
	t.Logf("mongo : 用户的钻石数量: %v,房卡的数量: %v", muser.GetDiamond(), muser.GetRoomCard())

	ruser := GetUserById(userId)
	t.Logf("redis : 用户的钻石数量: %v,房卡的数量: %v", ruser.GetDiamond(), ruser.GetRoomCard())

}
