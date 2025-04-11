package main

import "fmt"

// init：初始化，在main方法执行之前自动执行
func init() {
	fmt.Println("hello world")
}
func main() {
	// init 函数不需要传入参数，也没有返回值，任何地方不能调用 init()
	//init()
}
