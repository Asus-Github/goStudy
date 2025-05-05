package main

import (
	"fmt"
	"net/http"
)

func main() {
	//请求处理函数
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	//放在最后监听端口
	http.ListenAndServe("localhost:8080", nil)

}

/*
解释一下hello函数的两个参数，为什么第二个是指针类型?

req *http.Request
这是一个指向 http.Request 结构体的指针，表示客户端发来的 HTTP 请求。它包含了很多请求信息，比如请求的方法（GET、POST）、URL、Header、Body、表单数据等。
使用指针类型的原因是：
节省内存开销：http.Request 结构体可能非常大，使用指针避免复制整个结构体。
共享修改结果：如果需要修改请求（比如修改表单值、上下文等），通过指针可以直接作用在原始数据上。
一致性：Go 的标准库中很多结构体都采用指针方式传参，符合语言习惯。

w http.ResponseWriter
这是一个接口类型，用来构建并发送 HTTP 响应给客户端。你可以通过它写出文本、设置响应头、设置状态码等。
*/
func hello(resp http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	fmt.Println(req.URL)
	fmt.Println(req.Method)
	fmt.Println(req.RemoteAddr)
	/*
			为什么是一个随机端口号（比如 50296）？
			HTTP 是基于 TCP 的。
			当客户端（你本机的浏览器）向服务器（也是你本机）发起请求时，系统会分配一个临时端口用于建立连接。
			这个端口号就是 RemoteAddr 中显示的端口。
		RemoteAddr	客户端的 IP+端口（临时生成的）
		8080	你写的服务器监听的端口
	*/
	resp.Write([]byte("hello world"))
}

func login(resp http.ResponseWriter, req *http.Request) {

	mysqlUsername := "asus"
	mysqlPassword := "123456"
	fmt.Println("接收到了客户端请求！")
	reqData := req.URL.Query()
	username := reqData.Get("username")
	password := reqData.Get("password")

	if username != mysqlUsername || password != mysqlPassword {
		resp.Write([]byte("请重新输入账号密码！"))
		return
	} else {
		resp.Write([]byte("ok！登陆成功！"))
		return
	}

}

func register(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("接收到了register请求！")
	req.ParseForm() //⭐
	username := req.PostForm.Get("username")
	password := req.PostForm.Get("password")
	fmt.Println(username, password)
	fmt.Println(req.FormValue("username"))
}

/*
| 方式                       | 是否需要手动调用 `ParseForm()` |
| ------------------------- | --------------------- -  |
| `req.FormValue("key")`    | **不需要**                |
| `req.PostForm.Get("key")` | **需要**                 |
| `req.Form.Get("key")`     | **需要**                 |


*/
