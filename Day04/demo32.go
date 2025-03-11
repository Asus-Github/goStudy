package main

import "fmt"

// 匿名变量（没有名字的变量）
// 匿名函数（没有名字的函数）
func main() {

	f12()
	f2 := f12 //f2 f12本质上指向了同一个内存空间，空间中的代码一致
	f2()

	//匿名函数
	func() {
		fmt.Println("我是匿名函数1号")
	}() //加上()就不会报错，-> 因为加上()表示调用 不赋值就一次调用

	p := func() {
		fmt.Println("我是匿名函数2号")
	} //不加()表示一个变量，可以进行赋值，用语法糖赋值可以多次调用了

	p()

	func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2)

	add := func(a int, b int) int {
		res := a + b
		return res
	}
	fmt.Println(add(1, 2))

	ans := func(a, b int) int {
		return a + b
	}(1, 3)
	fmt.Println(ans)

}

func f12() {
	fmt.Println("我是f12")
}
