package main

import "fmt"

// 全局变量
// 全局变量必须使用var关键字
var a float64 = 3.14

func main() {
	//局部变量
	var a float64 = 2.18

	//全局变量有的首先找全局变量，否则找局部
	//程序的执行原则自上而下
	fmt.Println(a)

}
