package main

import (
	"fmt"
	"xueGo/helloWord/Day07/base/pojo"
)

func main() {
	var boss pojo.Boss

	boss.Name = "asus"
	boss.Age = 18
	//boss.money//小写字母不能被导出使用
	fmt.Println(boss)
}
