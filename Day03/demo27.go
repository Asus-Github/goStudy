package main

import "fmt"

func main() {

	a, _, b := swap("p1", "p2")

	fmt.Println(a, b)

}

func swap(str1 string, str2 string) (string, string, string) {

	return str2, str1, "test"

}
