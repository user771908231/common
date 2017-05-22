package hall

//比赛的实体
type T_bisai_model struct {
	Id        int32  //比赛的model
	GameId    int32  //游戏Id
	Des       string //说明
	UserLimit int32  //最低开始人数
	Status    int32  //状态，0默认状态，1，报名，2，开始游戏，3，结束，4，过期，取消
}

func ListTBisaiModels() []T_bisai_model {
	var ret []T_bisai_model
	return ret
}
