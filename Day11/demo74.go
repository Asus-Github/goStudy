package main

import (
	"fmt"
	"strconv"
)

func main() {

	s1 := "true"
	// string convert  = strconv    convert:转变
	// 转化 - 字符串转bool（parse：解析）
	s2, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T %t\n", s2, s2) //%t —— 格式化布尔值 只能用于 bool 类型的变量，否则会报错。
	}
	// bool转字符串（格式化 format）
	s3 := strconv.FormatBool(s2)
	fmt.Printf("%T %s\n", s3, s3)

	s4 := "10000"
	num, err := strconv.ParseInt(s4, 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T %d\n", num, num)
	}
	//base = 传入什么进制（会自动转换）
	s5 := strconv.FormatInt(num, 2)
	fmt.Printf("%T %s\n", s5, s5)

	// 10进制转换字符串，简便方法  atoi  itoa

	atoi, err := strconv.Atoi("-20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T %d\n", atoi, atoi)
	}
	//十进制数字转字符串
	/*
		func Itoa(i int) string {
			return FormatInt(int64(i), 10)
		}
	*/
	itoa := strconv.Itoa(30)
	fmt.Printf("%T %s\n", itoa, itoa)
}
