package userService

import (
	"testing"
	"casino_common/common/model/userDao"
	"casino_common/common/cfg"
	"casino_common/test"
)

func init() {
	test.TestInit()
}

func TestINCRUserCOIN(t *testing.T) {
	//10055
	//10168
	//10169
	//10170
	b, e := DECRUserCOIN(13702, 970000)
	if e != nil {
		t.Logf("出错%v", e)
	} else {
		t.Logf("增加之后的金额%v", b)
	}
}

func TestGetUserDiamond(t *testing.T) {
	var userId uint32 = 13702

	mongUser := userDao.FindUserById(userId)
	t.Logf("用mongod user :  %v", mongUser)

	dia := GetUserDiamond(userId)
	c := GetUserRoomCard(userId)
	coin := GetUserCoin(userId)
	t.Logf("用户的钻石数量: %v,房卡的数量: %v,金币的数量: %v", dia, c, coin)

	muser := userDao.FindUserById(userId)
	t.Logf("mongo : 用户的钻石数量: %v,房卡的数量: %v,金币的数量:%v", muser.GetDiamond(), muser.GetRoomCard(), muser.GetCoin())

	ruser := GetUserById(userId)
	t.Logf("redis : 用户的钻石数量: %v,房卡的数量: %v 金币的数量:%v", ruser.GetDiamond(), ruser.GetRoomCard(), ruser.GetCoin())
}

func TestSyncUserMoney(t *testing.T) {
	userId := uint32(10025)
	user := GetUserById(userId)
	t.Logf("玩家[%v]的信息:%v", userId, user)
	SetUserMoney(userId, cfg.RKEY_USER_DIAMOND, 999)
	SyncReidsUserMoney(user)
	t.Logf("玩家2[%v]的信息:%v", userId, user)
}

func TestUserInfo(t *testing.T) {
	mongUser := userDao.FindUserByUnionId("oKKIfxG1i1CzbjDn-KlT50B3Pcxg")
	t.Log("user:", mongUser)
	//userDao.
}
