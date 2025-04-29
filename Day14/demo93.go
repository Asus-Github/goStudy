package main

import (
	"fmt"
	"time"
)

func main() {

	/*timer := time.NewTimer(time.Second * 3)
	fmt.Println(time.Now())
	timeChan := <-timer.C
	fmt.Println(timeChan)*/

	select {
	case <-time.After(3 * time.Second): //为什么要加<-？ 在 select 中，每个 case 必须是一个通道的收发操作
		mail()
	}

	/*
		错误理解（❌）：
		如果通道发过来的是某个特定值，我才执行 case

		正确理解（✅）：
		只要通道能读/写，就立即触发对应的 case，读出来什么再说。

		✅ 多个 case 会怎么选？
		如果多个 case 同时准备好了（多个通道同时可操作），Go 会随机选择一个来执行。

		如果没有任何一个 case 可操作，并且你没有写 default，select 会阻塞等待。
	*/

	time.AfterFunc(time.Second*3, mail)

	time.Sleep(time.Second * 5)
	/*
	 Go 的一个并发模型特性：主函数不会自动等待其他 goroutine 或回调执行完成，它只负责自己的执行流程。
	 Go不会因为你注册了异步事件就默认阻塞主函数。
	*/
}

func mail() {
	fmt.Println("seed a mail")
}
