package main

import "fmt"

func main() {
	//定义数组的时候初始化
	var arr = [5]int{1, 2, 3, 4, 5} //可以理解为 先定义了arr变量，而这个变量有5个数 所以用[5]个来赋值 类型为 int 值为{}
	fmt.Println(arr)

	//语法糖快速初始化
	arr2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(arr2) //默认情况下，fmt.Println 打印 string 数组时不会带上双引号，而是直接显示字符串内容，使得 string 和 int 数组的打印结果在表面上看起来相同。
	fmt.Printf("%q\n", arr2)

	//-------------------------------
	str := "Hello\tWorld\nGo!"
	fmt.Printf("普通输出: %s\n", str)
	fmt.Printf("带引号且带转义: %q\n", str) //%q 会显示转义字符（如 \t、\n）并加上双引号。
	//-------------------------------

	/*
		比较特殊的点
		数据如果来自用户，我不知道用户给我多少个数据，数组
		代表数组的长度
		...赋值，自动推导长度
		Go的编译器会自动根据数组的长度来给
		注意点:这里的数组不是无限长的，也是固定的大小，大小取决于数组元素个数。
	*/
	var usrData = [...]int{1, 2, 3, 5, 4, 5, 6, 6, 6, 6, 7, 7, 78, 8, 9}
	fmt.Println(usrData)
	fmt.Println(len(usrData))

	//数组默认值 {index:data}
	var mRarr = [5]int{1: 100, 4: 200}
	//mRarr = [5]int
	fmt.Println(mRarr)

}
