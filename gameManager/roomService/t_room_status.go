package roomService

import (
	"gopkg.in/mgo.v2/bson"
	"casino_common/utils/db"
	"math/rand"
	"time"
	"fmt"
	"errors"
	"casino_common/common/consts/tableName"
	"casino_common/gameManager/gameService"
)

type RoomStatus int32

const (
	Waiting  RoomStatus = 1  //已创建,等待开局
	Gaming RoomStatus = 2  //游戏中
	Ended  RoomStatus = 3  //游戏已结束
)

//房间状态表
type TRoomInfoRow struct {
	ObjId  bson.ObjectId `bson:"_id"`  //objectid
	GameId	int32  //所属游戏id
	DeviceId  int32  //所属服务器
	RoomPwd  string  //房间号
	Status  RoomStatus  //房间状态
	Owner  uint32  //创建者
	CreateFee  int32  //房费

	MaxGamer  int32  //可容纳的玩家人数
	Users []uint32  //当前牌桌内用户
}

//更新进mongo
func (t *TRoomInfoRow) Upsert() error {
	return db.C(tableName.DBT_GAME_ROOM_INFO).Upsert(bson.M{
		"gameid": t.GameId,
		"deviceid": t.DeviceId,
		"roompwd": t.RoomPwd,
	}, t)
}

//删除
func (t *TRoomInfoRow) Remove() error {
	return db.C(tableName.DBT_GAME_ROOM_INFO).Remove(bson.M{
		"gameid": t.GameId,
		"deviceid": t.DeviceId,
		"roompwd": t.RoomPwd,
	})
}

//创建一个房间(该房间的唯一标识，以及房间号)
func CreateRoom(game_id int32, device_id int32, owner uint32, create_fee int32, max_gamer int32) (room_passwd string, err error) {
	exist_server := gameService.FindServerByGameAndDevice(game_id, device_id)
	if exist_server == nil {
		return "123456", errors.New("服务器不存在！")
	}

	new_room_passwd := 123456
	for {
		new_room_passwd = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899999) + 100000
		tmp_row := new(TRoomInfoRow)

		err = db.C(tableName.DBT_GAME_ROOM_INFO).Find(bson.M{"roompwd": new_room_passwd}, tmp_row)
		if err != nil {
			break
		}
	}

	new_room := TRoomInfoRow{
		ObjId: bson.NewObjectId(),
		GameId: game_id,
		DeviceId: device_id,
		RoomPwd: fmt.Sprintf("%d", new_room_passwd),
		Status: Waiting,
		Owner: owner,
		CreateFee: create_fee,
		MaxGamer: max_gamer,
		Users: []uint32{},
	}

	room_passwd = fmt.Sprintf("%d", new_room_passwd)
	err = new_room.Upsert()
	return
}

//查询room_info
func FindRoomByPasswd(room_passwd string) *TRoomInfoRow {
	room_row := new(TRoomInfoRow)
	err := db.C(tableName.DBT_GAME_ROOM_INFO).Find(bson.M{"roompwd": room_passwd}, room_row)
	if err != nil {
		return nil
	}
	return room_row
}

//更新房间状态
func UpdateRoomStatus(room_pwd string, room_status RoomStatus) error {
	if room_status != Ended {
		return db.C(tableName.DBT_GAME_ROOM_INFO).Update(bson.M{"roompwd": room_pwd}, bson.M{"$set": bson.M{"status": room_status}})
	}
	//如果是结束比赛
	room_row := FindRoomByPasswd(room_pwd)
	if room_row == nil {
		return errors.New("room not found.")
	}
	room_row.Status = Ended
	room_row.Upsert()
	//删除该条房间信息
	room_row.Remove()
	//更新该服游戏人数
	return db.C(tableName.DBT_GAME_SERVER_INFO).Update(bson.M{
		"gameid": room_row.GameId,
		"deviceid": room_row.DeviceId,
	},bson.M{
		"$inc": bson.M{"currgamer": -len(room_row.Users)},
	})
}

//新增用户
func NewUserEnter(room_pwd string, new_userid uint32) error {
	room_row := new(TRoomInfoRow)
	err := db.C(tableName.DBT_GAME_ROOM_INFO).Find(bson.M{"roompwd": room_pwd}, room_row)
	if err != nil {
		return err
	}

	for _,u := range room_row.Users {
		if u == new_userid {
			return errors.New("该用户已存在！")
		}
	}

	room_row.Users = append(room_row.Users, new_userid)

	db.C(tableName.DBT_GAME_SERVER_INFO).Update(bson.M{
		"gameid": room_row.GameId,
		"deviceid": room_row.DeviceId,
	},bson.M{
		"$inc": bson.M{"currgamer": 1},
	})

	return room_row.Upsert()
}

//通过房号查询房间
func FindServerInfoByRoomPWD(room_pwd string) (*gameService.TGameServerInfoRow, error) {
	room_info := new(TRoomInfoRow)
	err := db.C(tableName.DBT_GAME_ROOM_INFO).Find(bson.M{
		"roompwd": room_pwd,
	}, room_info)
	if err != nil {
		return nil, err
	}
	server_info := new(gameService.TGameServerInfoRow)
	err = db.C(tableName.DBT_GAME_SERVER_INFO).Find(bson.M{
		"gameid": room_info.GameId,
		"deviceid": room_info.DeviceId,
	}, server_info)

	return server_info, err
}
