package statisticsService

//局数统计器

type PlayCounter interface {
	//增加游戏局数
	IncPlayCount(win bool) (totalPlayCount int32, winPlayCount int32, e error)
}
