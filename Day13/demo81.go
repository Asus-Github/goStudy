package main

import "fmt"

func main() {
	//goroutine 调用：并发执行，和普通方法交替执行，一旦开始执行，它就自己执行自己的，继续运行下一行代码
	go hello()
	for i := 0; i < 100; i++ {
		fmt.Println("main - ", i)
	}
	//main函数结束了，所有的goroutine自动结束，瞬间销毁
}

func hello() int {
	for i := 0; i < 100; i++ {
		fmt.Println("hello - ", i)
	}

	return 1
}
