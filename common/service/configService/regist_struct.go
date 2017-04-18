package configService

import "casino_common/proto/ddproto"

//存在mongo中的配置
type LoginConfigMongo struct {
	CurVersion      int32 `bson:"CurVersion" title:"当前版本号"`  //curVersion的配置
	BaseDownloadUrl string `bson:"BaseDownloadUrl" title:"默认下载地址"` //默认下载地址
}

//服务器信息
type LoginServerInfo struct {
	GameId ddproto.CommonEnumGame `bson:"GameId" title:"游戏Id"`
	Name string `bson:"name" title:"游戏"`
	CurVersion          int32 `bson:"CurVersion" form:"CurVersion" binding:"" title:"当前版本"`
	IsUpdate            int32 `bson:"IsUpdate" form:"IsUpdate" binding:"" title:"是否需要强制升级"`
	IsMaintain          int32 `bson:"IsMaintain" form:"IsMaintain" binding:"" title:"是否在维护中"`
	MaintainMsg         string `bson:"MaintainMsg" form:"MaintainMsg" binding:"" title:"的维护信息"`
	ReleaseTag          int32 `bson:"ReleaseTag" form:"ReleaseTag" binding:"" title:"已经发布的版本"`
	DownloadUrl         string `bson:"DownloadUrl" form:"DownloadUrl" binding:"" title:"的下载连接"`
	LatestClientVersion int32 `bson:"LatestClientVersion" form:"LatestClientVersion" binding:"" title:"最后的版本号"`
	IP                  string `bson:"IP" form:"IP" binding:"" title:"ip"`
	PORT                int32 `bson:"PORT" form:"PORT" binding:"" title:"端口"`
	STATUS              int32 `bson:"STATUS" form:"STATUS" binding:"" title:"状态码"`
}
