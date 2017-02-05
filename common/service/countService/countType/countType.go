package countType

//统计类型
type CountType string

const (
	ALL_GAME_COUNT CountType = "all_game_count"   //所有游戏局数

	MJ_GAME_COUNT CountType = "mj_game_count"    //麻将游戏局数
	MJ_WIN_COUNT CountType = "mj_win_count"      //麻将赢的局数
	MJ_LOSE_COUNT CountType = "mj_lose_count"    //麻将输的局数

	DDZ_GAME_COUNT CountType = "ddz_game_count"   //斗地主游戏局数
	DDZ_WIN_COUNT CountType = "ddz_win_count"      //斗地主赢的局数
	DDZ_LOSE_COUNT CountType = "ddz_lose_count"    //斗地主输的局数

	ZJH_GAME_COUNT CountType = "zjh_game_count"    //扎金花游戏局数
	ZJH_WIN_COUNT CountType = "zjh_win_count"      //扎金花赢的游戏局数
	ZJH_LOSE_COUNT CountType = "zjh_lose_count"     //扎金花输的游戏局数
)
