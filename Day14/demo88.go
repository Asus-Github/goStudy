package main

import (
	"fmt"
	"time"
)

// 当一个数据被发送到通道时，在发送语句中被阻塞，直到另一个Goroutine从该通道读取数据
func main() {

	var flag chan bool
	flag = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		time.Sleep(time.Second * 3)
		flag <- true
	}()
	//通道是goroutine之间的连接，所有通道的发送和接收必须处在不同的goroutine中.
	read := <-flag
	fmt.Println(read)

}
