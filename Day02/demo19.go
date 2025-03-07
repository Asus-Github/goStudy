package main

import "fmt"

func main() {

	/*a := 0110 //
	/*在 Go 语言中，如果数值字面量以 0 开头，则被解释为八进制数。
	在你的代码中，`a := 0110` 是一个八进制数，其十进制值计算如下：
	因此，当你打印变量 `a` 时，输出的是 72。*/
	//b := 1001

	a := 5 //101
	b := 1 //001
	//100   ->   4
	fmt.Println(a, b)
	//位清空  当b为0的时候取a的值，当b为1的时候取0

	//关于位清空运算符应用场景 详见博客：https://blog.csdn.net/weixin_51461002/article/details/146097684?fromshare=blogdetail&sharetype=blogdetail&sharerId=146097684&sharerefer=PC&sharesource=weixin_51461002&sharefrom=from_link
	fmt.Println(a &^ b)
}
