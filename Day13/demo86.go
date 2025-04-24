package main

import (
	"fmt"
	"sync"
	"time"
)

//同步等待组

var wg sync.WaitGroup

func main() {
	// 公司最后关门的人   0
	// wg.Add(2) 判断还有几个线程、计数  num=2
	// wg.Done() 我告知我已经结束了  -1
	wg.Add(2)

	go test1()
	go test2()
	fmt.Println("main等待ing")
	wg.Wait()
	fmt.Println("over")
}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1: ", i)
		time.Sleep(time.Second * 3)
	}
	wg.Done()
}

func test2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("test2: ", i)
		time.Sleep(time.Second * 3)
	}
}
