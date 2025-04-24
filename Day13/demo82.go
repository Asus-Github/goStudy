package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	/*//获取GoRoot路径
	fmt.Println(runtime.GOROOT())
	//获取操作系统类型
	fmt.Println(runtime.GOOS)
	//获取cpu数量
	fmt.Println(runtime.NumCPU())*/

	go func() {
		fmt.Println("start")
		runtimeTest()
		fmt.Println("end")
	}()
	runtime2()
	time.Sleep(time.Second * 3) //⭐⭐main结束太快会 导致协程还没来得及执行就销毁了
}

func runtimeTest() {
	defer fmt.Println("defer1")
	//return // 只是终止了函数
	runtime.Goexit() // 终止当前的 goroutine
	fmt.Println("defer2")
}
func runtime2() {

	// goroutine是竞争cpu的  ，调度
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine", i)
		}
	}()
	//runtime.Goexit()
	//❌fatal error: no goroutines (main called runtime.Goexit) - deadlock!
	//终止当前 当前 当前 的 goroutine 如果没有goroutine则会报错
	for i := 0; i < 5; i++ {
		// gosched:礼让, 让出时间片，让其他的 goroutine 先执行
		// cpu是随机，相对来说，可以让一下，但是不一定能够成功
		// schedule
		runtime.Gosched()
		fmt.Println("main-", i)
	}
}
