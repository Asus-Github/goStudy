package main

import "fmt"

func main() {

	var str string
	str = "hello world"
	fmt.Printf("%T %s \n", str, str)

	// 字符和字符串同C
	A := 'A'
	B := "A"
	fmt.Printf("%T %d \n", A, A)
	fmt.Printf("%T %s \n", B, B)

	// ``可以写一些很长的带换行的字符串
	C := `111111
111
11
das

fas

`
	fmt.Printf("%T %s \n", C, C)

}
