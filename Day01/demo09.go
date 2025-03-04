package main

import "fmt"

func main() {

	//常量
	//常量 定义可以不使用，变量定义必须使用，因为常量和变量两个存放的内存地址不同
	/*
		栈，堆，常量池
	*/
	const URL string = "www.baidu.com"
	//隐式定义 常量的定义可以省略一些基础类型
	const URL2 = "http://www.baidu.com"

	const URL3, URL4 string = "www.baidu.com", "www.baidu2.com"

	fmt.Println(URL)
}
