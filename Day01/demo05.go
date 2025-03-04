package main

import "fmt"

func main() {

	/*
		变量交换
	*/
	var a int = 100
	var b int = 200
	// 语法糖 ，底层依旧是运用了临时变量
	a, b = b, a

	fmt.Println(a, b)

}
