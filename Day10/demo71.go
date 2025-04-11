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
	defer func() { //推入 defer 栈，尚未执行，捕捉panic
		if msg := recover(); msg != nil { // recover func recover() any 返回panic传递的值
			fmt.Println("recover执行了... panic msg:", msg)
			// 处理逻辑
			fmt.Println("---------程序已恢复----------")
		}
	}()
	//msg := recover()
	//fmt.Println(msg)
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

/*
为什么 recover() 能拿到 panic() 抛出的信息？
因为 Go 语言的 panic() 和 recover() 是一对“协作机制”，它们被设计成专门用于错误传递 + 捕获的。
*/
