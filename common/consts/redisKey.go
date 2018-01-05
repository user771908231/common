package consts

import "casino_common/common/consts/tableName"

//common使用的所有的rediskey都应该放在这里
const (
	RKEY_PRE_ASSETS_ALL         string = "AssetsInfo"              //热更新文件的redis-key ,这是前缀，后缀需要增加appClientID
	RKEY_PRE_ASSETS_ALL_WHITE   string = "AssetsInfoWhite"         //热更新文件的redis-key,这里是白名单的key  ,这是前缀，后缀需要增加appClientID
	RKEY_CONFIG_SYS             string = "config_sys"              //系统配置的redis-key
	RKEY_PRE_USER               string = tableName.DBT_T_USER      //用户的key
	RKEY_PRE_USER_AGENT_SESSION string = "agent_session"           //用户session的key
	RKEY_USER_COIN              string = "user_coin_redis_key"     //金币
	RKEY_USER_DIAMOND           string = "user_diamond_redis_key"  //钻石，金币场
	RKEY_USER_DIAMOND2          string = "user_diamond2_redis_key" //钻石朋友桌
	RKEY_USER_ROOMCARD          string = "user_roomcard_redis_key" //钻石朋友桌
	RKEY_MJSESSION_KEY_PRE      string = "redis_game_session"      //session redis key
	RKEY_QR_CODE                string = "redis_qr_code_key"       //二维码登陆成功后存储的key
	RKEY_SEQ_ID_KEY             string = "redis_seq_id_key"        //自增id key

	RKEY_AGENT_CREATE_DESK_LIST       string = "rkey_agent_create_desk_list"       //代开房间列表 key
	RKEY_AGENT_GROUP_CREATE_DESK_LIST string = "rkey_agent_group_create_desk_list" //代开房间列表 key

	RKEY_USER_GAME_BILL string = "rkey_user_game_bill" //玩家游戏账单数据
)

const (
	//大厅相关
	RKEY_CONF_HALL_GRAY_RELEASE_SWITH string = "rkey_conf_hall_gray_release_swith" //大厅的灰度测试开关
	RKEY_PRE_IP_ADDR                  string = "rkey_pre_ip_addr"                  //ip对应的实际地址
	RKEY_ROOMPWD_GAMEID               string = "rkey_roompwd_gameid"               //房间号对应gameid
)

const (
	//admin需要使用的
	RKEY_ALL_ONLINE_USERS      string = "rkey_all_online_users"      //记录所有在线的玩家
	RKEY_PRE_ONLINE_SINGE_GAME string = "rkey_pre_online_singe_game" //记录单个在线的游戏玩家

	//admin  游戏服务器配置相关
	ADMIN_GAMEID_ISMAINTAIN    string = "rkey_gameid_ison_maintain" //游戏服务器是否在维护
	ADMIN_GAMEID_MAINTAIN_INFO string = "rkey_gameid_maintain_info" //游戏服务器是否在维护
)

const (
	RKEY_GAME_WHITE_LIST string = "rkey_game_white_list" //白名单
)

