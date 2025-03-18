package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Go 不是一个面向对象的语言，只能通过方法来近似的模拟面向对象
// 匿名字段+提升字段
type Student struct {
	Person //匿名字段 实现了继承
	//p Person //不属于继承，属于结构体嵌套（组合）
	//name   string
	school string
}

func main() {

	s1 := new(Student)
	s1.Person.name = "asus"
	s1.Person.age = 18
	s1.school = "linda"
	fmt.Println(s1)

	var s2 = Student{Person: Person{"zjj", 18}, school: "linda"}
	fmt.Println(s2)
	fmt.Println(s2)

	//概念：提升字段，只有匿名字段才可以做到
	//Person在Student中是一个匿名字段，Person中的属性name，age就是提升字段
	//所有的提升字段就可以直接使用了，不用在通过匿名字段来点了
	//就近原则
	var s3 Student
	//s3.Person = Person{"zjj", 18}
	s3.name = "zcl"
	s3.age = 19
	s3.school = "linda"

}
