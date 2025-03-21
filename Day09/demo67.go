package main

import "fmt"

// 接口断言
// 检查一个接口类型的变量是不是符合你的预期值。
// 被断言的对象必须是接口类型，否则会报错

type I interface{}

// 通过switch来判断  switch i.(T)

func testAsserts(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case I:
		fmt.Println("变量为I类型")
	case nil:
		fmt.Println("变量为nil类型")
	case interface{}:
		fmt.Println("变量为interface{}类型")

	// .....
	default:
		fmt.Println("未知类型")

	}
}

func main() {
	testAsserts("string")
	testAsserts(1)

	var i I // 默认值为 nil
	testAsserts(i)
	var i2 I = 1 // 只有赋值了之后，才是对应的类型
	testAsserts(i2)
	//testAsserts(I)
}
