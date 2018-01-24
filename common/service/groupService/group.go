package groupService

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/log"
	"casino_common/common/service/agentService"
	"casino_common/common/service/roomAgent"
	"casino_common/common/userService"
	"casino_common/proto/ddproto"
	"casino_common/utils/db"
	"casino_common/utils/numUtils"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"sync"
	"time"
)

//所有群的集合
var Groups = map[int32]*GroupInfo{}

var UserGroupIds = map[uint32][]int32{}

//群结构
type GroupInfo struct {
	ddproto.GroupInfo //群结构
	sync.Mutex
}

//查询群列表
func GetUserGroups(uid uint32) (list []*GroupInfo, err error) {
	if groupids, ok := UserGroupIds[uid]; ok {
		for _, gid := range groupids {
			ginfo := GetGroupInfoById(gid)
			if ginfo != nil {
				list = append(list, ginfo)
			}
		}
		return
	}

	var ids = struct {
		Ids []int32
	}{}

	//从数据库查询
	err = db.C(tableName.DBT_GROUP_INFO).Pipe([]bson.M{
		bson.M{"$match": bson.M{"members.uid": uid}},
		bson.M{"$group": bson.M{"_id": nil, "ids": bson.M{"$addToSet": "$id"}}},
	}, &ids)
	if err != nil {
		return
	}

	//更新缓存
	UserGroupIds[uid] = ids.Ids

	for _, gid := range ids.Ids {
		ginfo := GetGroupInfoById(gid)
		if ginfo != nil {
			list = append(list, ginfo)
		}
	}
	return
}

//查询群
func GetGroupInfoById(gid int32) *GroupInfo {
	group, ok := Groups[gid]
	if ok {
		return group
	}

	res := ddproto.GroupInfo{}
	err := db.C(tableName.DBT_GROUP_INFO).Find(bson.M{
		"id": gid,
	}, &res)

	if err != nil {
		log.E("find groupinfo(gid=%d) err ", gid)
		return nil
	}

	new_group := &GroupInfo{
		GroupInfo: res,
	}
	Groups[res.GetId()] = new_group
	return new_group
}

//数据落地
func (g *GroupInfo) SaveToMongo() error {
	//保存到mongo
	err := db.C(tableName.DBT_GROUP_INFO).Upsert(bson.M{
		"id": g.GetId(),
	}, g.GroupInfo)
	return err
}

//群内广播
func (g *GroupInfo) BroadCast(msg interface{}) {
	for _, u := range g.GetMembers() {
		if u != nil {
			agentService.WriteMsg(u.GetUid(), msg)
		}
	}
}

//查询群成员
func (g *GroupInfo) GetMemberById(uid uint32) *ddproto.GroupMemberInfo {
	for _, u := range g.Members {
		if u.GetUid() == uid {
			return u
		}
	}
	return nil
}

//新增群成员
func (g *GroupInfo) AddMember(addUid uint32) error {
	g.Lock()
	g.Unlock()

	ex_member := g.GetMemberById(addUid)

	if ex_member != nil {
		return errors.New("已经在群里了")
	}

	user_info := userService.GetUserById(addUid)
	if user_info == nil {
		log.E("未找到用户%d，加群%d失败", addUid, g.GetId())
		return errors.New("加群失败，请稍后重试。")
	}

	new_member := &ddproto.GroupMemberInfo{
		Uid:      proto.Uint32(addUid),
		NickName: proto.String(user_info.GetNickName()),
		Remark:   proto.String(""),
		HeadImg:  proto.String(user_info.GetHeadUrl()),
		OpenId:   proto.String(user_info.GetOpenId()),
	}
	//加入成员
	g.Members = append(g.Members, new_member)
	//更新sync_id
	g.SyncId = proto.Int32(g.GetSyncId() + 1)
	//更新新成员的群列表
	if ugids, ok := UserGroupIds[addUid]; ok {
		ugids = append(ugids, g.GetId())
		UserGroupIds[addUid] = ugids
	}

	//发送入群广播
	bc_msg := &ddproto.HallGroupSendMsgBc{
		ResItem: &ddproto.GroupMsgItem{
			FromUid:     proto.Uint32(0),
			ToGroup:     proto.Int32(g.GetId()),
			SendTime:    proto.Int64(time.Now().Unix()),
			MsgType:     ddproto.GroupMsgType_group_msg_type_MemberIn.Enum(),
			TxtContent:  proto.String("欢迎 " + user_info.GetNickName() + " 进群！"),
			DeskContent: nil,
		},
	}
	g.BroadCast(bc_msg)
	//保存数据
	g.SaveToMongo()
	return nil
}

