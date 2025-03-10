package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(max(1, 2))
	m := max(1, 2)
	fmt.Println(m)
}

/*

	func 函数名（参数名 类型，...）返回值类型{

	}

*/

func max(num1 int, num2 int) int {

	var res int
	if num1 > num2 {
		res = num1
	} else {
		res = num2
	}
	return res

}
