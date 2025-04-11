package main

import "fmt"

type myError struct {
	code int
	msg  string
}

/*
type error interface {
	Error() string
}
*/
//实现error接口 = 实现该接口中所有的方法
/*
只有 *myError （也就是 myError 的指针）有 Error() 方法。
换句话说，myError 本身没有实现接口，只有 *myError 实现了接口。
&myError{...} 是一个指针。
它实现了 Error() 方法。
所以它可以当作 error 类型返回。
*/
func (e *myError) Error() string {
	return fmt.Sprintf("code:%d msg:%s", e.code, e.msg)
}

func test(i int) (int, error) {
	if i != 0 {
		return i, &myError{code: i, msg: "ok"}
		//return i, myError{code: i, msg: "ok"}  // 注意：不是指针了
		//编译会报错：❌ 因为 myError（值）没有实现 Error() 方法。
	}
	return i, nil
}
func main() {
	re, err := test(1)
	fmt.Println(re, err)
}