//群成员退出
func (g *GroupInfo) DelMember(delUid uint32) error {
	g.Lock()
	g.Unlock()

	ex_member := g.GetMemberById(delUid)

	if ex_member == nil {
		return errors.New("您未在该群中，退群失败。")
	}

	user_info := userService.GetUserById(delUid)
	if user_info == nil {
		log.E("未找到用户%d，退群%d失败", delUid, g.GetId())
		return errors.New("退群失败，请稍后重试。")
	}

	//发送出群广播
	bc_msg := &ddproto.HallGroupSendMsgBc{
		ResItem: &ddproto.GroupMsgItem{
			FromUid:     proto.Uint32(0),
			ToGroup:     proto.Int32(g.GetId()),
			SendTime:    proto.Int64(time.Now().Unix()),
			MsgType:     ddproto.GroupMsgType_group_msg_type_MemberOut.Enum(),
			TxtContent:  proto.String(user_info.GetNickName() + " 退出了群聊。"),
			DeskContent: nil,
		},
	}
	g.BroadCast(bc_msg)

	//删除成员
	for i, u := range g.Members {
		if u.GetUid() == delUid {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
			break
		}
	}
	//更新sync_id
	g.SyncId = proto.Int32(g.GetSyncId() + 1)
	//更新新成员的群列表
	if ugids, ok := UserGroupIds[delUid]; ok {
		for i, id := range ugids {
			if id == g.GetId() {
				ugids = append(ugids[:i], ugids[i+1:]...)
			}
		}
		UserGroupIds[delUid] = ugids
	}

	//保存数据
	g.SaveToMongo()
	return nil
}

//是否在群中
func (g *GroupInfo) IsMember(uid uint32) bool {
	for _, u := range g.GetMembers() {
		if u.GetUid() == uid {
			return true
		}
	}
	return false
}

//接收消息
func (g *GroupInfo) OnReceiveMsg(msg *ddproto.GroupMsgItem, agent gate.Agent) {
	if !g.IsMember(msg.GetFromUid()) {
		log.E("非群成员发消息：%v", msg)
		return
	}

	switch msg.GetMsgType() {
	case ddproto.GroupMsgType_group_msg_type_TXT, ddproto.GroupMsgType_group_msg_type_Voice, ddproto.GroupMsgType_group_msg_type_Image:
		log.T("群发消息:%v", msg)
		//文本消息，直接广播
		bc := &ddproto.HallGroupSendMsgBc{
			ResItem: msg,
		}
		bc.ResItem.SendTime = proto.Int64(time.Now().Unix())
		bc.ResItem.DeskContent = nil
		g.BroadCast(bc)
	case ddproto.GroupMsgType_group_msg_type_CreateDesk:
		log.T("群创建房间:%v", msg)
		//创建房间
		game_option := msg.GetTxtContent()
		if len(g.GetGameOpts()) == 1 {
			game_option = g.GetGameOpts()[0].GetOption()
		}
		new_desk, err := roomAgent.CreateDeskByOption(g.GetOwner(), msg.GetToGroup(), game_option)
		//开房失败
		if err != nil {
			res := &ddproto.HallGroupSendMsgBc{
				ResItem: &ddproto.GroupMsgItem{
					FromUid:     proto.Uint32(msg.GetFromUid()),
					ToGroup:     proto.Int32(msg.GetToGroup()),
					SendTime:    proto.Int64(time.Now().Unix()),
					MsgType:     ddproto.GroupMsgType_group_msg_type_SysFail.Enum(),
					TxtContent:  proto.String("创建房间失败！错误:" + err.Error()),
					DeskContent: nil,
				},
			}
			agent.WriteMsg(res)
			return
		}
		//开房成功
		res := &ddproto.HallGroupSendMsgBc{
			ResItem: &ddproto.GroupMsgItem{
				FromUid:     proto.Uint32(msg.GetFromUid()),
				ToGroup:     proto.Int32(msg.GetToGroup()),
				SendTime:    proto.Int64(time.Now().Unix()),
				MsgType:     ddproto.GroupMsgType_group_msg_type_CreateDesk.Enum(),
				TxtContent:  proto.String("创建房间成功"),
				DeskContent: new_desk,
			},
		}
		g.BroadCast(res)
	case ddproto.GroupMsgType_group_msg_type_MemberOut:
		//退出群聊,房主踢人,解散房间
		if msg.GetFromUid() == g.GetOwner() {
			if del_member_id := numUtils.String2Uint32(msg.GetTxtContent()); g.IsMember(del_member_id) {
				//群主踢人
				g.DelMember(del_member_id)
			}
			if strings.ToUpper(msg.GetTxtContent()) == "DESTROY" {
				//群主解散群
				log.T("群主%d解散群成功！", g.GetOwner())
			}
		} else {
			//成员自己退出
			g.DelMember(msg.GetFromUid())
		}
	}
}
