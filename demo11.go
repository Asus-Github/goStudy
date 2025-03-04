package main

import "fmt"

func main() {
	/*bool 类型只有true or false 不可以用0 or 1*/
	var b1 bool = true
	var b2 bool = false
	fmt.Println(b1, b2)
	/* %T 类型，%t 数值*/
	fmt.Printf("%T,%t\n", b1, b1)
}
