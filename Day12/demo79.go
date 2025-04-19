package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test.go")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name())
	file.Close()
	/*
			无论何时只要你使用 os.Create、os.Open、os.OpenFile 打开了文件，
			都要记得 defer file.Close() 或手动关闭，避免资源泄露或文件占用问题。
			如果不手动关闭，没办法删除
		⭐报错：remove D:\GoWorks\src\xueGo\helloWord\test.go: The process cannot access the file because it is being used by another process.
	*/
	err1 := os.Remove("D:\\GoWorks\\src\\xueGo\\helloWord\\test.go")
	if err1 != nil {
		fmt.Println(err1)
	}
}
