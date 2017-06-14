package sessionService

//是否处于维护中,第二个返回值为维护中的提示消息
func IsOnMaintain(gid int32) (bool, string) {

	return false, ""
}
