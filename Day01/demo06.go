package main

import "fmt"

func main() {

	var num int = 1000
	var name string = "刘华硕123"
	fmt.Println(num)
	// Print format 格式化输出
	fmt.Printf("num = %d\n 地址 = %p\n", num, &num)

	fmt.Printf("name is %s \n地址= %p", name, &name)
}
