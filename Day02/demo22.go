package main

import "fmt"

func main() {

	var score int = 60
	var msg string
	switch score { //此处case后面不可以写范围，范围返回bool值，但是score是int值，不匹配
	case 100, 90, 85: //case后面可以写多个条件
		msg = "A"
	case 80, 70:
		msg = "B"
	case 60:
		msg = "C"
	default:
		msg = "不及格"
	}
	fmt.Println(msg)

	var flag bool

	switch flag { //此处可以写范围，因为flag是个bool类型
	case score > 100:
		flag = true
	case score > 80:
		flag = false
	} //可以没有default
	fmt.Println(flag)

	var fallThroughFlag bool

	switch fallThroughFlag { //此处可以写范围，因为flag是个bool类型
	case score > 100:
		fallThroughFlag = true
		fallthrough //case穿透，加上这个关键字之后，则会强制执行下面这一个case（只有一个）
	case score > 80:
		fallThroughFlag = false
		break //break退出
		fmt.Println("test break happen!")
	case score > 60:
		fallThroughFlag = true
	} //可以没 有default
	fmt.Println(fallThroughFlag)
}
