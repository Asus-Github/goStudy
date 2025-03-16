package main

import "fmt"

// 结构体嵌套
type userInfo struct {
	Name    string
	Age     int
	Address addrInfo
}
type addrInfo struct {
	Road string
	Id   string
	//Road,Id string
}

func main() {

	user1 := new(userInfo)
	user1.Name = "asus"
	user1.Age = 18
	//user1.Address.Road = "成华大道"
	//user1.Address.Id = "001"
	user1.Address = addrInfo{ //赋值初始化，不要忘记{}前的对象类型⭐
		Road: "成华大道",
		Id:   "001",
	}
	fmt.Println(user1)
}
