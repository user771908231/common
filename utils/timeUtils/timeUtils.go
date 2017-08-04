package timeUtils

import (
	"time"
	"casino_common/common/log"
	"casino_common/utils/rand"
	"errors"
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

/**
	判断两个日期是否是同一天
 */
func EqualDate(time1 time.Time, time2 time.Time) bool {
	time1 = String2YYYYMMDDHHMMSS(Format(time1))
	time2 = String2YYYYMMDDHHMMSS(Format(time2))

	log.T("判断时间是不是同一天time1[%v],time2[%v]", time1, time2)
	log.T("y[%v],2[%v]", time1.Year(), time2.Year())
	log.T("m[%v],2[%v]", time1.Month(), time2.Month())
	log.T("r[%v],2[%v]", time1.Day(), time2.Day())

	return time1.Year() == time2.Year() &&
		time1.Month() == time2.Month() &&
		time1.Day() == time2.Day()
}

//取两者的差
func DiffSec(time1, time2 time.Time) int64 {
	a := String2YYYYMMDDHHMMSS(Format(time1))
	b := String2YYYYMMDDHHMMSS(Format(time2))
	return b.Unix() - a.Unix()
}

func RandDuration(a, b int32) time.Duration {
	d := rand.Rand(a, b)
	return time.Second * time.Duration(d)
}

//取得时间区间内的每一天的列表
func RangeDayList(start_time time.Time, end_time time.Time)([]time.Time, error) {
	if start_time.After(end_time) {
		return []time.Time{}, errors.New("end_time必须大于start_time")
	}
	sub_day_sum := int(end_time.Sub(start_time).Hours()/24 + 1)

	day_list := make([]time.Time, sub_day_sum)

	start_day := start_time
	for i,_ := range day_list {
		day_list[i] =  start_day
		start_day = start_day.AddDate(0, 0,1)
	}

	return day_list, nil
}
