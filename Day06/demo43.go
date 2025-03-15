package main

import "fmt"

// 切片不存储数据，切片指向数组
func main() {

	var arr = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	// 通过数组创建切片 [start,end)
	s1 := arr[2:4] //切片不存储数据，切片指向数组，基于数组，所以cap是从开始到末尾
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1[:cap(s1)]) // [3 4 5 6 7 8]  虽然切片到cap的内容依旧为原数组内容，但是当append的时候 会覆盖原来元素，即原来>len(slice)的切片部分默认值是原数组的值

	/*
		Go 语言规定 cap(slice) = len(底层数组) - 起始索引
		为什么这么规定？
		1.确保切片可以动态扩展⭐
		如果 cap(slice) 只等于 len(slice)，那 append 就没法利用底层数组的剩余空间。
		因为 cap(slice) = 4，所以可以直接利用 arr[5] 的位置，避免不必要的内存分配。如果 cap(slice) 只是 len(slice)，那么 append 必须马上分配新数组，会降低性能。
		2.共享底层数组，提升效率
		切片本质上是对底层数组的引用，而不是新建数组。如果 cap(slice) 只等于 len(slice)，切片就不能利用数组剩余部分，导致资源浪费。
		3.遵循 Go 语言的“零拷贝”设计理念
		如果 cap(slice) = len(slice)，那么 append 必须创建新的数组，增加了额外的内存分配和数据拷贝。
		Go 语言的设计目标之一就是高效，通过共享底层数组，切片可以：
		减少内存分配（append 只有超出 cap 才分配新数组）
		避免不必要的拷贝（多个切片可以引用同一个数组）
	*/

	//修改数组 -> 切片也会发生变化 本质上修改的是切片的底层数组
	arr[2] = 111
	fmt.Println(s1)

	s1 = append(s1, 1, 2, 3, 4, 5) //切片发生扩容之后（扩容->新建一个数组然后把原来的数据copy过去），再对原数组进行修改，就不会影响切片了，因为底层数组变了
	fmt.Println(s1)
	arr[2] = 666
	fmt.Println(s1)

	newSlice := make([]int, 3, 5) //
	copy(newSlice, s1)
	fmt.Println(newSlice)
}
