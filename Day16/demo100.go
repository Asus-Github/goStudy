package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	urlStr := "http://localhost:8080/login"

	data := url.Values{}
	data.Set("username", "asus")
	data.Set("password", "123456")

	urlNew, _ := url.ParseRequestURI(urlStr) //解析一个字符串形式的 URL，并检查它是否是一个合法的 HTTP 请求目标 URI。 返回一个URL指针（对象）
	urlNew.RawQuery = data.Encode()          //查询参数（RawQuery）：username=asus&password=123456
	/*
		RawQuery 表示查询参数（不含 ?）
		String() 负责生成完整的 URL 字符串 加上'?'
		url.URL 是一个结构体，而它的 String() 方法会在打印时自动调用。
		url.URL.String() 的实现中有这么一段逻辑（简化描述）：
			if u.RawQuery != "" {
				buf.WriteByte('?')
				buf.WriteString(u.RawQuery)
			}
	*/
	fmt.Println(urlNew)

	resp, _ := http.Get(urlNew.String())
	defer resp.Body.Close() //及时关闭
	fmt.Println(resp.Body)
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println(string(buf[:n]))
			break
		}
	} //分块读取，适合大数据
	fmt.Println("读取完毕！")
	//buf, _ := io.ReadAll(resp.Body) 直接读取全部，适合小数据
	//fmt.Println(string(buf))

	/*
		第一次打印 resp.Body：这会打印的是 resp.Body 本身，通常会显示类似于：&{<内存地址>}，这是一个 io.Reader 接口，指向 HTTP 响应体的数据流。
		第二次读取：你调用 io.ReadAll(resp.Body) 从 resp.Body 读取了响应体的全部内容，并将其存储到 buf 中。然后 fmt.Println(string(buf)) 打印了 buf 中的实际数据内容。
	*/

	/*resp, err := http.PostForm(urlStr, data)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(resp)
	}*/
}
