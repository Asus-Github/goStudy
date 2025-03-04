package main

import "fmt"

func main() {
	var f1 float64
	f1 = 3.14

	var f2 float32
	f2 = 5
	//默认输出保留6位小数
	//保留几位小数，以及自动四舍五入等同C语言
	fmt.Printf("%.2f %f\n", f1, f2)

}