//数据恢复
const (
	RKEY_PDK_DATARECOVER_ID_LIST string = "rkey_pdk_datarecover_id_list" //跑得快数据恢复房间ID列表
	RKEY_PDK_DATARECOVER_DATA    string = "rkey_pdk_datarecover_data"    //跑得快数据恢复数据

	RKEY_ZZHZ_DATARECOVER_ID_LIST string = "rkey_zzhz_datarecover_id_list" //转转红中朋友桌麻将数据恢复房间ID列表
	RKEY_ZZHZ_DATARECOVER_DATA    string = "rkey_zzhz_datarecover_data"    //转转红中朋友桌麻将数据恢复数据

	RKEY_ZZHZ_DATARECOVER_ID_LIST_COIN string = "rkey_zzhz_datarecover_id_list_coin" //转转红中金币场麻将数据恢复房间ID列表
	RKEY_ZZHZ_DATARECOVER_DATA_COIN    string = "rkey_zzhz_datarecover_data_coin"    //转转红中金币场麻将数据恢复数据

	RKEY_SHENQI_DATARECOVER_ID_LIST string = "rkey_shenqi_datarecover_id_list" //神奇朋友桌麻将数据恢复房间ID列表
	RKEY_SHENQI_DATARECOVER_DATA    string = "rkey_shenqi_datarecover_data"    //神奇朋友桌麻将数据恢复数据

	RKEY_SHENQI_DATARECOVER_ID_LIST_COIN string = "rkey_shenqi_datarecover_id_list_coin" //神奇金币场麻将数据恢复房间ID列表
	RKEY_SHENQI_DATARECOVER_DATA_COIN    string = "rkey_shenqi_datarecover_data_coin"    //神奇金币场麻将数据恢复数据

	RKEY_ZXZ_DATARECOVER_ID_LIST string = "rkey_zxz_datarecover_id_list" //捉瞎子麻将数据恢复房间ID列表
	RKEY_ZXZ_DATARECOVER_DATA    string = "rkey_zxz_datarecover_data"    //捉瞎子麻将数据恢复数据

	RKEY_MJ_DATARECOVER_ID_LIST string = "rkey_mj_datarecover_id_list" //四川、长沙麻将数据恢复房间ID列表
	RKEY_MJ_DATARECOVER_DATA    string = "rkey_mj_datarecover_data"    //四川、长沙麻将数据恢复数据

	RKEY_SJH_DATARECOVER_ID_LIST string = "rkey_sjh_datarecover_id_list" //松江河麻将数据恢复房间ID列表
	RKEY_SJH_DATARECOVER_DATA    string = "rkey_sjh_datarecover_data"    //松江河麻将数据恢复数据

	RKEY_PHZ_DATARECOVER_ID_LIST string = "rkey_phz_datarecover_id_list" //跑胡子娄底放炮罚数据恢复_房间ID列表
	RKEY_PHZ_DATARECOVER_DATA    string = "rkey_phz_datarecover_data"    //跑胡子娄底放炮罚数据恢复_数据

	RKEY_PHZ_SHAOYANGZIPAI_DATARECOVER_ID_LIST string = "rkey_phz_shaoyangzipai_datarecover_id_list" //跑胡子邵阳字牌数据恢复_房间ID列表
	RKEY_PHZ_SHAOYANGZIPAI_DATARECOVER_DATA    string = "rkey_phz_shaoyangzipai_datarecover_data"    //跑胡子邵阳字牌数据恢复_数据

	RKEY_NIUNIU_SNAPSHOT_ID_LIST string = "rkey_niuniu_snapshot_id_list" //牛牛数据恢复id列表
	RKEY_NIUNIU_SNAPSHOT_DATA    string = "rkey_niuniu_snapshot_data"    //牛牛数据恢复数据

	RKEY_PAOYAO_SNAPSHOT_ID_LIST string = "rkey_paoyao_snapshot_id_list" //牛牛数据恢复id列表
	RKEY_PAOYAO_SNAPSHOT_DATA    string = "rkey_paoyao_snapshot_data"    //牛牛数据恢复数据
)

const (
	RKEY_SHENQI_WHITE_USER_WINCOIN_INFO string = "rkey_shenqi_white_user_wincoin_info" //神奇朋友桌麻将白名单玩家输赢数据
	RKEY_ZZHZ_WHITE_USER_WINCOIN_INFO   string = "rkey_zzhz_white_user_wincoin_info"   //转转红中朋友桌麻将白名单玩家输赢数据
	RKEY_SJH_WHITE_USER_WINCOIN_INFO    string = "rkey_sjh_white_user_wincoin_info"    //松江河朋友桌麻将白名单玩家输赢数据
)

const (
	RKEY_PDK_REXIPAI_CHANCE_BOOM string = "rkey_pdk_rexipai_boom"
	RKEY_PDK_REXIPAI_CHANCE_AIR  string = "rkey_pdk_rexipai_air"
	RKEY_PDK_REXIPAI_CHANCE_SHUN string = "rkey_pdk_rexipai_shun"
	RKEY_PDK_WHITEUSER_WINCOUNT  string = "rkey_pdk_whiteuser_wincount"
)

const (
	RKEY_DESK_PRE_XIPAI_DATA string = "rkey_desk_pre_xipai_data" //牌桌预洗牌
)

const (
	RKEY_HZH_ROBOT_GAME_BILL    string = "rkey_hzh_robot_game_bill"    //红中机器人策略控制用的赢局统计信息
	RKEY_SHENQI_ROBOT_GAME_BILL string = "rkey_shenqi_robot_game_bill" //神奇机器人策略控制用的赢局统计信息
)

const (
	RKEY_ROBOT_WIN_RATE_HZH    string = "rkey_robot_win_rate_hzh"    //红中机器人胜率
	RKEY_ROBOT_WIN_RATE_SHENQI string = "rkey_robot_win_rate_shenqi" //神奇机器人胜率
)

//配置信息
const (
	RKEY_CONFIG_ROBOT_HEADERURL_PREFIX string = "rkey_config_robot_headerurl_prefix" //机器人头像前缀
)

const (
	RKEY_ROBOT_AVAILABLE string = "rkey_robot_available" //机器人可用情况
)
