package main

import "fmt"

//接口嵌套

type A interface {
	test1()
}
type B interface {
	test2()
}

type C interface {
	A
	B
	test3()
}
type DogX struct {
	name string
}

func (dog DogX) test1() {
	fmt.Println("dog test1")
}
func (dog DogX) test2() {
	fmt.Println("dog test2")
}
func (dog DogX) test3() {
	fmt.Println("dog test3")
}

func main() {

	//由于DogX对象实现了所有的接口方法，因此它有四种形态：DogX，A，B，C
	dog := DogX{"dudu"}
	dog.test1()
	dog.test2()
	dog.test3()

	fmt.Println("===============")
	var a A
	a = dog
	//向上转型后只能调用自己的方法
	a.test1()
	var b B
	b = dog
	b.test2()
	var c C
	c = dog
	c.test3()

}
