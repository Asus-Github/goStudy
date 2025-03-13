package main

import "fmt"

// 如果你希望 slice 变成一个完全独立的副本，而不是共享 arr 的底层数组，可以使用 copy：
func main() {
	arr := [6]int{0, 1, 2, 3, 4, 5}
	slice := arr[2:5] // [2, 3, 4]

	// 创建新的独立切片
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	// 修改新切片
	newSlice[0] = 99

	fmt.Println("原始数组:", arr)     // [0 1 2 3 4 5] (未被修改)
	fmt.Println("原始切片:", slice)   // [2 3 4] (未被修改)
	fmt.Println("新切片:", newSlice) // [99 3 4] (独立副本)
}
