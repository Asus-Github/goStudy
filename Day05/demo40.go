package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	var s1 []int //带数字的是数组，不带数字的是切片
	fmt.Println(s1)
	if s1 == nil {
		fmt.Println("切片为空")
	}

	s2 := []int{} //切片是可变长的
	fmt.Println(s2)

	ts := make([]int, 5, 10) //make创建切片

	fmt.Println(len(ts), cap(ts))

	ts = append(ts, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(ts)
	//[0 0 0 0 0 1 2 3 4 5 6 7]  append增加的元素是从len(ts)之后的位置存放的,切片自动扩容到原来的2倍
	fmt.Println(len(ts), cap(ts))

}
