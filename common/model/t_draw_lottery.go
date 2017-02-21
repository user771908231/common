package model

type T_Draw_Lottery struct {
	Sort int32 //排序
	Type int32 //类型
	Name string //显示名称
	Number float64 //数量
	Percent float64 //百分比概率 如 24.50% 存入24.50
	Version float32 //版本
}
