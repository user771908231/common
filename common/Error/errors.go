package Error

//创建房间失败

var (
	ERROR_CREATEDESK error = NewError(-1, "创建房间失败")
	ERROR_ENTERDESK  error = NewError(-1, "房间号错误")
)
