package main

import "fmt"

func main() {

	a, _ := swap("p1", "p2")

	fmt.Println(a)

}

func swap(str1 string, str2 string) (string, string) {

	return str2, str1

}
