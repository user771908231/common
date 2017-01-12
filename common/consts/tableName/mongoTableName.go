package tableName

const (
	//数据库相关
	DB_ENSURECOUNTER_KEY string = "id" //自增键

	//大厅
	DBT_T_USER                 string = "t_user"                 //用户表
	DBT_T_USER_DIAMOND_DETAILS string = "t_user_diamond_details" //钻石交易记录
	DBT_T_TH_NOTICE            string = "t_th_notice"            //公告
	DBT_T_RECHARGE_DETAILS     string = "t_recharge_details"     //充值记录
	DBT_T_USER_TASK            string = "t_user_task"            //用户任务状态表
	DBT_T_USER_BAG             string = "t_user_bag"             //用户背包道具表
	DBT_T_MAIL_CONTENT         string = "t_mail_content"         //公共邮件表
	DBT_T_USER_MAIL_list       string = "t_user_mail_list"       //用户邮件表
	DBT_T_USER_FRIENDS_LIST    string = "t_user_friends_list"    //用户好友关系表
	DBT_T_PAYBASEDETAILS       string = "t_PayBaseDetails"       //重新订单信息
	DBT_T_USER_STRONGBOX       string = "t_user_strongbox"       //保险箱
	DBT_T_GOODS_INFO           string = "t_goods_info"           //商品信息表
	DBT_T_CONFIG_SYS           string = "t_config_sys"           //系统的配置表
	DBT_T_USER_ATTACH          string = "t_user_attach"          //用户签到/抽奖/补助 领取记录表
	DBT_T_DRAW_LOTTERY         string = "t_draw_lottery"         //转盘物品信息表

	//麻将
	DBT_MJ_DESK             = "t_mj_desk"           //桌子的信息
	DBT_MJ_DESK_ROUND       = "t_mj_desk_round"     //一把麻将结束
	DBT_T_TH_GAMENUMBER_SEQ = "t_th_gamenumber_seq" //麻将 编号

	//斗地主
	DBT_DDZ_DESK_ROUND = "t_ddz_desk_round" //一把斗地主结束

	//扎金花
)
