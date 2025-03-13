package main

import "fmt"

func main() {

	//定义一个二维数组

	var arr = [3][3]int{
		{1, 2, 3}, //用大括号括起来 不是[]
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	fmt.Println(arr[0])

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d\t", arr[i][j])
		}
		fmt.Printf("\n")
	}

	for index, value := range arr {
		fmt.Printf("%d\t%d\n", index, value)
	}

	//定义一个三维数组

	var arr3 = [3][3][3]int{
		{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}},
	}
	fmt.Println(arr3)

}
