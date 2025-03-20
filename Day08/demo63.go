package main

//同一个包下所有东西都不能重名的

//多态

type AnimalX interface {
	sleep()
	eat()
}

type DogS struct {
	name string
}

func (d DogS) sleep() {
	println(d.name, "正在睡觉...")
}
func (d DogS) eat() {
	println(d.name, "正在吃饭...")
}

func testX(a AnimalX) {
	a.sleep()
	a.eat()
}

// 接口的实现类都拥有多态特性：除了自己本身还是他对应接口的类型。
func main() {

	d := DogS{"dudu"}
	d.sleep()
	d.eat()
	//testX传入的是一个Animal类型，由于DogS类实现了所有的接口定义方法，所以这个类实现了这个接口，所以DogS类型=Animal类型
	//即d 有两重身份 除了自身类型DogS，还有它实现的接口类型Animal
	testX(d)

	var animal AnimalX
	animal = d
	animal.sleep()
}
