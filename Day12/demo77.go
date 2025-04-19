package main

import (
	"fmt"
	"os"
)

func main() {
	//  Windows 下路径里的 \ 需要转义为 \\
	file, err := os.Stat("D:\\GoWorks\\src\\xueGo\\helloWord\\Day12\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name())
	fmt.Println(file.Size())
	fmt.Println(file.ModTime())
	fmt.Println(file.IsDir())
	fmt.Println(file.Mode()) //权限
}
