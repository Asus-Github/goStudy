package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	ch := make(chan int, 5)

	ch <- 1

	fmt.Println(cap(ch), len(ch))

	fmt.Println("--------------------------")
	chStr := make(chan string, 5)

	go test1(chStr)
	time.Sleep(time.Second * 5)
	for data := range chStr {
		fmt.Println(data)
		fmt.Println(cap(chStr), len(chStr))
	}

}

func test1(chStr chan string) {
	for i := 0; i < 10; i++ {
		chStr <- strconv.Itoa(i)
		fmt.Println("ttt")
	}
	close(chStr)
}

/*
缓冲通道，可以定义缓冲区的数量

如果缓冲区没有满，可以继续存放，如果满了，也会阻塞等待

如果缓冲区空的，读取也会等待，如果缓冲区中有多个数据，依次按照先进先出的规则进行读取。

如果缓冲区满了，同时有两个线程在读或者写，这个时候和普通的chan一样。一进一出。 类似队列

*/
