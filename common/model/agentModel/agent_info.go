package agentModel

import (
	"casino_common/common/consts/tableName"
	"casino_common/common/model/userDao"
	"casino_common/utils/db"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type AgentType int

const (
	AGENT_TYPE_1 AgentType = 1 //一级总代理
	AGENT_TYPE_2 AgentType = 2 //一级普通代理
	AGENT_TYPE_3 AgentType = 3 //普通代理
)

//代理信息表
type AgentInfo struct {
	Id       bson.ObjectId `bson:"_id"`
	UserId   uint32        //游戏内user_id
	NickName string        //昵称
	RealName string        //真实姓名
	Phone    string        //手机号
	OpenId   string
	UnionId  string
	RootId   uint32    //所属根代理id
	Pid      uint32    //所属上级代理id
	Level    int32     //代理级别：一级 二级 三级
	Type     AgentType //代理类型： 总代 普通带理

	Attr struct {
		Level int32 //1：高级代理 2：中级代理 3：低级代理  (黄圣特有)
	}

	CoinFeeRebate float64 //金币房费提成(单位：人民币元)
}

//插入表
func (r *AgentInfo) Insert() error {
	r.Id = bson.NewObjectId()
	return db.C(tableName.DBT_AGENT_INFO).Insert(r)
}

//保存
func (r *AgentInfo) Save() error {
	return db.C(tableName.DBT_AGENT_INFO).Update(bson.M{
		"_id": r.Id,
	}, r)
}

//自增金币房费
func (r *AgentInfo) Inc(key string, val float64) error {
	return db.C(tableName.DBT_AGENT_INFO).Update(bson.M{"_id": r.Id}, bson.M{
		"$inc": bson.M{key: val},
	})
}

//通过Agent_id获取agent_info
func GetAgentInfoById(agent_id uint32) *AgentInfo {
	info := new(AgentInfo)
	err := db.C(tableName.DBT_AGENT_INFO).Find(bson.M{
		"userid": agent_id,
	}, info)
	if err != nil {
		return nil
	}
	return info
}

//判断一个用户是否为代理商
func IsAgent(user_id uint32) bool {
	row := new(ApplyRecord)
	err := db.C(tableName.DBT_AGENT_INFO).Find(bson.M{
		"userid": user_id,
	}, row)
	if err == nil {
		return true
	}
	return false
}

//获取代理下面的子代理数量
func GetAgentChildNum(agent_id uint32) int {
	num, _ := db.C(tableName.DBT_AGENT_INFO).Count(bson.M{
		"pid": agent_id,
	})
	return num
}

//获取子代理列表
func GetAgentChildrens(agent_id uint32) []*AgentInfo {
	childs := []*AgentInfo{}

	db.C(tableName.DBT_AGENT_INFO).FindAll(bson.M{
		"pid": agent_id,
	}, &childs)

	return childs
}

//获取子代理的所有多级子代理列表
func GetAgentChildsTree(agent_id uint32) []*AgentInfo {
	list := GetAgentChildrens(agent_id)
	for _, agent := range list{
		list = append(list, GetAgentChildsTree(agent.UserId)...)
	}
	return list
}

//=======设置代理的上级========
func SetAgentParent(agent_id uint32, new_pid uint32) error {
	agent_info := GetAgentInfoById(agent_id)
	if agent_info == nil {
		return errors.New("代理不存在！")
	}
	switch new_pid {
	case 1:
		agent_info.Pid = 0
		agent_info.RootId = 0
		agent_info.Level = 1
		agent_info.Type = AGENT_TYPE_1
	case 2:
		agent_info.Pid = 0
		agent_info.RootId = 0
		agent_info.Level = 1
		agent_info.Type = AGENT_TYPE_2
	default:
		parent_info := GetAgentInfoById(new_pid)
		if parent_info == nil {
			return errors.New("未找到父级代理！")
		}
		agent_info.Pid = parent_info.UserId
		if parent_info.Level == 1 {
			agent_info.RootId = parent_info.UserId
		} else {
			agent_info.RootId = parent_info.RootId
		}

		agent_info.Level = parent_info.Level + 1
		agent_info.Type = AGENT_TYPE_3
	}
	//更新数据
	agent_info.Save()

	//更新子代理的下级代理关系
	childs := GetAgentChildrens(agent_info.UserId)
	for _, child := range childs {
		SetAgentParent(child.UserId, agent_info.UserId)
	}

	return nil
}

//代理登录
func LoginByPwd(agent_id uint32, pwd string) (*AgentInfo, error) {
	agent_info := GetAgentInfoById(agent_id)
	user_info := userDao.FindUserById(agent_id)

	if agent_info == nil || user_info == nil {
		return nil, errors.New("代理不存在！")
	}

	if user_info.GetPwd() == "" || pwd == "" {
		return nil, errors.New("密码为空")
	}

	if user_info.GetPwd() == pwd {
		return agent_info, nil
	}
	return nil, errors.New("密码错误！")
}
