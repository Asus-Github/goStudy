package main

import "fmt"

func main() {

	map1 := make(map[int]string)
	map1[1] = "hello"
	// map取值 隐藏返回值ok->返回是否存在该key对应的value值
	value, ok := map1[1]
	value2, ok2 := map1[2]
	fmt.Println(value, ok)
	if ok {
		fmt.Println("存在该key对应的value")
	} else {
		fmt.Println("不存在该key对应的value")
	}

	fmt.Println(value2, ok2)

	map1[3] = "xxxxxxxxx"
	fmt.Println(map1)
	delete(map1, 3) //map删除函数，delete(map,key)
	fmt.Println(map1)
}
