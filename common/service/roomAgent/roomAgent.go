package roomAgent

import (
	"casino_common/common/consts"
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/service/rpcService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/redisUtils"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"casino_common/common/Error"
)

//从数据库获取当前未开局的房间表
func GetAgentRooms(creator uint32) []*ddproto.CommonDeskByAgent {
	redis_obj := &ddproto.RedisDeskByAgent{
		Data: []*ddproto.CommonDeskByAgent{},
	}
	redisUtils.GetObj(fmt.Sprintf("%s_%d", consts.RKEY_AGENT_CREATE_DESK_LIST, creator), redis_obj)

	log.T("GetAgentRooms(creator:%d):%v", creator, redis_obj.Data)
	return redis_obj.Data
}

//获取俱乐部代开的房间
func GetAgentRoomsByGroup(group_id int32) []*ddproto.CommonDeskByAgent {
	redis_obj := &ddproto.RedisDeskByAgent{
		Data: []*ddproto.CommonDeskByAgent{},
	}
	redisUtils.GetObj(fmt.Sprintf("%s_%d", consts.RKEY_AGENT_GROUP_CREATE_DESK_LIST, group_id), redis_obj)

	log.T("GetAgentRoomsByGroup(group_id:%d):%v", group_id, redis_obj.Data)
	return redis_obj.Data
}

//查询空闲房间
func GetAgentFreeRoomByOption(group_id int32, gamer_num int, keywords []string) *ddproto.CommonDeskByAgent {
	rooms := GetAgentRoomsByGroup(group_id)
	for _, room := range rooms {
		log.T("room status[%v] room len(users)[%v] gamer_num[%v]", room.GetStatus(), len(room.GetUsers()), gamer_num)
		if room.GetStatus() != 0 || len(room.GetUsers()) >= gamer_num {
			log.W("匹配失败 room[%v]", room.GetTips())
			continue
		}

		has_all_keywords := true
		log.T("开始匹配 roomTips[%v], keywords[%v]", room.GetTips(), keywords)
		for _, keyword := range keywords {
			log.T("开始匹配 roomTips[%v], keyword[%v]", room.GetTips(), keyword)
			if !strings.Contains(room.GetTips(), keyword) {
				log.W("匹配失败 roomTips[%v], keyword[%v]", room.GetTips(), keyword)
				has_all_keywords = false
				break
			}
		}

		if has_all_keywords {
			log.T("匹配成功 return room[%v]", room.GetTips())
			return room
		}
	}
	return nil
}

//查询单个牌桌信息
func saveToRedis(creator uint32, gameId int32, deskId int32, groupId int32) error {

	//保存至redis
	saveToRedisByCreator(creator)
	if groupId > 0 {
		saveToRedisByGroupid(groupId)
	} else if groupId == 0 {
		desk := new(ddproto.CommonDeskByAgent)

		err := db.C(tableName.DBT_AGENT_CREATED_ROOM).Find(bson.M{
			"creator": creator,
			"gameid":  gameId,
			"deskid":  deskId,
		}, desk)

		if err != nil {
			log.E("saveToRedis(creator:%d,gameId:%d,deskId:%d) err:%v", creator, gameId, deskId, err)
			return err
		}
		if desk.GetGroupId() > 0 {
			saveToRedisByGroupid(desk.GetGroupId())
		}
	}

	return nil
}

//更新redis
func saveToRedisByCreator(creator uint32) error {
	list := []*ddproto.CommonDeskByAgent{}

	err := db.C(tableName.DBT_AGENT_CREATED_ROOM).FindAll(bson.M{
		"creator": creator,
	}, &list)

	if err != nil {
		log.E("saveToRedisByCreator(%d) err:%v", creator, err)
		return err
	}

	redis_obj := &ddproto.RedisDeskByAgent{
		Data: list,
	}

	err = redisUtils.SetObj(fmt.Sprintf("%s_%d", consts.RKEY_AGENT_CREATE_DESK_LIST, creator), redis_obj)
	if err != nil {
		log.E("saveToRedisByCreator(%d) err:%v", creator, err)
		return err
	}

	return nil
}

