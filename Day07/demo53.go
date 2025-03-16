package main

import "fmt"

// 指针数组 顾名思义 存储指针的数组
func main() {

	a := 1
	b := 2
	c := 3
	d := 4

	var ptrArr = [4]*int{&a, &b, &c, &d} //*int是一个指针类型，变成指针数组只需要在前面加上个数[size] 即可

	//通过指针数组修改数值
	fmt.Printf("a变量的地址为：%p\n", ptrArr[0])
	fmt.Printf("a变量的值为：%d\n", *ptrArr[0])

	*ptrArr[0] = 100

	fmt.Printf("修改后的a变量值为%d\n", *ptrArr[0])

	a = 200
	fmt.Printf("修改后的a变量值为%d\n", *ptrArr[0])

}
