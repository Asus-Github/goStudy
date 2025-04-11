package main

import (
	"errors"
	"fmt"
)

func main() {

	// 自己定义一个错误
	// 1、errors.New("xxxxx")
	// 2、fmt.Errorf()
	// 都会返回  error 对象， 本身也是一个类型

	age_err := setAge(-1)
	if age_err != nil {
		fmt.Println(age_err)
	}
	fmt.Println(age_err)
	//errInfo2 := fmt.Errorf("我是一个错误信息：%d\n", 404)
	errInfo := fmt.Errorf("我是一个错误信息 %d", 500)
	// 处理这个错误
	if errInfo != nil {
		fmt.Println(errInfo)
	}
}

func setAge(age int) error {

	if age < 0 {
		age = 3
		return errors.New("年龄不合法")
	}

	return nil

}
