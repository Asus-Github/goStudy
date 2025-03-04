package main

import "fmt"

func main() {

	/*
		我们一般认为iota定义为常量使用
		iota 本身就是个常量，所以只能定义常量的时候使用
	*/
	const (
		/* 每次定义新的常量iota会自动 +1 ，而不是每次使用iota*/
		a = iota
		b = iota
		c = iota
		d = 3
		e = iota
		f
		g
		h = iota
	)

	fmt.Println(a, b, c, d, e, f, g, h)
	const (
		i = iota
		j
		k
	)
	fmt.Println(i, j, k)
}
