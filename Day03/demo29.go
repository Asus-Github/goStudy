package main

import "fmt"

func main() {
	fmt.Println(getSum1(3))
}

/*
递归函数太多层可能导致栈溢出
*/
func getSum1(num int) int {

	if num == 1 {
		return 1
	}
	return getSum1(num-1) + num
}
