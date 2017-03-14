package handlers

import (
	"casino_common/proto/ddproto"
	"casino_common/common/service/statisticsService"
	"casino_common/utils/agentUtils"
	"github.com/name5566/leaf/gate"
	"time"
)

func HandlerClickStatistic(args []interface{}) {
	m := args[0].(*ddproto.CommonReqClickStatistic)
	a := args[1].(gate.Agent)
	c := statisticsService.NewDefaultBtnClickCounter()
	c.C(m.GetHeader().GetUserId(), agentUtils.GetIP(a), int32(m.GetBtnId()), time.Now())
}
