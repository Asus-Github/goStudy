package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go test(ch)

	/*for {
		time.Sleep(time.Second * 1)
		data, ok := <-ch
		if !ok {
			fmt.Println("读取完毕", ch)
			break
		} else {
			fmt.Println(data)
		}
	}*/
	//for range 可以简化开发
	for data := range ch {
		time.Sleep(time.Second)
		fmt.Println(data)
	}
}

func test(ch chan int) {

	for i := 0; i <= 10; i++ {
		ch <- i
	}
	// 关闭通道，告诉接收方，不会在往ch中放入数据
	close(ch)
}
