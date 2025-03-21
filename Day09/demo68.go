package main

import "fmt"

// 创建了一种新类型！
// 这是定义了一个新类型 MyInt，是int转换过来的，和int一样，但是不能通int发生操作，类型不同
type MyInt int

// 2、type 起别名
// - type xxx = 类型 ，将某个类型赋值给 xxx，相当于这个类型的别名
type diyint = int

func main() {

	var a MyInt
	a = 1
	fmt.Println(a)
	fmt.Printf("%T\n", a) //main.MyInt 全新的类型
	var b int
	b = 1
	fmt.Printf("%T\n", b) //int
	//p := a+b 不同类型不能计算
	p := int(a) + b //强转
	fmt.Println(p)

	var c diyint = 30
	fmt.Println(c)
	fmt.Printf("%T\n", c) //int
}

/*
type关键字的理解：
1、type 定义一个类型
 - type User struct 定义结构体类型
 - type User interface 定义接口类型
 - type Diy (int、string、....) 自定义类型，全新的类型

2、type 起别名
 - type xxx = 类型 ，将某个类型赋值给 xxx，相当于这个类型的别名
 - 别名只能在写代码的时候使用，增加代码的可阅读性。
 - 真实在项目的编译过程中，它还是原来的类型。
*/
