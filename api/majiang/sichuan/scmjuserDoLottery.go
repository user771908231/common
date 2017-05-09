package sichuan

import (
	"casino_common/common/userService"
)

func (u *SCMJUser) AfterLottery() error {
	userService.UpdateUser2MgoById(u.GetUserId()) //afterlottery 之后更新玩家的信息
	u.S = MJUSER_STATUS_SEATED                    //设置状态为坐下
	u.Banker = false                              //设置为不是庄
	u.Ready = false                               //设置为没有准备
	u.SCGameData.HandPais = nil                   //所有的手牌
	u.SCGameData.MoPai = nil                      //摸的牌
	u.SCGameData.PengPais = nil                   //碰的牌
	u.SCGameData.GangPais = nil                   //杠的牌
	u.SCGameData.HuInfo = nil                     //胡牌信息
	u.SCGameData.OutPais = nil                    //打出去的牌
	return nil
}
