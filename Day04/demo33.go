package main

import "fmt"

// 回调函数：函数作为了另一个函数的参数
// 高阶函数：函数接收了另一个函数作为参数
/*
	fun1()
	fun2(fun1)
*/
func main() {
	//add作为了fmt的回调函数
	fmt.Println(add(1, 2))

	fmt.Println(oper(1, 2, add))

	// 接收并调用匿名函数
	fun1 := func(x, y int) int {
		return x * y
	}
	fmt.Println(oper(1, 2, fun1))

	//直接调用匿名函数
	ans := oper(1, 2, func(i1 int, i2 int) int {
		return i1 + i2
	})
	fmt.Println(ans)

	res := operS(2, 2, add)
	fmt.Println(res)
}

// 高阶函数，其中一个参数是另外一个函数
func oper(x int, y int, fun func(int, int) int) int {
	res := fun(x, y)
	return res
}

func add(x int, y int) int {
	return x + y
}

// 高阶函数，其中一个参数是另外一个函数
// +可变参数
func operS(x int, y int, fun ...func(int, int) int) int {
	res := fun[0](x, y) //此时需要数组取参
	return res
}
