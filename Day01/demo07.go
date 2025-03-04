package main

import "fmt"

func test() (int, int) {
	return 100, 200
}

func main() {

	//a, b := test()
	//fmt.Println(a, b)

	// 只想要一个结果 可以用 匿名变量 _ 相当于一个黑洞，不想要的东西就丢进去
	a, _ := test()
	fmt.Println(a)
	_, b := test()
	fmt.Println(b)
}
