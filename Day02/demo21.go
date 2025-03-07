package main

import "fmt"

func main() {
	//练习：if的练习

	var password1 int
	var password2 int

	fmt.Println("请输入密码：")
	//Scan 接受的是一个指针，所以这里需要用变量的地址传入
	fmt.Scan(&password1)
	if password1 == 123456 {
		fmt.Println("请再次输入密码：")
		fmt.Scan(&password2)
		if password2 == 123456 {
			fmt.Println("登陆成功！")
		} else {
			fmt.Println("两次密码不一致！")
		}
	} else {
		fmt.Println("密码错误！")
	}

}
