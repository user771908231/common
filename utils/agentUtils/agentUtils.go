package agentUtils

import (
	"github.com/name5566/leaf/gate"
	"strings"
	"net"
)

//通过链接得到ip
func GetIP(a gate.Agent) string {
	ipDefault := "127.0.0.1"
	if a == nil {
		return ipDefault
	}

	ipi := a.RemoteAddr()
	if ipi == nil {
		return ipDefault
	}
	ip := ipi.(net.Addr)
	if ip == nil {
		return ipDefault
	}

	iparrs := strings.Split(ip.String(), ":")
	if iparrs != nil {
		return iparrs[0] //返回ip地址，不需要端口
	} else {
		return ipDefault
	}
}
