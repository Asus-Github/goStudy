package main

import "fmt"

type Animal struct {
	name string
	age  int
}

type CatX struct {
	Animal
	color string
}
type DogX struct {
	Animal
	color string
}

func (dog DogX) say() {
	fmt.Println(dog.name, "正在wangwang...")
}

func (cat CatX) say() {
	fmt.Println(cat.name, "正在miaomiao...")
}

func (animal Animal) eat() {
	fmt.Println(animal.name, "正在吃...")
}

func main() {

	dog := DogX{Animal{"dudu", 18}, "黄色"}

	dog.say()
	//子类可以操作父类的方法，父类可以操作子类的吗？不可以
	dog.eat()
	//父类不能调用子类自己的扩展的方法
	//a := Animal{"test", 18}
	//a.say
}
