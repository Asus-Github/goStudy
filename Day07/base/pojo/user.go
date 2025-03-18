package pojo

type Boss struct { //结构体名称首字母小写也不能被导出，大写才可以 同下
	Name  string
	Age   int
	money int //成员首字母大写=可以被导出使用，首字母小写=不可以被导出使用
}
