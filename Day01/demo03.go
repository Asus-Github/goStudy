package main

import "fmt"

func main() {

	var name string
	// 不初始化赋给默认值
	var age int

	var (
		name2 string
		age2  int
	)

	name = "刘华硕123"
	// ctrl + d 自动在下一行生成当前行
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(age)
	fmt.Println(age2)
}
