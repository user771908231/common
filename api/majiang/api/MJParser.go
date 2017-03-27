package api

//跑得快解析器
type MJParser interface {
	//出牌的时候比较大小
	CanOut(outpai interface{}, check interface{}) (bool, error)
	//通过一副牌的id解析牌型
	Parse(pids []int32) (interface{}, error)
	//洗一副扑克牌出来
	XiPai() interface{} //洗牌
}
