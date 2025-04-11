package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "liuhuashuo,ai,zhangjingjing"
	// 1、判断某个字符是否包含了指定的内容 Contains
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains(str, "liuhuashuo"))
	// 2、判断某个字符串是否包含了 多个字符 中的某一个 有其中一个就是对的
	fmt.Println(strings.ContainsAny(str, "p"))  //false
	fmt.Println(strings.ContainsAny(str, "po")) //true

	// 3、统计这个字符在指定字符串中出现的数量 Count() 计数
	// func Count(s, substr string) int
	fmt.Println(strings.Count(str, "n")) // 3

	//寻找这个字符串第一次出现的位置 Index()
	fmt.Println(strings.Index(str, "i"))
	//寻找这个字符串最后一次出现的位置 LastIndex()
	fmt.Println(strings.LastIndex(str, "ua"))

	fileName := "2024112645.mp3"
	//判断是否用给定字符（串）开头，HasPrefix() bool
	temp := strings.HasPrefix(fileName, "2")
	fmt.Println(temp)
	/*
		if strings.HasPrefix(fileName, "2024") {
		      fmt.Println("找到2024开头的文件：", fileName)
		}
	*/
	if strings.HasSuffix(fileName, "mp3") {
		fmt.Println("找到了mp3文件")
	}
	// 8、拼接字符串, 数组或者切片拼接 ，前端给了我们多个参数。保存为一个字符串
	// Join()
	str1 := []string{"I", "am", "a", "boy"}
	fmt.Println(strings.Join(str1, " "))

	str2 := strings.Join(str1, "-")
	fmt.Println(str2)
	p := strings.Split(str2, "-")
	fmt.Println(p)
	fmt.Println(p[0])

	// 大小写ToUpper()
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	// 替换 ,-1=改所有的，  1=就是一个  2就是2两个
	fmt.Println(strings.Replace(str, "a", "我", 1))

	// 截取某个字符串
	str5 := str[0:5]
	fmt.Println(str5)
}
