package countType

//统计类型
type CountType string

const (
	ALL_GAME_COUNT CountType = "all_game_count"   //所有游戏局数
	ALL_GAME_WIN CountType = "all_game_win"      //所有游戏赢的局数

	//房费统计
	ALL_COIN_FEE_COUNT CountType = "all_coin_fee_count"  //所有房费统计
	DDZ_COIN_FEE_COUNT CountType = "ddz_coin_fee_count"  //斗地主房费统计
	MJ_COIN_FEE_COUNT CountType = "mj_coin_fee_count"  //麻将房费统计
	ZJH_COIN_FEE_COUNT CountType = "zjh_coin_fee_count"  //炸金花房费统计

	//麻将
	MJ_GAME_COUNT CountType = "mj_game_count"    //麻将游戏局数
	MJ_WIN_COUNT CountType = "mj_win_count"      //麻将赢的局数
	MJ_LOSE_COUNT CountType = "mj_lose_count"    //麻将输的局数

	MJ_WIN_FRIEND_COUNT CountType = "mj_win_friend_count" //斗地主好友桌赢的局数
	MJ_LOSE_FRIEND_COUNT CountType = "mj_lose_friend_count" //斗地主好友桌输的局数

	MJ_WIN_LV1_COUNT CountType = "mj_win_lv1_count"  //麻将房间1赢
	MJ_WIN_LV2_COUNT CountType = "mj_win_lv2_count"  //麻将房间2赢
	MJ_WIN_LV3_COUNT CountType = "mj_win_lv3_count"  //麻将房间3赢
	MJ_WIN_LV4_COUNT CountType = "mj_win_lv4_count"  //麻将房间4赢

	MJ_LOSE_LV1_COUNT CountType = "mj_lose_lv1_count"  //麻将房间1输
	MJ_LOSE_LV2_COUNT CountType = "mj_lose_lv2_count"  //麻将房间2输
	MJ_LOSE_LV3_COUNT CountType = "mj_lose_lv3_count"  //麻将房间3输
	MJ_LOSE_LV4_COUNT CountType = "mj_lose_lv4_count"  //麻将房间4输


	//斗地主
	DDZ_GAME_COUNT CountType = "ddz_game_count"   //斗地主游戏局数
	DDZ_WIN_COUNT CountType = "ddz_win_count"      //斗地主赢的局数
	DDZ_LOSE_COUNT CountType = "ddz_lose_count"    //斗地主输的局数

	DDZ_WIN_FRIEND_COUNT CountType = "ddz_win_friend_count" //斗地主好友桌赢的局数
	DDZ_LOSE_FRIEND_COUNT CountType = "ddz_lose_friend_count" //斗地主好友桌输的局数

	DDZ_WIN_LV1_COUNT CountType = "ddz_win_lv1_count"  //斗地主房间1赢
	DDZ_WIN_LV2_COUNT CountType = "ddz_win_lv2_count"  //斗地主房间2赢
	DDZ_WIN_LV3_COUNT CountType = "ddz_win_lv3_count"  //斗地主房间3赢
	DDZ_WIN_LV4_COUNT CountType = "ddz_win_lv4_count"  //斗地主房间4赢

	DDZ_LOSE_LV1_COUNT CountType = "ddz_lose_lv1_count"  //斗地主房间1输
	DDZ_LOSE_LV2_COUNT CountType = "ddz_lose_lv2_count"  //斗地主房间2输
	DDZ_LOSE_LV3_COUNT CountType = "ddz_lose_lv3_count"  //斗地主房间3输
	DDZ_LOSE_LV4_COUNT CountType = "ddz_lose_lv4_count"  //斗地主房间4输

	//炸金花
	ZJH_GAME_COUNT CountType = "zjh_game_count"    //扎金花游戏局数
	ZJH_WIN_COUNT CountType = "zjh_win_count"      //扎金花赢的游戏局数
	ZJH_LOSE_COUNT CountType = "zjh_lose_count"     //扎金花输的游戏局数

	ZJH_WIN_FRIEND_COUNT CountType = "zjh_win_friend_count" //炸金花好友桌赢的局数
	ZJH_LOSE_FRIEND_COUNT CountType = "zjh_lose_friend_count" //炸金花好友桌输的局数

	ZJH_WIN_LV1_COUNT CountType = "zjh_win_lv1_count"  //炸金花房间1赢
	ZJH_WIN_LV2_COUNT CountType = "zjh_win_lv2_count"  //炸金花房间2赢
	ZJH_WIN_LV3_COUNT CountType = "zjh_win_lv3_count"  //炸金花房间3赢
	ZJH_WIN_LV4_COUNT CountType = "zjh_win_lv4_count"  //炸金花房间4赢
	ZJH_WIN_LV5_COUNT CountType = "zjh_win_lv5_count"  //炸金花房间5赢
	ZJH_WIN_LV6_COUNT CountType = "zjh_win_lv6_count"  //炸金花房间6赢

	ZJH_LOSE_LV1_COUNT CountType = "zjh_lose_lv1_count"  //炸金花房间1输
	ZJH_LOSE_LV2_COUNT CountType = "zjh_lose_lv2_count"  //炸金花房间2输
	ZJH_LOSE_LV3_COUNT CountType = "zjh_lose_lv3_count"  //炸金花房间3输
	ZJH_LOSE_LV4_COUNT CountType = "zjh_lose_lv4_count"  //炸金花房间4输
	ZJH_LOSE_LV5_COUNT CountType = "zjh_lose_lv5_count"  //炸金花房间5输
	ZJH_LOSE_LV6_COUNT CountType = "zjh_lose_lv6_count"  //炸金花房间6输

)
