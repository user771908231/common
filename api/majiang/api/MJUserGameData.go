package api

import "casino_common/api/majiang"

type MJUserGameData interface {
	GetPengPais() []*majiang.PengPai
	GetGangPais() []*majiang.GangPai
	GetChiPais() []*majiang.ChiPai
	GetHandPais() []*majiang.MJPAI
}
