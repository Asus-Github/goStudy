package main

import "fmt"

type Student struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	var stu Student
	stu.Name = "asus"
	stu.Age = 18
	stu.Sex = "男"

	tstu := stu
	tstu.Name = "zjj"

	fmt.Println(stu)
	fmt.Println(tstu) //结构体是值传递,只是把值copy给了tstu

	var stuPtr *Student
	stuPtr = &stu //用指针指向的是内存地址，所以通过指针stuptr修改变量值可以完成原变量的修改
	stuPtr.Name = "updataAsus"
	fmt.Printf("stu的地址为：%p\n", &stu)
	fmt.Printf("指针指向的地址为：%p\n", stuPtr)

	fmt.Println("修改stuPtr的Name值之后stuName的值为：", stu.Name)
	fmt.Println(stuPtr)
	fmt.Println(*stuPtr)
	fmt.Println((*stuPtr).Name)
	fmt.Println(stuPtr.Name) //Go语言特有语法糖：可以直接用变量表示指针 即 (*stuPtr).Name = stuPtr.Name

	stu2 := new(Student) //Go语言提供了创建对象的函数 new() 该函数返回一个指针，由于语法糖，所以我们可以直接使用变量名
	stu2.Name = "asus"

	updataUserInfo(stu2, "zjj")
	fmt.Println(stu2.Name)
	/*
		// 内置函数 new 创建对象。new 关键字创建的对象，都返回指针，而不是结构体对象,
		// func new(Type)*Type
		// 通过这种方式创建的结构体对象更加灵活，突破了结构体是值类型的限制。
		// 使得函数调用更方便
	*/

}

func updataUserInfo(student *Student, newName string) {
	student.Name = newName
}
