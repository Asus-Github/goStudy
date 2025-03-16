package main

import "fmt"

// 指针函数 返回一个指针的函数
func main() {

	ptr1 := backPtr1()

	fmt.Println(ptr1)

	arr := [5]int{1, 2, 3, 4, 5}
	backPtr2(&arr)
	fmt.Println(arr)

}

// 通过此种方式，可以把函数中的数组从栈上逃逸到堆上，因为没有被销毁
// 补充：什么时候定义的变量会销毁，没有任何东西再指向它的时候，垃圾回收（GC）会回收。
func backPtr1() *[5]int {
	arr := [5]int{1, 2, 3, 4, 5}
	return &arr
}

func backPtr2(arr *[5]int) { //参数是指针类型，所以我要传进来一个地址 &arr
	arr[0] = 100
}
