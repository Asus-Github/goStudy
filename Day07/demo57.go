package main

import "fmt"

// 匿名结构体 : 即没有名字的结构体
func main() {

	s1 := struct {
		name string
		age  int
	}{ //匿名结构体，直接可以在函数内部创建出来，创建后就需要赋值使用
		name: "asus",
		age:  18,
	}
	fmt.Println(s1)

	var teacher = Teacher{"asus", 18}
	// 如何打印这个匿名字段，默认使用数据类型当做字段名称
	fmt.Println(teacher.string, teacher.int)
}

// 匿名字段
// 结构体中的匿名字段，没有名字的字段，这个时候属性类型不能重复，因为默认使用数据类型当做字段名称
type Teacher struct {
	string
	int
}
