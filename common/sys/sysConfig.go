package sys

import (
	"casino_common/proto/ddproto"
	"casino_common/common/model/configDao"
	"casino_common/utils/redisUtils"
	"errors"
	"fmt"
	"casino_common/common/consts"
)

var GAMEENV struct {
	PRODMODE bool
	DEVMODE  bool
}

var CONFIG_SYS *ddproto.TConfigSys

//初始化系统配置
func initSysConfig() error {
	fmt.Println("开始初始化 sysConfig")
	CONFIG_SYS = configDao.GetConfigSys()
	if CONFIG_SYS != nil {
		redisUtils.SetObj(consts.RKEY_CONFIG_SYS, CONFIG_SYS) //把系统配置存放在sysconfig中
		return nil
	} else {
		//为空表示异常
		fmt.Println("error:	没有找到sys_config 的配置...")
		return errors.New("初始化系统配置的时候出错")
	}
	return nil
}

//初始化游戏环境变量
func initGAMEENV(prodMode bool) error {
	GAMEENV.PRODMODE = prodMode
	GAMEENV.DEVMODE = !prodMode
	return nil
}
