package majiangMsgRegister

import (
	"github.com/name5566/leaf/network/protobuf"
	"casino_common/proto/ddproto/mjproto"
	"casino_common/proto/ddproto"
)

func OnInit(mjProcessor *protobuf.Processor) {
	mjProcessor.Register(&ddproto.Heartbeat{})                   //0
	mjProcessor.Register(&mjproto.Game_QuickConn{})              //1 接入服务器
	mjProcessor.Register(&mjproto.Game_AckQuickConn{})           //2
	mjProcessor.Register(&ddproto.CommonReqGameLogin{})          //3 登录游戏
	mjProcessor.Register(&ddproto.CommonAckGameLogin{})          //4
	mjProcessor.Register(&mjproto.Game_CreateRoom{})             //5 创建房间
	mjProcessor.Register(&mjproto.Game_AckCreateRoom{})          //6
	mjProcessor.Register(&mjproto.Game_EnterRoom{})              //7 进入房间
	mjProcessor.Register(&mjproto.Game_AckEnterRoom{})           //8
	mjProcessor.Register(&mjproto.Game_SendGameInfo{})           //9 卓内游戏数据
	mjProcessor.Register(&mjproto.Game_Ready{})                  //10 准备
	mjProcessor.Register(&mjproto.Game_AckReady{})               //11
	mjProcessor.Register(&mjproto.Game_ExchangeCards{})          //12 换3张
	mjProcessor.Register(&mjproto.Game_AckExchangeCards{})       //13 换3张-回复
	mjProcessor.Register(&mjproto.Game_DingQue{})                //14 定缺
	mjProcessor.Register(&mjproto.Game_Opening{})                //15 开始(表示都已经准备完了)
	mjProcessor.Register(&mjproto.Game_DealCards{})              //16 发牌
	mjProcessor.Register(&mjproto.Game_GetInCard{})              //17 摸牌
	mjProcessor.Register(&mjproto.Game_SendOutCard{})            //18 出牌
	mjProcessor.Register(&mjproto.Game_AckSendOutCard{})         //19 出牌-ack
	mjProcessor.Register(&mjproto.Game_ActPeng{})                //20 碰
	mjProcessor.Register(&mjproto.Game_AckActPeng{})             //21
	mjProcessor.Register(&mjproto.Game_ActGang{})                //22 杠
	mjProcessor.Register(&mjproto.Game_AckActGang{})             //23
	mjProcessor.Register(&mjproto.Game_ActGuo{})                 //24 过
	mjProcessor.Register(&mjproto.Game_AckActGuo{})              //25
	mjProcessor.Register(&mjproto.Game_ActHu{})                  //26 胡
	mjProcessor.Register(&mjproto.Game_AckActHu{})               //27
	mjProcessor.Register(&mjproto.Game_BroadcastBeginDingQue{})  //28 开始定缺(广播)
	mjProcessor.Register(&mjproto.Game_BroadcastBeginExchange{}) //29 开始换牌(广播)
	mjProcessor.Register(&mjproto.Game_OverTurn{})               //30 轮到下一人
	mjProcessor.Register(&mjproto.Game_SendCurrentResult{})      //31 本局结果
	mjProcessor.Register(&mjproto.Game_SendEndLottery{})         //32 牌局结束
	mjProcessor.Register(&mjproto.Game_DissolveDesk{})           //33 解散房间
	mjProcessor.Register(&mjproto.Game_AckDissolveDesk{})        //34
	mjProcessor.Register(&ddproto.CommonReqLeaveDesk{})          //35 离开房间
	mjProcessor.Register(&ddproto.CommonAckLeaveDesk{})          //36
	mjProcessor.Register(&ddproto.CommonReqMessage{})            //37 发送聊天消息
	mjProcessor.Register(&ddproto.CommonBcMessage{})             //38 广播聊天
	mjProcessor.Register(&mjproto.Game_DingQueEnd{})             //39 定缺结束
	mjProcessor.Register(&mjproto.Game_GameRecord{})             //40 查询战绩
	mjProcessor.Register(&mjproto.Game_AckGameRecord{})          //41 战绩回复
	mjProcessor.Register(&mjproto.Game_ExchangeCardsEnd{})       //42 换三张 结束之后的广播
	mjProcessor.Register(&ddproto.CommonReqNotice{})             //43通知信息
	mjProcessor.Register(&ddproto.CommonAckNotice{})             //44通知信息回复
	mjProcessor.Register(&ddproto.CommonReqLogout{})             //45请求推出
	mjProcessor.Register(&ddproto.CommonAckLogout{})             //46回复请求推出
	//任务、在线奖励、托管(被动/主动）、换桌 道具
	mjProcessor.Register(&ddproto.AwardReqOnline{})                 //47 在线奖励
	mjProcessor.Register(&ddproto.WardAckOnline{})                  //48 在线奖励回复
	mjProcessor.Register(&ddproto.HallReqTask{})                    //49 任务
	mjProcessor.Register(&ddproto.HallAckTask{})                    //50 任务回复
	mjProcessor.Register(&ddproto.CommonReqEnterAgentMode{})        //51 进入托管
	mjProcessor.Register(&ddproto.CommonAckEnterAgentMode{})        //52 进入托管回复
	mjProcessor.Register(&ddproto.CommonReqQuitAgentMode{})         //53 退出托管
	mjProcessor.Register(&ddproto.CommonAckQuitAgentMode{})         //54 退出托管回复
	mjProcessor.Register(&ddproto.CommonReqReg{})                   //55 注册
	mjProcessor.Register(&ddproto.CommonAckReg{})                   //56 注册回复
	mjProcessor.Register(&ddproto.CommonReqGameState{})             //57 玩家游戏状态
	mjProcessor.Register(&ddproto.CommonAckGameState{})             //58 玩家游戏状态回复
	mjProcessor.Register(&ddproto.CommonReqFeedback{})              //59 反馈
	mjProcessor.Register(&ddproto.CommonReqApplyDissolve{})         //60 申请解散房间
	mjProcessor.Register(&ddproto.CommonBcApplyDissolve{})          //61 申请解散房间回复
	mjProcessor.Register(&ddproto.CommonReqApplyDissolveBack{})     //62 同意或拒绝解散房间回复
	mjProcessor.Register(&ddproto.CommonAckApplyDissolveBack{})     //63 同意或拒绝解散房间回复
	mjProcessor.Register(&ddproto.CommonBcKickout{})                //64 强制退出
	mjProcessor.Register(&mjproto.Game_ActChi{})                    //65吃牌
	mjProcessor.Register(&mjproto.Game_AckActChi{})                 //66吃牌回复
	mjProcessor.Register(&mjproto.Game_ChangShaAckActGang{})        //67长沙杠回复
	mjProcessor.Register(&mjproto.Game_ActChangShaQiShouHu{})       //68长沙起手胡
	mjProcessor.Register(&mjproto.Game_AckActChangShaQiShouHu{})    //69回复长沙起手胡
	mjProcessor.Register(&mjproto.Game_ChangshQiShouHuOverTurn{})   //70回复长沙起手胡
	mjProcessor.Register(&mjproto.Game_ChangShaOverTurnAfterGang{}) //71长沙麻将杠牌
	mjProcessor.Register(&mjproto.Game_AckActHuChangSha{})          //72长沙胡牌
	mjProcessor.Register(&mjproto.Game_DealHaiDiCards{})            //73 发送海底牌
	mjProcessor.Register(&mjproto.Game_ReqDealHaiDiCards{})         //74 玩家是否需要海底牌
	mjProcessor.Register(&mjproto.Game_AckDealHaiDiCards{})         //75 广播玩家是否需要海底牌
	mjProcessor.Register(&ddproto.CommonBcUserBreak{})              //76 广播短线
	mjProcessor.Register(&ddproto.CommonReqReconnect{})             //77 重新建立链接
	mjProcessor.Register(&mjproto.GameReqBuxiazi{})                 //78 请求补虾
	mjProcessor.Register(&mjproto.GameAckBuxiazi{})                 //79 回复补虾
	mjProcessor.Register(&ddproto.CommonReqOffline{})               //80 亲求离线
	mjProcessor.Register(&ddproto.CommonAckReconnect{})             //81 回复
	mjProcessor.Register(&mjproto.GameBcBaoting{})                  //82 广播某玩家已经报听
	mjProcessor.Register(&mjproto.GameReqBaoting{})                 //83 请求报听
	mjProcessor.Register(&mjproto.GameAckBaoting{})                 //84 回复报听的听张
	mjProcessor.Register(&mjproto.GameBcFenzhang{})                 //85 广播分张

	mjProcessor.Register(&mjproto.GameBcPiao{})  //86 开始飘的广播
	mjProcessor.Register(&mjproto.GameReqPiao{}) //87 请求飘
	mjProcessor.Register(&mjproto.GameAckPiao{}) //88 回复飘

	mjProcessor.Register(&mjproto.GameReqFly{}) //89 请求飞
	mjProcessor.Register(&mjproto.GameAckFly{}) //90 回复飞
	mjProcessor.Register(&mjproto.GameReqTi{})  //91 请求提
	mjProcessor.Register(&mjproto.GameAckTi{})  //92 回复提

	mjProcessor.Register(&ddproto.CommonBcLeaveTimeout{}) //93 广播玩家离线状态

	mjProcessor.Register(&mjproto.GameAckJiaoinfos{}) //94 向玩家推送下叫提示

	mjProcessor.Register(&mjproto.GameReqGang{}) //95 请求杠 new
	mjProcessor.Register(&mjproto.GameAckGang{}) //96 回复杠 new

	mjProcessor.Register(&ddproto.CommonReqKickout{}) //97 房主请求踢人

	mjProcessor.Register(&mjproto.GameAckTinginfos{}) //98 可胡牌列表(听牌)的推送

	mjProcessor.Register(&mjproto.GameAckBuhua{}) //99 补花的推送

	mjProcessor.Register(&mjproto.GameBcShangga{}) //100 上嘎 给每个玩家单独推送 让其上嘎
	mjProcessor.Register(&mjproto.GameReqShangga{}) //101 上嘎请求
	mjProcessor.Register(&mjproto.GameAckShangga{}) //102 上嘎ack
}