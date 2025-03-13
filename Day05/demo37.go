package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// for range循环 返回一个值 只是数组下标，返回两个值 是 下标和下标所对应的值 快捷键：arr.forr
	for index, value := range arr {
		fmt.Println(index, value)
	}
	//区分值传递和地址传递
	arr2 := arr
	fmt.Println(arr2)
	fmt.Println(arr)
	arr2[1] = 100
	fmt.Println(arr2)
	fmt.Println(arr)
}
