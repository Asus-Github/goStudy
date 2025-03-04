package main

import "fmt"

func main() {

	a := 5.65
	var b = int(a)
	//浮点数转整数 向下取整，不是四舍五入
	fmt.Println(a, b)
	fmt.Printf("%T %.2f\n", a, a)
	fmt.Printf("%T %d\n", b, b)
	// 保留几位小数才会自动四舍五入
	fmt.Printf("%.1f", a)
}
