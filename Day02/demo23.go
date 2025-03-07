package main

import "fmt"

func main() {
	var sum int = 2 //go语言中只有int ，其中int多少位32 or 64根据cpu架构决定，也可以指定eg:int32 int64
	for i := 1; i < 3; i++ {
		fmt.Println(sum)
		sum *= i
	}
	fmt.Println(sum)

	for {
		//死循环
	}
	fmt.Println("test for end！")
}
