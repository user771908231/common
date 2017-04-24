package cfg

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
)
