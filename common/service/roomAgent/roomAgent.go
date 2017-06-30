package roomAgent

import (
	"casino_common/proto/ddproto"
	"errors"
	"github.com/golang/protobuf/proto"
	"casino_common/utils/redisUtils"
	"fmt"
	"casino_common/common/consts"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
)

//获取代理创建的房间列表
func GetAgentRooms(creator uint32) []*ddproto.CommonDeskByAgent {
	redis_key := fmt.Sprintf("%s_%d", consts.RKEY_AGENT_CREATE_DESK_LIST, creator)
	redis_obj := &ddproto.RedisDeskByAgent{}

	obj := redisUtils.GetObj(redis_key, redis_obj)
	if obj == nil {
		return []*ddproto.CommonDeskByAgent{}
	}

	return redis_obj.Data
}

//更新redis
func updateToRedis(creator uint32, gameId int32, deskId int32, item_list []*ddproto.CommonDeskByAgent) error {

	redis_key := fmt.Sprintf("%s_%d", consts.RKEY_AGENT_CREATE_DESK_LIST, creator)

	return redisUtils.SetObj(redis_key, &ddproto.RedisDeskByAgent{Data: item_list})
}

//插入mongo
func insertToMongo(item *ddproto.CommonDeskByAgent) error {

 	return db.C(tableName.DBT_AGENT_CREATE_ROOM_RECORD).Upsert(bson.M{"creator": item.GetCreator(), "gameid": item.GetGameId(), "deskid": item.GetDeskId()}, item)
}

//检查合法性
func checkItem(item *ddproto.CommonDeskByAgent) error {
	if item.GetPassword() == "" || item.GetGameId() == 0 || item.GetDeskId() == 0 || item.GetCreator() == 0 || item.GetCreateTime() == 0 {
		return errors.New("参数不合法！")
	}
	return nil
}

//创建房间
func CreateDesk(gameId int32, password string, deskId int32, creator uint32, tips string, createTime int64) error {
	item_list := GetAgentRooms(creator)
	for _, item := range item_list {
		if item.GetCreator() == creator && item.GetGameId() == gameId && item.GetDeskId() == deskId {
			//如果已存在，则更新到redis
			item.Password = &password
			item.Tips = &tips
			item.CreateTime = &createTime
			if err := checkItem(item); err != nil {
				return err
			}
			return updateToRedis(creator, gameId, deskId, item_list)
		}
	}

	new_item := &ddproto.CommonDeskByAgent{
		Password: &password,
		GameId: &gameId,
		DeskId: &deskId,
		Creator: &creator,
		Tips: &tips,
		CreateTime: &createTime,
		Status: proto.Int32(0),
		Users: []string{},
	}


	if err := checkItem(new_item); err != nil {
		return err
	}

	item_list = append(item_list, new_item)
	return updateToRedis(creator, gameId, deskId, item_list)
}

//牌桌开局
func DoStart(creator uint32, gameId int32, deskId int32) error {
	item_list := GetAgentRooms(creator)
	for _, item := range item_list {
		if item.GetCreator() == creator && item.GetGameId() == gameId && item.GetDeskId() == deskId {
			item.Status = proto.Int32(1)
			return updateToRedis(creator, gameId, deskId, item_list)
		}
	}
	return errors.New("item not found.")
}

//牌桌结束
func DoEnd(creator uint32, gameId int32, deskId int32) error {
	item_list := GetAgentRooms(creator)
	for i, item := range item_list {
		if item.GetCreator() == creator && item.GetGameId() == gameId && item.GetDeskId() == deskId {
			item.Status = proto.Int32(2)
			//先把历史开房记录存入mongo
			if err := insertToMongo(item); err != nil {
				return err
			}
			//再删除并保存进redis
			item_list = append(item_list[:i], item_list[i+1:]...)
			return updateToRedis(creator, gameId, deskId, item_list)
		}
	}
	return errors.New("item not found.")
}

//添加用户
func DoAddUser(creator uint32, gameId int32, deskId int32, new_user string) error {
	item_list := GetAgentRooms(creator)
	for _, item := range item_list {
		if item.GetCreator() == creator && item.GetGameId() == gameId && item.GetDeskId() == deskId {
			if new_user == "" {
				return errors.New("user is empty.")
			}
			for _, u := range item.Users {
				if u == new_user {
					return errors.New("user is exist.")
				}
			}
			item.Users = append(item.Users, new_user)
			return updateToRedis(creator, gameId, deskId, item_list)
		}
	}
	return errors.New("item not found.")
}
