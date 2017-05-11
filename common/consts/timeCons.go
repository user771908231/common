package consts

import "time"

const (
	APPLYDISSOLVE_DURATION = time.Second * 2 * 60 //申请解散房间 等待的时间
	COIN_RAEDY_DURATION    = time.Second * 60     //金币场，一局结束之后，需要准备，超时不准备需要提出房间
)
