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
}
