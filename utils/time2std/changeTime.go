package time2std

import (
	"time"
	"fmt"
)

func Time2std(res int64) string{
	fmt.Println("res:",res)
	// 假设原始的Unix时间戳为 1652090400000
    timestamp := int64(res) / 1000  // 将毫秒级别的时间戳转换为秒级别
	fmt.Println(timestamp)
	// 使用 time.Unix 函数将 Unix 时间戳转换为 time.Time 类型
	tm:=time.Unix(timestamp,0)
	
	// 使用 time.Time.Format 函数将时间格式化为指定的字符串格式
	// fmt.Println(tm.Format("2006-01-02 15:04:05"))
    ret:=tm.Format("2006-01-02 15:04:05")
	fmt.Println(ret)

	return ret
}