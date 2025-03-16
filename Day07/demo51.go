package main

import "fmt"

// 指针
func main() {

	var a int = 5
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 10
	fmt.Println(a)

	var c *int
	c = &a

	fmt.Printf("变量存储的指针地址%p\n", c)
	fmt.Printf("变量本身的地址%p\n", &c)
	fmt.Printf("变量存储的值%d\n", *c)

	//fmt.Println(*c)

	var ptr **int //第一个*表示是指针类型，*int表示的是ptr指向的是一个指针类型 即ptr是指针的指针
	ptr = &c      //√ c本身是指针类型，ptr指向指针的地址
	//ptr = &a //报错 因为a不是一个指针类型
	fmt.Printf("ptr存储指针的地址为：%p\n", ptr)
	fmt.Printf("ptr的地址为%p\n", &ptr)
	fmt.Printf("ptr变量存储的地址为%p\n", *ptr) //即c的地址
	fmt.Printf("ptr存储的地址存储的值为：%d\n", **ptr)

	**ptr = 123
	fmt.Println(a)
}
