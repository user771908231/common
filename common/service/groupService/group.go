package groupService

import (
	"casino_common/proto/ddproto"
	"casino_common/common/service/agentService"
	"errors"
	"github.com/golang/protobuf/proto"
	"casino_common/common/userService"
	"casino_common/common/log"
	"time"
	"sync"
	"casino_common/utils/db"
	"casino_common/common/consts/tableName"
	"gopkg.in/mgo.v2/bson"
)

//所有群的集合
var Groups = map[int32]*GroupInfo{}

//群结构
type GroupInfo struct {
	ddproto.GroupInfo //群结构
	sync.Mutex
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
	for _,u := range g.GetMembers() {
		if u != nil {
			agent, ok := agentService.AgentMap[u.GetUid()]
			if ok && agent != nil {
				agent.WriteMsg(msg)
			}
		}
	}
}

//查询群成员
func (g *GroupInfo) GetMemberById(uid uint32) *ddproto.GroupMemberInfo {
	for _,u := range g.Members {
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
		Uid: proto.Uint32(addUid),
		NickName: proto.String(user_info.GetNickName()),
		Remark: proto.String(""),
		HeadImg: proto.String(user_info.GetHeadUrl()),
	}
	//加入成员
	g.Members = append(g.Members, new_member)
	//更新sync_id
	g.SyncId = proto.Int64(g.GetSyncId()+1)

	//发送入群广播
	bc_msg := &ddproto.HallGroupSendMsgBc{
		ResItem: &ddproto.GroupMsgItem{
			FromUid: proto.Uint32(0),
			ToGroup: proto.Int32(g.GetId()),
			SendTime: proto.Int64(time.Now().Unix()),
			MsgType: ddproto.GroupMsgType_group_msg_type_MemberIn.Enum(),
			TxtContent: proto.String("欢迎 "+user_info.GetNickName()+" 进群！"),
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

	//删除成员
	for i,u := range g.Members {
		if u.GetUid() == delUid {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
			break
		}
	}
	//更新sync_id
	g.SyncId = proto.Int64(g.GetSyncId()+1)

	//发送入群广播
	bc_msg := &ddproto.HallGroupSendMsgBc{
		ResItem: &ddproto.GroupMsgItem{
			FromUid: proto.Uint32(0),
			ToGroup: proto.Int32(g.GetId()),
			SendTime: proto.Int64(time.Now().Unix()),
			MsgType: ddproto.GroupMsgType_group_msg_type_MemberOut.Enum(),
			TxtContent: proto.String(user_info.GetNickName()+" 退出了群聊。"),
			DeskContent: nil,
		},
	}
	g.BroadCast(bc_msg)
	//保存数据
	g.SaveToMongo()
	return nil
}

