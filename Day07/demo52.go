package main

import "fmt"

// 数组指针 顾名思义 指向数组的指针
func main() {

	var arr = [5]int{1, 2, 3, 4, 5}

	var ptr *[5]int

	ptr = &arr
	fmt.Printf("arr的地址：%p\n", &arr)
	fmt.Printf("ptr指向的地址：%p\n", ptr)
	fmt.Printf("ptr指向的值：%d\n", *ptr)
	fmt.Printf("ptr本身的地址为：%p\n", &ptr)
	/*
		语法糖：由于ptr指向了arr这个数组，所以可以直接用p1来操控数组  ⭐Go语言特性
		ptr指向了谁，这个指针就可以代表谁。
		ptr = arr
	*/
	fmt.Println(ptr[0])

}
