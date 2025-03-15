package main

import "fmt"

// map集合
func main() {
	//创建一个map，也是一个变量，数据类型是 map (%T)
	var map1 map[int]string //map加上key和value

	if map1 == nil {
		fmt.Println("map1 == nil")
	}
	map2 := make(map[int]string) //make创建map
	fmt.Println(map2)
	//map初始化
	var map3 = map[int]string{1: "test1", 2: "test2"}
	fmt.Println(map3)

	fmt.Printf("%T\n", map3)
}
