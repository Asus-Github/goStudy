package main

import (
	"fmt"
	"sync"
	"time"
)

/*
临界资源：指并发环境中多个进程、线程、协程共享的资源

发现结果和预想的不同，多线程加入之后，原先单线程的逻辑出现了问题。

出现了临界资源安全问题。
*/

var ticketTotal int = 20

var mutex sync.Mutex

func main() {

	go saleTicket("张菁菁")
	go saleTicket("刘华硕")
	go saleTicket("陈正乐")

	time.Sleep(time.Second * 5)
}
func saleTicket(name string) {

	for {
		mutex.Lock()
		if ticketTotal > 0 {
			ticketTotal--
			time.Sleep(time.Millisecond * 1) //如果不加这个，可能其他goroutine得不到服务，只被一个goroutine买完了
			fmt.Println(name, "抢到一张票， 剩余：", ticketTotal)
			mutex.Unlock()
		} else {
			fmt.Println("票已售完")
			mutex.Unlock()
			return
		}
	}

}
