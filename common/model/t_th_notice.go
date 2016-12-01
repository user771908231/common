package model

//表示公告
type T_th_notice struct {
	Id            int32
	NoticeType    int32    //公告类型
	NoticeTitle   string   //公告标题
	NoticeContent string   //公告内容
	NoticeMemo    string   //公告备注
	NoticeFileds  []string //公告的字段
}