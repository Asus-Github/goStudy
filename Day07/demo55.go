package main

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

func main() {
	//创建对象方式一
	var user User //创建了自己的类型，在定义的时候可以使用非基本类型了！
	user.name = "asus"
	user.age = 18
	user.sex = "男"
	fmt.Println(user)

	//创建对象方式二 语法糖
	user2 := User{}
	user2.name = "zjj"
	user2.age = 18
	fmt.Println(user2)

	//创建对象方式三⭐
	user3 := User{
		name: "linda",
		age:  18,
		sex:  "男",
	}
	fmt.Println(user3)

	var user4 User //不赋值，默认int为0，string为""空
	fmt.Println(user4)

	//创建对象方式四，不声明属性，需要值一一匹配
	user5 := User{"wls", 12, "男"}
	fmt.Println(user5)
}
