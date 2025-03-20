package main

import "fmt"

// 区分方法和函数
/*
	方法:
	某个类的行为功能，需要指定调用者
	一段独立的代码功能，必须要使用⭐调用者来调用
	多个类的方法可以重名，单个类不行
	方法是某个类的动作
	函数：
	一段独立的代码功能，可以直接调用
	命名完全不能冲突
	函数是一个特殊的类型
*/

type Dog struct {
	name string
	age  int
}

type Cat struct {
	name string
	age  int
}

// GO语言中的方法：func + 调用者 + 方法名
func (dog Dog) eat() {
	fmt.Println("Dog eating...")
}

func (cat Cat) eat() {
	fmt.Println("Cat eating...")
}

func main() {

	dog := Dog{
		"doudou",
		18,
	}
	dog.eat()

	cat := Cat{"didi", 18}
	cat.eat()

}
