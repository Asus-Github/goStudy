package main

import "fmt"

//空接口 = any
/*
不包含任何方法
因此：所有的结构体都默认实现了空接口 = 所有（结构体）类型都可以使用空接口来接受
因此空接口可以存储任何的类型
*/
//interface{} = any

type A interface {
}
type Dogg struct {
	name string
}

func testXX(a A) {
	fmt.Println(a)
}

func main() {

	d := Dogg{"zjj"}
	testXX(d)

	//map
	m1 := make(map[string]any)
	//m1 := makce(map[string]interface{})
	m1["zjj"] = 1
	m1["asus"] = "love"
	m1["zjj and asus"] = "marry"
	fmt.Println(m1)

	//slice
	s1 := make([]interface{}, 0, 10)
	s1 = append(s1, 1, 2, "3", "zjj")
	fmt.Println(s1)

	//数组
	var arr1 = [4]interface{}{1, 2, 3}
	var arr2 = [4]any{12, 2}
	fmt.Println(arr1, arr2)
}
