package main

import "fmt"

func main() {
	arr := [6]int{0, 1, 2, 3, 4, 5}
	slice := arr[2:5] // [2, 3, 4]

	fmt.Println("原始数组:", arr)   // [0 1 2 3 4 5]
	fmt.Println("原始切片:", slice) // [2 3 4]

	// 修改切片
	// 切片本身不存储数据，它指向的是一个数组
	// 因为切片和数组共享相同的底层数据，所以修改 slice 也会影响 arr：

	slice[0] = 99
	fmt.Println("修改后切片:", slice) // [99 3 4]
	fmt.Println("修改后数组:", arr)   // [0 1 99 3 4 5]
	// slice[0] = 99 实际上修改的是 arr[2]，因为 slice 仍然 指向 arr 的底层数据。

	//如何让 slice 拥有自己的数据（避免修改 arr）？
	// -> demo42
}
