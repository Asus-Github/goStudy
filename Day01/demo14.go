package main

import "fmt"

func main() {

	//浮点数： 符号位 + 指数位 + 尾数位（可能会丢失，造成精度丢失）
	var num1 float32 = -123.0000901 //自动截断
	var num2 float64 = -123.0000901
	//go语言中 浮点数默认使用float64
	fmt.Println(num1, "\n", num2)

}