//更新redis
func saveToRedisByGroupid(group_id int32) error {
	list := []*ddproto.CommonDeskByAgent{}

	err := db.C(tableName.DBT_AGENT_CREATED_ROOM).FindAll(bson.M{
		"groupid": group_id,
	}, &list)

	if err != nil {
		log.E("saveToRedisByGroupid(%d) err:%v", group_id, err)
		return err
	}

	redis_obj := &ddproto.RedisDeskByAgent{
		Data: list,
	}

	err = redisUtils.SetObj(fmt.Sprintf("%s_%d", consts.RKEY_AGENT_GROUP_CREATE_DESK_LIST, group_id), redis_obj)
	if err != nil {
		log.E("saveToRedisByGroupid(%d) err:%v", group_id, err)
		return err
	}

	return nil
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

//============================================房间事件===============================================

//创建房间--from-mongo
func CreateDesk(gameId int32, password string, deskId int32, creator uint32, tips string, createTime int64, maxCircle int32, maxGammer int32, groupId int32) error {
	count, err := db.C(tableName.DBT_AGENT_CREATED_ROOM).Count(bson.M{"creator": creator})
	if err != nil {
		return nil
	}
	if count >= 20 {
		return errors.New("您最多只能创建20个房间！")
	}

	new_item := &ddproto.CommonDeskByAgent{
		Password:   &password,
		GameId:     &gameId,
		DeskId:     &deskId,
		Creator:    &creator,
		Tips:       &tips,
		CreateTime: &createTime,
		Status:     proto.Int32(0),
		Users:      []string{},
		MaxCircle:  proto.Int32(maxCircle),
		MaxGammer:  proto.Int32(maxGammer),
		GroupId:    proto.Int32(groupId),
	}

	if err := checkItem(new_item); err != nil {
		return err
	}

	switch {
	case db.C(tableName.DBT_AGENT_CREATED_ROOM).Insert(new_item) != nil:
		return errors.New("insert err.")
	case saveToRedisByCreator(creator) != nil:
		return errors.New("save redis creator err.")
	case groupId > 0 && saveToRedisByGroupid(groupId) != nil:
		return errors.New("save redis group err.")
	}

	return nil
}

//牌桌开局
func DoStart(creator uint32, gameId int32, deskId int32) error {
	err := db.C(tableName.DBT_AGENT_CREATED_ROOM).Update(bson.M{
		"creator": creator,
		"gameid":  gameId,
		"deskid":  deskId,
	}, bson.M{
		"$set": bson.M{"status": 1},
	})

	if err != nil {
		return err
	}

	return saveToRedis(creator, gameId, deskId, 0)
}

//牌桌结束
func DoEnd(creator uint32, gameId int32, deskId int32) error {
	ex_desk := new(ddproto.CommonDeskByAgent)
	query := bson.M{
		"creator": creator,
		"gameid":  gameId,
		"deskid":  deskId,
	}
	err := db.C(tableName.DBT_AGENT_CREATED_ROOM).Find(query, ex_desk)

	if err != nil {
		return err
	}

	ex_desk.Status = proto.Int32(2)
	//删除在线房间记录
	err = db.C(tableName.DBT_AGENT_CREATED_ROOM).Remove(query)
	if err != nil {
		return err
	}

	//更新redis
	saveToRedisByCreator(ex_desk.GetCreator())
	if ex_desk.GetGroupId() > 0 {
		saveToRedisByGroupid(ex_desk.GetGroupId())
	}

	//插入已结束表
	insertToMongo(ex_desk)

	//推送解散消息
	_, err = rpcService.GetHall().SendDeskEventMsg(context.Background(), &ddproto.HallRpcDeskEventMsg{
		Msg:  proto.String("房间已结束游戏"),
		Desk: ex_desk,
	})

	return err
}

//牌桌解散
func DoDissolve(creator uint32, gameId int32, deskId int32) error {
	ex_desk := new(ddproto.CommonDeskByAgent)
	query := bson.M{
		"creator": creator,
		"gameid":  gameId,
		"deskid":  deskId,
	}
	err := db.C(tableName.DBT_AGENT_CREATED_ROOM).Find(query, ex_desk)

	if err != nil {
		return err
	}
	//删除在线房间记录
	err = db.C(tableName.DBT_AGENT_CREATED_ROOM).Remove(query)

	if err != nil {
		return err
	}

	//更新状态
	ex_desk.Status = proto.Int32(3)
	err = insertToMongo(ex_desk)
	if err != nil {
		return err
	}

	//更新redis
	err = saveToRedisByCreator(ex_desk.GetCreator())
	if err != nil {
		return err
	}
	if ex_desk.GetGroupId() > 0 {
		saveToRedisByGroupid(ex_desk.GetGroupId())

		//推送解散消息
		res, err := rpcService.GetHall().SendDeskEventMsg(context.Background(), &ddproto.HallRpcDeskEventMsg{
			Msg:  proto.String("房间已解散"),
			Desk: ex_desk,
		})

		if res == nil {
			if err != nil {
				return errors.New("rpc networkerr:" + err.Error())
			} else {
				return errors.New("未知错误")
			}
		} else {
			if res.GetCode() < 0 {
				return errors.New(res.GetError())
			} else {
				return nil
			}
		}
	}
	return nil
}

//添加用户
func DoAddUser(creator uint32, gameId int32, deskId int32, new_user string) error {
	query := bson.M{
		"creator": creator,
		"gameid":  gameId,
		"deskid":  deskId,
	}

	//更新redis
	desk := new(ddproto.CommonDeskByAgent)
	err := db.C(tableName.DBT_AGENT_CREATED_ROOM).Find(query, desk)
	if err != nil {
		return err
	}

	for _, u := range desk.Users {
		if u == new_user {
			return errors.New("用户已在房间里，重复请求进房。")
		}
	}
	desk.Users = append(desk.Users, new_user)

	//更新用户列表
	err = db.C(tableName.DBT_AGENT_CREATED_ROOM).Update(query, bson.M{
		"$addToSet": bson.M{"users": new_user},
	})
	if err != nil {
		return err
	}

	//保存到redis
	group_id := desk.GetGroupId()
	if group_id <= 0 {
		group_id = -1
	}
	err = saveToRedis(creator, gameId, deskId, group_id)
	if err != nil {
		return err
	}

	go func() {
		_, err = rpcService.GetHall().SendDeskEventMsg(context.Background(), &ddproto.HallRpcDeskEventMsg{
			Msg:  proto.String("新用户进房"),
			Desk: desk,
		})
	}()

	//更新Redis
	return nil
}

//删除用户
func DoRmUser(creator uint32, gameId int32, deskId int32, del_user string) error {
	query := bson.M{
		"creator": creator,
		"gameid":  gameId,
		"deskid":  deskId,
	}

	//更新redis
	desk := new(ddproto.CommonDeskByAgent)
	err := db.C(tableName.DBT_AGENT_CREATED_ROOM).Find(query, desk)
	if err != nil {
		return err
	}

	for i, u := range desk.Users {
		if u == del_user {
			desk.Users = append(desk.Users[:i], desk.Users[i+1:]...)
			break
		}
	}

	//更新用户列表
	err = db.C(tableName.DBT_AGENT_CREATED_ROOM).Update(query, bson.M{
		"$pull": bson.M{"users": del_user},
	})
	if err != nil {
		return err
	}

	//保存到redis
	group_id := desk.GetGroupId()
	if group_id <= 0 {
		group_id = -1
	}
	err = saveToRedis(creator, gameId, deskId, group_id)
	if err != nil {
		return err
	}

	go func() {
		defer Error.ErrorRecovery("rpcService.GetHall().SendDeskEventMsg()")
		_, err = rpcService.GetHall().SendDeskEventMsg(context.Background(), &ddproto.HallRpcDeskEventMsg{
			Msg:  proto.String("新用户进房"),
			Desk: desk,
		})
	}()

	//更新Redis
	return nil
}
