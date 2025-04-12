package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	// 时间格式化 2023-02-23 20:43:49
	// 格式化模板： yyyy-MM-dd HH:mm:ss
	// Go语言诞生的时间作为格式化模板：2006年1月2号下午3点4分
	// Go语言格式化时间的代码：2006-01-02 15:04:05  （记忆方式：2006 12 345） 固定的时间，别的不行
	// 固定的："2006-01-02 15:04:05"
	fmt.Println(now.Format("2006-01-02 15:04:05")) // 24小时制

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	time1()
	time4()
}

func time1() {
	// 将字符串格式化为Time对象 (获取到网页传递的时间字符串，需要转化为Time才能在代码中使用)
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loc)
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(timeStr)
	// layout 格式 时间字符串  时区位置 , 需要和前端传递的格式进行匹配
	// func ParseInLocation(layout, value string, loc *Location)
	timeTemp, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeTemp)
}

// 时间戳：更多时候和随机数结合
func time4() {
	fmt.Println("================")
	// 格林威治时间自1970年1月1日(00:00:00 GMT)至当前时间的总秒数
	// 时间戳 Unix ： 1970.1.1 00:00:00 - 当下的一个毫秒数，Unix 时间戳，不会重复的。
	now := time.Now()
	timestamp1 := now.Unix()     // 时间戳
	timestamp2 := now.UnixNano() // 纳秒的时间数

	fmt.Println(timestamp1)
	fmt.Println(timestamp2)
	//
	// 通过 Unix 转换time对象
	/*
			tsNano := int64(1712908800123456789) // 纳秒级时间戳
			sec := tsNano / 1e9
			nsec := tsNano % 1e9

			timeObj := time.Unix(sec, nsec) //第二个参数是纳秒偏移，精确到纳秒
			//如果你直接把纳秒时间戳放到 time.Unix(tsNano, 0)，会得出一个错误的时间，因为你把“纳秒”当“秒”用了。

		//timeObj := time.Unix(1712908800, 123456789) // 2024-04-12 00:00:00.123456789
		//你想表示的时间是 “在某一秒内，偏移了 123456789 纳秒”，这在性能分析、日志打点中非常有用。
	*/
	timeObj := time.Unix(timestamp1, 0) // 返回time对象
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	// 2023-2-23 20:40:31
	// Printf ： 整数补位--02如果不足两位，左侧用0补齐输出
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
