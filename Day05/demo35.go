package main

import "fmt"

// 数组
func main() {

	//var arr int
	var arr [5]int //相对于单个变量 数组只是增加了个数[5]
	for i := 1; i < 5; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	var anyarr [5]any
	anyarr[0] = "123"
	anyarr[1] = 456
	anyarr[2] = "a"
	fmt.Println(anyarr)
	fmt.Println(len(anyarr)) //虽然我只向数组里放置了3个元素，但是数组的长度是不变的，还是初始化的5
}
