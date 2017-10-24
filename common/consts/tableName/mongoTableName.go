package tableName

const (
	//数据库相关
	DB_ENSURECOUNTER_KEY string = "id" //自增键

	//大厅
	DBT_T_USER                      string = "t_user"                       //用户表
	DBT_T_USER_DIAMOND_DETAILS      string = "t_user_diamond_details"       //钻石交易记录
	DBT_T_TH_NOTICE                 string = "t_th_notice"                  //公告
	DBT_T_RECHARGE_DETAILS          string = "t_recharge_details"           //充值记录
	DBT_T_ACTIVE_LIST               string = "t_config_active_list"         //活动列表
	DBT_T_TASK_INFO                 string = "t_config_task_info"           //公共任务信息表
	DBT_T_USER_TASK                 string = "t_user_task"                  //用户任务状态表
	DBT_T_USER_BAG                  string = "t_user_bag"                   //用户背包道具表
	DBT_T_MAIL_CONTENT              string = "t_mail_content"               //公共邮件表
	DBT_T_USER_MAIL_list            string = "t_user_mail_list"             //用户邮件表
	DBT_T_USER_FRIENDS_LIST         string = "t_user_friends_list"          //用户好友关系表
	DBT_T_PAYBASEDETAILS            string = "t_PayBaseDetails"             //重新订单信息
	DBT_T_USER_STRONGBOX            string = "t_user_strongbox"             //保险箱
	DBT_T_GOODS_INFO                string = "t_goods_info"                 //商品信息表
	DBT_T_CONFIG_SYS                string = "t_config_sys"                 //系统的配置表
	DBT_T_USER_ATTACH               string = "t_user_attach"                //用户签到/抽奖/补助 领取记录表
	DBT_T_DRAW_LOTTERY              string = "t_config_draw_lottery"        //转盘物品信息表
	DBT_T_FRIEND_DRAW_LOTTERY       string = "t_config_friend_draw_lottery" //朋友桌分享转盘抽奖
	DBT_T_SIGN_REWARD               string = "t_config_sign_reward"         //签到奖励表
	DBT_T_GAME_LOG                  string = "t_game_log"                   //游戏记录表
	DBT_T_GAME_COUNT                string = "t_game_count"                 //游戏记录表
	DBT_T_GAME_DAY_COUNT            string = "t_game_day_count"             //游戏记录表
	DBT_T_BONUS_TASK_COMPUTE_CONFIG string = "t_bonus_task_compute_config"  //红包任务计算配置表
	DBT_T_TRADE_LOG                 string = "t_game_trade_log"             //游戏货币流水记录表
	DBT_T_GRAY_RELEASE_USER         string = "t_gray_release_user"          //灰度发布的数据库表
	DBT_T_HALL_SERVER               string = "t_hall_server"                //客服
	DBT_T_HALL_BAGPWD               string = "t_user_bag_password"          //背包密码
	DBT_T_FUDAI_LIST                string = "t_fudai_list"                 //福袋列表
	DBT_T_FUDAI_LIST_USER           string = "t_fudai_list_user"            //福袋user
	DBT_T_FUDAI_XIANGXI_LIST        string = "t_fudai_xiangxi_list"         //福袋详细信息

	DBT_GROUP_INFO string = "t_group_info" //群组信息及成员列表

	DBT_STATISTIC_BTN_CLICK   string = "T_statistic_btn_click"   //统计相关的数据库表名字
	DBT_ADMIN_EXCHANGE_RECORD string = "t_admin_exchange_record" //红包、实物兑换记录表
	DBT_IP_ADDRESS            string = "dbt_ip_address"          //存储Ip对应的地址

	DBT_ROBOT_WECHAT_GROUP_INFO string = "t_robot_wechat_group_info" //机器人微信群表

	DBT_APPLY_AGENTPRO_RECORD string = "t_apply_agentpro_record" //代理申请记录

	DBT_GAME_SERVER_INFO = "t_game_server_info"
	DBT_GAME_ROOM_INFO   = "t_game_room_info"

	//游戏session表
	DBT_GAME_SESSION string = "t_game_session"
	//代开当前房间列表
	DBT_AGENT_CREATED_ROOM string = "t_agent_created_room"
	//代开房间记录
	DBT_AGENT_CREATE_ROOM_RECORD string = "t_agent_create_room_record"

	//统计
	DBT_STATISTICS_ROOMCARD             string = "t_statistics_roomcard"             //个人房卡游戏内实际消耗
	DBT_STATISTICS_ROOMCARD_DAY_DETAILS string = "t_statistics_roomcard_day_details" //游戏房卡消耗汇总

	DBT_ROOMCARD_LOG string = "t_roomcard_log" //房卡消耗记录表

	//代理充值系统
	DBT_AGENT_GOODS        = "t_agent_goods"        //商品信息表
	DBT_AGENT_RECHARGE_LOG = "t_agent_recharge_log" //代理商充值记录表
	DBT_AGENT_SALES_LOG    = "t_agent_sales_log"    //代理商销售记录表
	DBT_AGENT_APPLY_LOG    = "t_agent_apply_log"    //代理商申请记录表
	DBT_AGENT_REBATE_LOG   = "t_agent_rebate_log"   //代理商领取返利记录表
	DBT_AGENT_INFO         = "t_agent_info"         //代理关系表

	//数据分析
	ADMIN_USER_ATHOME          = "t_user_athome"
	ADMIN_USER_ONLINEHOUR      = "t_online_count"
	ADMIN_USER_ONLINEDAY       = "t_data_onlineday"
	ADMIN_USER_RECHARGE_RECORD = "t_user_recharge_record" //充值记录表

	//用户管理
	DBT_USER_GIVE_RECORD string = "t_user_give_record" //用户赠送记录表

	//接入服务器
	DBT_GAME_CONFIG_LOGIN               = "t_game_config_login"               //登陆服配置	//游戏配置表
	DBT_GAME_CONFIG_LOGIN_LIST          = "t_game_config_login_list"          //登陆服配置	//登陆服游戏列表（阿里云北京）
	DBT_GAME_CONFIG_LOGIN_LIST_SHENZHEN = "t_game_config_login_list_shenzhen" //登录服配置（阿里云深圳）

	//麻将
	DBT_MJ_DESK             = "t_mj_desk"           //桌子的信息
	DBT_MJ_DESK_ROUND       = "t_mj_desk_round"     //一把麻将结束
	DBT_MJ_DESK_ROUND_ALL   = "t_mj_desk_round_all" //全局麻将结束
	DBT_T_TH_GAMENUMBER_SEQ = "t_th_gamenumber_seq" //麻将 编号

	DBT_MJ_BS_DESK_ROUND     = "t_mj_bs_desk_round"     //一把白山麻将结束
	DBT_MJ_BS_DESK_ROUND_ALL = "t_mj_bs_desk_round_all" //全局白山麻将结束

	DBT_MJ_SJH_DESK_ROUND     = "t_mj_sjh_desk_round"     //一把松江河麻将结束
	DBT_MJ_SJH_DESK_ROUND_ALL = "t_mj_sjh_desk_round_all" //全局松江河麻将结束

	DBT_MJ_ZXZ_DESK_ROUND     = "t_mj_zxz_desk_round"     //一把捉虾子麻将结束
	DBT_MJ_ZXZ_DESK_ROUND_ALL = "t_mj_zxz_desk_round_all" //全局捉虾子麻将结束

	DBT_MJ_ZHZH_DESK_ROUND     = "t_mj_zhzh_desk_round"     //一把转转麻将结束
	DBT_MJ_ZHZH_DESK_ROUND_ALL = "t_mj_zhzh_desk_round_all" //全局转转麻将结束

	DBT_MJ_HZH_DESK_ROUND     = "t_mj_hzh_desk_round"     //一把红中麻将结束
	DBT_MJ_HZH_DESK_ROUND_ALL = "t_mj_hzh_desk_round_all" //全局红中麻将结束

	//斗地主
	DBT_DDZ_DESK_ROUND     = "t_ddz_desk_round"     //一把斗地主结束
	DBT_DDZ_DESK_ROUND_ALL = "t_ddz_desk_round_all" //全局斗地主结束

	//跑得快
	DBT_PDK_DESK_ROUND          = "t_pdk_desk_round"          //一把跑得快结束战绩
	DBT_PDK_DESK_ROUND_ALL      = "t_pdk_desk_round_all"      //全局跑得快结束战绩
	DBT_PDK_DESK_ROUND_PLAYBACK = "t_pdk_desk_round_playback" //一把跑得快回放

	//跑胡子
	DBT_PHZ_DESK_ROUND          = "t_phz_desk_round"          //跑胡子一局游戏战绩
	DBT_PHZ_DESK_ROUND_ALL      = "t_phz_desk_round_all"      //跑胡子全局游戏结束战绩
	DBT_PHZ_DESK_ROUND_PLAYBACK = "t_phz_desk_round_playback" //一把跑胡子回放

	//牛牛
	DBT_NIU_DESK_ROUND_ONE = "t_niuniu_desk_round_one" //牛牛1局结束战绩
	DBT_NIU_DESK_ROUND_ALL = "t_niuniu_desk_round_all" //牛牛10局结束战绩

	//拼二章
	DBT_PEZ_DESK_ROUND     = "t_pinerzhang_desk_round"     //拼二张一局游戏战绩
	DBT_PEZ_DESK_ROUND_ALL = "t_pinerzhang_desk_round_all" //拼二张全局游戏结束的战绩，包含每一局的战绩

	//百人牛牛
	DBT_BAINIU_AREA_BILL_COUNT = "t_bainiu_area_bill_count" //百人牛牛区域胜负统计

	DBT_USER_GAME_BILL = "t_user_game_bill" //玩家游戏账单记录

	DBT_T_ROBOT_INFO string = "t_robot_info" //机器人信息表

)
