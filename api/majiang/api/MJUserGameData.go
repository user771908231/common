package api

import "casino_common/api/majiang"

type MJUserGameData interface {
	GetPengPais() []*majiang.PengPai
	GetGangPais() []*majiang.GangPai
	GetChiPais() []*majiang.ChiPai
	GetMoPai() *majiang.MJPAI //得到摸的牌
	GetHandPais() []*majiang.MJPAI
}
