package protocol

//统一的回复
type CommonAckRpc struct {
	Msg  string
	Code int32
}
