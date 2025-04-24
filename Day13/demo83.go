package main

import (
	"fmt"
	"time"
)

/*
临界资源：指并发环境中多个进程、线程、协程共享的资源

发现结果和预想的不同，多线程加入之后，原先单线程的逻辑出现了问题。

出现了临界资源安全问题。
*/

var ticketsTotal int = 10

func main() {

	go saleTickets("张菁菁")
	go saleTickets("刘华硕")
	go saleTickets("陈正乐")

	time.Sleep(time.Second * 5)
}
func saleTickets(name string) {

	for {
		if ticketsTotal > 0 {
			ticketsTotal--
			fmt.Println(name, "抢到一张票， 剩余：", ticketsTotal)
		} else {
			fmt.Println("票已售完")
			return
		}
	}

}
