package gameService

import (
	"casino_common/utils/db"
	"gopkg.in/mgo.v2/bson"
	"casino_common/common/consts/tableName"
	"errors"
)

type GameStatus int32

const (
	Stoped GameStatus = 0  //停机中
	Runing GameStatus = 1  //运行中，可正常服务

	Maintain GameStatus = 2  //处于维护中
	Accident GameStatus = 3  //处于故障中
)

//游戏服务器信息
type TGameServerInfoRow struct {
	ObjId            bson.ObjectId  `bson:"_id"`  //objectId
	GameId           int32	//所属游戏id
	DeviceId         int32  //机器id 每个GameServer对应一个唯一id
	WsAddr           int32	//对外websocket地址
	MaxConn          int32  //最大承载连接数
	Status           GameStatus  //服务器状态
	Tag              string  //标签，方便标识与分类
	CurrGamer        int32  //当前人数
}

func FindServerByGameAndDevice(game_id int32, device_id int32) *TGameServerInfoRow {
	row := new(TGameServerInfoRow)
	err := db.C(tableName.DBT_GAME_SERVER_INFO).Find(bson.M{
		"gameid": game_id,
		"deviceid": device_id,
	}, row)
	if err != nil {
		return nil
	}
	return row
}

//通过game_id返回一个游戏服务器
func GetMappingServerByGameId(game_id int32) (*TGameServerInfoRow, error) {
	list := []*TGameServerInfoRow{}

	err, count := db.C(tableName.DBT_GAME_SERVER_INFO).Page(bson.M{
		"gameid": game_id,
		"status": Runing,
		"maxconn": bson.M{"$gt": 0},
	}, &list, "", 1, 999)

	if err != nil {
		return nil, err
	}

	if count <= 0 {
		return nil, errors.New("server not found.")
	}

	min_server := list[0]
	min_num := float64(list[0].CurrGamer)/float64(list[0].MaxConn)
	for _,item := range list {
		num := float64(item.CurrGamer)/float64(item.MaxConn)
		if num < min_num {
			min_num = num
			min_server = item
		}
	}

	return min_server, nil
}
