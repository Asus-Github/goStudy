package main

import "fmt"

//接口：USB、typec、插座
//1、Go语言提供了接口数据类型。
//2、接口就是把一些共性的方法集合在一起定义。
//3、如果有实现类将接口定义的方法全部实现了，那么就代表实现了这个接口
//4、隐式实现GO，假设A实现了B接口中的所有方法，不需要显示声明
//5、接口是方法的定义集合，不需要实现具体的方法内容。名字约束

//接口只定义方法，并不是实现方法

type USB interface {
	input()
	output()
}

type Mouse struct {
	name string
}

// 结构体定义了接口的所有方法，说明实现了这个接口，否则不算完成了这个接口
func (mouse Mouse) input() {
	fmt.Println(mouse.name, "正在input...")
}

func (mouse Mouse) output() {
	fmt.Println(mouse.name, "正在output...")
}

type Key struct {
	name string
}

func (key Key) input() {
	fmt.Println(key.name, "正在input...")
}

func (key Key) output() {
	fmt.Println(key.name, "正在output...")
}

func test(usb USB) {
	usb.input()
	usb.output()
}
func main() {

	m1 := Mouse{"罗技"}
	//⭐⭐⭐test传入的是一个usb类型，如果一个结构体实现了这个接口的所有方法，那这个结构体就是这个接口类型的
	//所以此处Moust=USB（interfacec）
	test(m1)
	/*
		m1.input()
		m1.output()
	*/
	fmt.Println("============")
	k1 := Key{"雷蛇"}
	test(k1)
	fmt.Println("============")
	k1.input()
	k1.output()
	fmt.Println("============")

	//定义了高级类型，赋值=向上转型
	var usb USB
	usb = k1
	//接口无法使用实现类的属性
	//usb.name
	fmt.Println(usb)
	usb.input()
}
