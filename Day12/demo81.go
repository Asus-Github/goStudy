package main

import (
	"fmt"
	"os"
)

// 读取  file.Read([]byte)  ，将file中的数据读取到  []byte 中， n，err  n读取到的行数，err 错误，EOF错误，就代表文件读取完毕了
func main() {

	file, err := os.Open("D:\\GoWorks\\src\\xueGo\\helloWord\\Day12\\test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bs := make([]byte, 4, 1024)
	n, err := file.Read(bs)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(string(bs))
	n, err = file.Read(bs)
	fmt.Println(string(bs))
	fmt.Println(n)
	fmt.Println(err)
	//file.Seek(0, 0)
	n, err = file.Read(bs)
	/*
		abcdefghi
		string(bs) 会输出 "ifgh"，但实际上只读了 'i'，剩下的 fgh 是上次的残留
		err = io.EOF 表示读到文件末尾
		❗ 关键点：残留数据（坑）
		因为 bs 是复用的，并没有每次都用新的切片，所以当你只读取了 1 个字节（比如 'i'），剩下的 bs[1:] 还是保留着上次读的内容。
		你打印 string(bs) 实际上是 "i" + "fgh"，让人误以为你读到了 ifgh，其实只读了 i！
		✅ 正确做法（建议）
		只打印 bs[:n]，即真实读取的部分：
	*/
	fmt.Println(string(bs))
	fmt.Println(string(bs[:n]))
	fmt.Println(n)
	fmt.Println(err)
	/*

		utils.Copy(source, dest, 1024) 是啥？
		虽然 utils.Copy 不是 Go 标准库的函数，应该是你项目自定义的工具函数，但它很可能实现了文件复制或者流数据复制功能，类似下面这种封装：
		✅ 参数 1024 是什么意思？
		这个就是传入的 bufSize，表示：
			每次从 source 读取和写入到 dest 的最大字节数是 1024 字节（即 1KB）
			🔍 为什么要用缓冲区？⭐⭐⭐
			避免一次性读/写占用过多内存
			控制内存使用（比如用于大文件时）
			提高效率（读写固定大小通常更快）

	*/
}
