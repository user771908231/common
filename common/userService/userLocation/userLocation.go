package userLocation

import (
	"casino_common/common/service/agentService"
	"casino_common/utils/iputils"
)

//计算两玩家之间距离
func GetTwoUsersDistance(userA, userB uint32) float64 {
	//请求他人的userData 需要计算距离
	distance := iputils.GetDistanceByGi(userA, userB)
	//如果通过经纬度没有找到对应的地址，再通过ip来找
	if distance == 0 {
		distance = iputils.GetDistance(agentService.GetUserIp(userA), agentService.GetUserIp(userB))
	}
	return distance
}
