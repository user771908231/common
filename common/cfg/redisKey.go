package cfg

import "casino_common/common/consts/tableName"

//common使用的所有的rediskey都应该放在这里
const (
	RKEY_PRE_ASSETS_ALL         string = "AssetsInfo"         //热更新文件的redis-key ,这是前缀，后缀需要增加appClientID
	RKEY_CONFIG_SYS             string = "config_sys"         //系统配置的redis-key
	RKEY_PRE_USER               string = tableName.DBT_T_USER //用户的key
	RKEY_PRE_USER_AGENT_SESSION string = "agent_session"      //用户session的key

)
