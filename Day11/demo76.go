package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//num1 := rand.Intn(100)
	//fmt.Println(num1)
	//num2 := rand.Intn(100)
	//fmt.Println(num2)
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100)) // 每次运行都会输出一样的结果
	}
}
