package main

import "fmt"

/*
函数的本质：其实也是一个变量,可以进行赋值
*/

func main() {

	d(1, 2, "sd")
	// Go语言中函数 加（）表示调用这个函数，不加（）代表它本身就是一个变量
	tempd := d
	var fd func(int, int, ...string) int //把f定义成与d相同的函数类型，则可以把d赋值为f
	fd = d
	fmt.Println(fd(1, 2, "pd"))
	fmt.Println(tempd(1, 2, "pd"))

}

// 可变参数也是要有变量名字的 str[] 用数组下表取参0~len(str)-1
func d(num1, num2 int, str ...string) int {
	return 11
}
