package main

import (
	"fmt"
)

func main() {

	//单向通道
	//ch := make(<-chan int) //只读
	//ch := make(chan<- int) //只写

	//双向通道
	//ch := make(chan int,1)

	//select
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 100
	}()

	go func() {
		//time.Sleep(1 * time.Second)
		ch2 <- 200
	}()
	// 读取chan数据，无论谁先放入，我们就用谁，抛弃其他的.
	// select 和 swtich 只是在通道中使用，case表达式需要是一个通道结果
	select {
	case data := <-ch1:
		fmt.Println(data)
	case data := <-ch2:
		fmt.Println(data)
	}

	/*
		1、每一个case必须是一个通道的操作  <-

		2、所有chan操作都有要结果（通道表达式都必须会被求值）

		3、如果任意的通道拿到了结果。它就会立即执行该case、其他就会被忽略

		4、如果有多个case都可以运行，select是随机选取一个执行，其他的就不会执行。

		5、如果存在default，执行该语句，如果不存在，阻塞等待 select 直到某个通道可以运行。

	*/
}
