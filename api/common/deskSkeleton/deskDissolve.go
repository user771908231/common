package deskSkeleton

import (
	"casino_common/proto/ddproto"
	"github.com/golang/protobuf/proto"
	"casino_common/common/Error"
	"casino_common/common/log"
	"time"
	"casino_common/common/consts"
)

var MJUER_APPLYDISSOLVE_S_REFUSE int32 = -1 //拒绝解散
var MJUER_APPLYDISSOLVE_S_DEFAULT int32 = 0 //没有处理
var MJUER_APPLYDISSOLVE_S_AGREE int32 = 1   //同意解散

var ERR_REQ_REPETITION = Error.NewError(consts.ACK_RESULT_ERROR, "申请失败")
//申请解散
func (d *DeskSkeleton) ApplyDissolve(userId1 uint32) error {
	log.T("锁日志: %v ApplyDissolve(%v)的时候等待锁", d.DlogDes(), userId1)
	//如果已经是在解散的状态中，这个时候请求解散房间，那么直接同意解散
	if d.ApplyDis == true {
		go func() {
			defer Error.ErrorRecovery("离线的人自动统一")
			d.ApplyDissolveBack(userId1, true)
		}()
		return nil
	}

	d.ApplyDis = true //有人申请解散房间:设置为等待解散的状态
	//4,如果房间还没有解散，那么现在开始做定时处理
	for _, u := range d.ListBaseUserSkeletons() {
		if u != nil {
			userId := u.GetUserId()
			//短线离线的人 默认同意.其他的人增加定时器
			if u.IsBreak() || u.IsLeave() || u.GetUserId() == userId1 {
				go func() {
					defer Error.ErrorRecovery("离线的人自动同意解散房间")
					d.ApplyDissolveBack(userId, true)
				}()
			} else if u.ApplyDissolve == MJUER_APPLYDISSOLVE_S_DEFAULT {
				log.T("开始给user %v 设置同意解散房间倒计时", userId)
				u.DissolveTimer = time.AfterFunc(time.Second*5*60, func() {
					defer Error.ErrorRecovery("倒计时解散房间")
					//超时的时候，需要默认设置为同意
					log.T("玩家[%v]同意或者拒绝解散房间的倒计时超时...现在开始进行超时处理.", userId)
					d.ApplyDissolveBack(userId, true)
				})
			}
		}
	}

	//2,申请解散 回复
	ack := new(ddproto.CommonBcApplyDissolve)
	ack.UserId = proto.Uint32(userId1)
	d.BroadCastProto(ack)
	return nil
}

//申请解散房间
func (d *DeskSkeleton) ApplyDissolveBack(args ...interface{}) error {
	userId := args[0].(uint32)
	agree := args[1].(bool)

	log.T("ApplyDissolveBack(%v,%v),,desk.GetApplyDis():%v", userId, agree, d.ApplyDis)
	if !d.ApplyDis {
		//不在 判断拒绝或者同意的状态
		//return ERR_REQ_REPETITION
		log.W("玩家重复请求解散房间")
		return nil
	}

	//如果有人拒绝则表示申请解散失败...
	if !agree {
		log.T("%v 由于玩家[%v]agree:%v 所以解散失败..", d.DlogDes(), userId, agree)
		//解散失败，需要把其他人申请的状态设置为 没有统一
		for _, u := range d.ListBaseUserSkeletons() {
			if u != nil {
				u.ApplyDissolve = MJUER_APPLYDISSOLVE_S_DEFAULT //有人拒绝以后设置为默认状态
				if u.DissolveTimer != nil {
					u.DissolveTimer.Stop() //停止定时器
				}
			}
		}

		d.ApplyDis = false //有人不同意解散房间：设置不在申请的阶段...
		ack := new(ddproto.CommonAckApplyDissolveBack)
		ack.UserId = proto.Uint32(userId)
		d.BroadCastProto(ack)
		return nil
	}

	//得到用户，并判断是否是重复请求，如果不是重复请求就设置状态，并返回ack
	user := d.GetBaseUserSkeleton(userId)
	if user.ApplyDissolve != MJUER_APPLYDISSOLVE_S_DEFAULT {
		return ERR_REQ_REPETITION
	}
	//设置等待状态
	user.ApplyDissolve = MJUER_APPLYDISSOLVE_S_AGREE //设置申请解散状态
	if user.DissolveTimer != nil {
		user.DissolveTimer.Stop() //停止定时器
	}

	//回复解散的信息
	ack := new(ddproto.CommonAckApplyDissolveBack) //回复ack
	ack.UserId = proto.Uint32(userId)
	ack.Agree = proto.Bool(agree)
	d.BroadCastProto(ack)

	//如果可以解散 直接解散房间，如果所有的人都同意，那么就解散

	if d.allAgreeDissolve() {
		d.ApplyDis = false                 //可以解散房间，设置为不在申请的状态
		err := d.GetRoom().DissolveDesk(d) //统一解散房间 开始解散房间
		if err != nil {
			log.E("解散房间的时候发生错误:%v", err)
		}
		return err
	}
	return nil
}

//所有人都同意解散
func (d *DeskSkeleton) allAgreeDissolve() bool {
	for _, u := range d.ListBaseUserSkeletons() {
		if u == nil {
			continue
		}
		if u.ApplyDissolve == MJUER_APPLYDISSOLVE_S_DEFAULT {
			return false
		}

		if u.ApplyDissolve == MJUER_APPLYDISSOLVE_S_REFUSE {
			return false
		}
	}
	return true
}
