package agentUtils

import (
	"github.com/name5566/leaf/gate"
	"strings"
)

//通过链接得到ip
func GetIP(a gate.Agent) string {
	if a == nil {
		return "127.0.0.1"
	}
	ip := a.RemoteAddr()
	iparrs := strings.Split(ip.(string), ":")
	if iparrs != nil && len(iparrs) == 2 {
		return iparrs[0] //返回ip地址，不需要端口
	} else {
		return "127.0.0.1"
	}
}
