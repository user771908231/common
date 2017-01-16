package model

//签到表
type T_sign_reward struct {
	Id int32
	Day int32               //连续签到天数
	Type int32        //奖品类型
	Name string       //奖品名
	Amount float32       //奖品数量
}
