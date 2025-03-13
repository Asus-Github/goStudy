package main

import "fmt"

func main() {

	arr := [10]int{1, 12, 3, 4, 50, 6, -7, 8, 9, 10}
	//交换排序
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				/*temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp*/

				arr[i], arr[j] = arr[j], arr[i] //Go中的交换很简单呀！！！
			}
		}
		fmt.Println(arr)
	}
	fmt.Println("-------------------------")
	fmt.Println(arr)
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	//冒泡排序
	arr2 := [10]int{1, 12, 3, 4, 50, 6, -7, 8, 9, 10}

	for i := 1; i <= len(arr2); i++ {
		for j := 0; j < len(arr2)-i; j++ {
			if arr2[j] > arr2[j+1] {
				arr2[j], arr2[j+1] = arr2[j+1], arr2[j]
			}
		}
		fmt.Println(arr2)
	}
	fmt.Println("-------------------------")
	fmt.Println(arr2)

}
