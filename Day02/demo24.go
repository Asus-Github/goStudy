package main

import "fmt"

func main() {

	for i := 1; i <= 5/2+1; i++ {
		for j := 5 - i; j >= 1; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 5 / 2; i >= 1; i-- {
		for j := 5 - i; j >= 1; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
