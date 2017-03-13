package redpacketService

//红包管理器
type RedPacketManager interface {
	TryGetRP() (amount int64, e error) //尝试得到红包
	//通过传入场次数据来判断是否能得到红包：考虑到判断是否得到红包，都在是统计数据之后，所以一般是调用第二种方法，这种方法的效率是最高的
	TryGetRP2(winPlayCount int32) (amount int64, e error)
}
