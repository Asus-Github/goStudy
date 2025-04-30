package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	//源文件地址
	stFile := "C:\\Users\\15728\\Pictures\\联想截图\\联想截图_20230621194014.png"
	//目的地址
	destFile := "D:\\GoWorks\\src\\xueGo\\helloWord\\Day15\\server\\demo95.png"
	//临时记录文件
	tempFile := "D:\\GoWorks\\src\\xueGo\\helloWord\\Day15\\temp.txt"

	file1, _ := os.Open(stFile)
	file2, _ := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	file3, _ := os.OpenFile(tempFile, os.O_RDWR|os.O_CREATE, os.ModePerm)

	defer file1.Close()
	defer file2.Close()

	file3.Seek(0, io.SeekStart)
	buf := make([]byte, 1024, 1024) //文件在底层存储中是以字节（byte）为单位的，而不是以整数（int）为单位的。
	fileLen, _ := file3.Read(buf)
	countStr := string(buf[:fileLen])
	count, _ := strconv.ParseInt(countStr, 10, 64) //（Itoa = Int to ASCII）
	fmt.Println(count)

	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	fmt.Println("已经将光标移动到上次传输的地方")

	buf1 := make([]byte, 1024, 1024)

	total := int(count) //Go 中整数类型不能混加，int ≠ int64，需要手动显式转换！
	fmt.Printf("%T,%T\n", total, count)
	//index := 0
	for {
		readNum, err := file1.Read(buf1)
		//index++
		//fmt.Println(index)
		if err == io.EOF {
			file3.Close()
			fmt.Println("文件传输完毕")
			os.Remove(tempFile) //Remove传输的是一个路径
			break
		}
		writeNum, _ := file2.Write(buf1[:readNum])

		total += writeNum //Go 中整数类型不能混加，int ≠ int64，需要手动显式转换！

		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(int(total)))
	}
}

/*
假设你的文件大小是 2500 字节，而你的缓冲区是 1024 字节。

读取过程：
第一次读：读了 1024 字节，readNum = 1024, err = nil

第二次读：读了 1024 字节，readNum = 1024, err = nil

第三次读：只剩 452 字节，读了 452 字节，readNum = 452, err = nil（❗注意这里仍然不是 EOF）

第四次读：什么都读不到了，readNum = 0, err = io.EOF

只要这次读取确实读到了数据（readNum > 0），即使已经到了文件末尾，err 也还是 nil。
*/
