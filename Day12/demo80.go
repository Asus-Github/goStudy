package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("D:\\GoWorks\\src\\xueGo\\helloWord\\Day12\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println(file)
	/*
		这个第三个参数 perm 是一个文件权限（os.FileMode 类型），只有在文件“被创建”时才生效
		os.OpenFile("example.txt", os.O_CREATE|os.O_WRONLY, 0644)
		如果 example.txt 不存在，它会被创建。
		创建时的权限就是 0644。
		如果文件已经存在，这个权限参数会被忽略。

	*/
	file2, er2 := os.OpenFile("D:\\GoWorks\\src\\xueGo\\helloWord\\Day12\\test.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)

	if er2 != nil {
		fmt.Println(er2)
		return
	}
	defer file2.Close()
	fmt.Println(file2)
	/*
		问题 | 答案
		os.Open 和 os.OpenFile 区别？ | os.Open 是只读简化版，os.OpenFile 是万能打开函数，可读可写可创。
		file 是文件内容吗？ | 不是，它是一个指向文件对象的指针，用于操作该文件。
		类似：&{0xc0000ae000}
	*/
}
