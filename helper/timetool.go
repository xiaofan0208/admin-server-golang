package helper

import (
	"time"
)

// 时间戳（秒）
func GetTimestampSecond() (int64){
	return time.Now().Unix()
}

// 时间戳（毫秒）
func GetMillisecond() (int64){
	return time.Now().UnixNano()/ 1e6
}

// 秒 转换成 字符串 (格式 ：  2006-01-02 15:04:05 )
func SecondFormatDetailStr(tt int64) (string){
	t := time.Unix(tt,0)
	return t.Format("2006-01-02 15:04:05")
}

// 秒 转换成 字符串 (格式 ：  2006-01-02 )
func SecondFormatStr(tt int64) (string){
	t := time.Unix(tt,0)
	return t.Format("2006-01-02")
}

// 毫秒 转换成 字符串 (格式 ：  2006-01-02 15:04:05 )
func MillisecondFormatDetailStr(tt int64) (string){
	tt = int64(tt / 1000)
	t := time.Unix(tt,0)
	return t.Format("2006-01-02 15:04:05")
}

// 毫秒 转换成 字符串 (格式 ：  2006-01-02 )
func MillisecondFormatStr(tt int64) (string){
	tt = int64(tt / 1000)
	t := time.Unix(tt,0)
	return t.Format("2006-01-02")
}