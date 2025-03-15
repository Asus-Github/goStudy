package main

import "fmt"

/*
⭐1.map是无序的，每次打印出来的map可能都不一样，它不能通过index获取，只能通过key来获取
2.map的长度是不固定的，是引用类型的
3.len可以用于map查看map中数据的数量，但是cap无法使用
4.map的key可以是所有可以比较的类型。布尔类型，整数，浮点数，字符串
*/
func main() {

	map1 := map[string]string{
		"lhs": "asus",
		"zjj": "pgd",
		"czl": "bj",
	}
	//map只能通过for range来遍历

	for key, value := range map1 {
		fmt.Println(key, value) //map是无序的，每次打印出来的顺序都不一样
	}
	fmt.Println(len(map1))
}
