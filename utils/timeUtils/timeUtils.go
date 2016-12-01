package timeUtils

import (
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"
const TIME_LAYOUT_YYYY_MM_DD = "2006-01-02"

func Format(t time.Time) string {
	return t.Format(TIME_LAYOUT)
}

func String2YYYYMMDDHHMMSS(s string) time.Time {
	result, _ := time.Parse(TIME_LAYOUT, s)
	return result
}

func FormatYYYYMMDD(t time.Time) string {
	return t.Format(TIME_LAYOUT_YYYY_MM_DD)
}

/**
	字符转转时间
 */
func StringYYYYMMDD2time(s string) time.Time {
	result, _ := time.Parse(TIME_LAYOUT_YYYY_MM_DD, s)
	return result
}

func NowYYYYMMDD() time.Time {
	t := time.Now()
	s := t.Format(TIME_LAYOUT_YYYY_MM_DD)
	return StringYYYYMMDD2time(s)
}


/**
	判断两个时间多少天
 */
func DiffDays(t1, t2 time.Time) int32 {
	dur := t2.Sub(t1)
	hs := dur.Hours()
	var result int32 = int32(hs)
	return result
}