package main

import "fmt"

func main() {

	//自动推导，一个短变量声明
	name := "asus"
	fmt.Println(name)

	age := 181
	//语法糖（简化开发）
	//:= 相当于快速定义变量 如果赋值了，那么会自动推导
	fmt.Println(age)

	Name := "asus2"
	// go还是区分大小写的
	fmt.Println(Name)
	fmt.Println(name)
}
