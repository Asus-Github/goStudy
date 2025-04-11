package main

import "fmt"

func main() {

	defer fmt.Println("msgMain - 1")
	defer fmt.Println("msgMain - 2")
	fmt.Println("msgMain - 3")
	testPanic(1)
	defer fmt.Println("msgMain - 4")
}

func testPanic(num int) {
	defer fmt.Println("msg - 1")
	defer fmt.Println("msg - 2")
	fmt.Println("msg - 3")
	// 如果在函数中一旦触发了 panic，会终止后面要执行的代码。
	// 如果存在defer，正常按照defer逻辑执行 (压栈，出栈)
	if num == 1 {
		panic("出现了预定的异常----panic")
	}
	fmt.Println("msg - 4")
}
