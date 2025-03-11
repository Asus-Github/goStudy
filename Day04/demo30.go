package main

import "fmt"

/*
defer延迟函数：让这个函数延迟执行，采用压栈操作，即逆序执行defer函数
作用场景：
对文件的I/O操作，对文件.open之后立刻defer 文件.close,由于是压栈操作，所以会在整个程序执行结束之后才会close
*/
func main() {

	f("1")
	fmt.Println("2")
	defer f("3")
	f("4")
	fmt.Println("5")
	defer f("6")
	fmt.Println("7")
	defer fmt.Println("8")
	fmt.Println("\n")
	num := 10

	defer df(num) //此时的defer传进去的num = 10,传参不受影响，只影响执行顺序

	num++
	fmt.Println("EndNum = ", num)

}
func f(str string) {

	fmt.Println(str)
	return
}
func df(num int) int {
	fmt.Println(num)
	return num
}
