package main

import "fmt"

func main() {

	// byte 0-255 byte = uint8
	var num1 byte = 255

	fmt.Println(num1)
	fmt.Printf("%T %d\n", num1, num1)

	//rune = uint32
	var num2 rune = 1000000000
	fmt.Printf("%T %d", num2, num2)
}
