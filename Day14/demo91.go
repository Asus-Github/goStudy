package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int) // 可读可写
	go writeOnly(ch1)
	go readOnly(ch1)

	time.Sleep(time.Second * 3)
}

// 作为函数的参数或者返回值之类的。
// 指定函数去写， 不让他读取，防止通道滥用
func writeOnly(ch chan<- int) {
	// 函数的内部，处理一些写数据的操作
	ch <- 100
}

// 指定函数去读， 不让他写，防止通道滥用
func readOnly(ch <-chan int) int {
	// 取出通道的值，做一些操作，不可写的。
	data := <-ch
	fmt.Println(data)
	return data
}
