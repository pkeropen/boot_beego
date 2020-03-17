package utils

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	uuid "github.com/satori/go.uuid"
)

func String2Int(val string) int {

	goodsId_int, err := strconv.Atoi(val)
	if err != nil {
		return -1
	} else {
		return goodsId_int
	}
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}

func Int642String(val int64) string {
	return strconv.FormatInt(val, 10)
}

func Float642String(val float64) string {
	return strconv.FormatFloat(val, 'E', -1, 64)
}

func GetUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	} else {
		return uuid.String()
	}
}

//the result likes 1423361979
func GetTimestamp() int64 {
	return time.Now().Unix()
}

//the result likes 2015-02-08 10:19:39 AM
func FormatTimestamp(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

func ExactMapValues2Int64Array(maparray []orm.Params, key string) []int64 {

	var vals []int64
	for _, value := range maparray {
		vals = append(vals, value[key].(int64))
	}
	return vals
}

func ExactMapValues2StringArray(maparray []orm.Params, key string) []string {

	var vals []string
	for _, value := range maparray {
		vals = append(vals, value[key].(string))
	}
	return vals
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}


func Now() string{
	return time.Now().Format("2006-01-02 15:04:05")
}


