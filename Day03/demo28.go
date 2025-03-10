package main

import "fmt"

func main() {

	res := getSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)
}

/*
可变参数
用数组取下标
可用len()取参数数量
*/
func getSum(num ...int) int {
	sum := 0
	l := len(num)

	for i := 0; i < l; i++ {
		sum += num[i]
	}
	return sum
}
