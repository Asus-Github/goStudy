package main

import "fmt"

// map是引用类型不同变量可以共享同一个 map 的底层数据（不会复制整个 map）。
// 多个 map 变量可以指向同一块内存。
func main() {

	user1 := make(map[string]string)
	user1["name"] = "asus"
	user1["age"] = "18"
	user1["sex"] = "男"
	user1["addr"] = "山东"

	user2 := make(map[string]string)
	user2["name"] = "zjj"
	user2["age"] = "18"
	user2["sex"] = "女"
	user2["addr"] = "黑龙江"

	userData := make([]map[string]string, 0, 3) //对比 make([]int,0,3) int类型的切片，  []map[string]string为map[string]string类型的切片
	/*userData := make([]map[string]string, 1, 3) //[map[] map[addr:山东 age:18 name:asus sex:男] map[addr:黑龙江 age:18 name:zjj sex:女]]

	//userData[0] = make(map[string]string) //对预先分配len 分配内存

	userData[0]["name"] = "asus" //❌ panic: assignment to entry in nil map
	/*
		为什么 userData[0]["name"] = "asus" 会 panic？
		make([]map[string]string, 1, 3) 创建了一个长度为 1，容量为 3 的切片。
		但 map 本身是引用类型，Go 不会自动初始化 map，而是默认将切片内的 map 设为 nil。
		userData[0] 目前是 nil，尝试向 nil map 赋值会导致 panic: assignment to entry in nil map。

		nil map 没有分配空间，m == nil。
		len(m) == 0 是合法的。
		但不能直接赋值，否则会 panic。

		var m map[string]string // m 是 nil 只是定义了 并没有分配内存 ⭐
		fmt.Println(m == nil)   // ✅ 输出: true⭐
		fmt.Println(len(m))     // ✅ 输出: 0（合法操作）
		// ❌ 试图向 nil map 赋值，会 panic
		m["name"] = "asus" // panic: assignment to entry in nil map

		m := make(map[string]string) // ✅ 通过 make 初始化 -> 相当于分配内存 ⭐
		fmt.Println(m == nil)        // ✅ 输出: false（因为已初始化）⭐
		fmt.Println(len(m))          // ✅ 输出: 0（没有元素）


	*/
	//fmt.Println(userData[0])
	/*
		len非0的时候，自动赋值为nil
		强行取值：panic: runtime error: index out of range [0] with length 0
		需要对预分配len 分配内存 即 make(map[string]string)
	*/

	userData = append(userData, user1)
	userData = append(userData, user2)
	fmt.Println(userData[0])
	fmt.Println(userData)

}

/*
map 是引用类型，但它本身是一个指针（引用）指向的结构，Go 不会自动给它分配内存，需要手动 make 或 {} 初始化。
*/
