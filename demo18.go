package main

import "fmt"

func main() {

	var a int = 2
	var b int = 7
	var c int
	c = b / a
	c = b / 2.0
	fmt.Printf("%T %v\n", c, c)
	//d := a + 1.6
	//var d float64 = float64(a*b + c*c + 1.6)
	//fmt.Printf("%T %.2f\n", d, d)

}
