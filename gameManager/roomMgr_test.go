package gameManager

import (
	"casino_common/test"
	"testing"
	"casino_common/gameManager/roomService"
	"casino_common/proto/ddproto"
	"casino_common/gameManager/gameService"
	"time"
)

func init() {
	test.TestInit()
}

const (
	roomPwd = "385802"
)

func TestNewRoom(t *testing.T) {
	passwd, err := roomService.CreateRoom(int32(ddproto.CommonEnumGame_GID_NIUNIUJINGDIAN), 1, 12, 2, 6)
	t.Log(passwd, err)
}

func TestUpdateRoomStatus(t *testing.T) {
	err := roomService.UpdateRoomStatus(roomPwd, roomService.Gaming)
	t.Log(err)
}

func TestAddNewUser(t *testing.T) {
	err := roomService.NewUserEnter(roomPwd, 33)
	t.Log(err)
}

func TestGameEnd(t *testing.T) {
	err := roomService.UpdateRoomStatus(roomPwd, roomService.Ended)
	t.Log(err)
}


//服务器负载均衡
func TestMappingService(t *testing.T) {
	time_start := time.Now()
	server, err := gameService.GetMappingServerByGameId(11)
	t.Log(server, err)
	time_spend := time.Now().Sub(time_start).Seconds()*1000
	t.Logf("spemd %.2fms", time_spend)
}

//通过房号查询所在对应的服务器
func TestFindServerByPWD(t *testing.T) {
	server_info, err := roomService.FindServerInfoByRoomPWD(roomPwd)
	t.Log(server_info, err)
}
