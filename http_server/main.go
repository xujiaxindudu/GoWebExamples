package main

import (
	"fmt"
	"net/http"
)

/*
创建一个基本的http服务器。一个基本的http服务器需要具备几个关键工作
处理动态请求:http.HandleFunc,他的第一个参数需要一个匹配的路径，第二个参数是一个要执行的函数，
处理静态资产：
接受连接：http.ListenAndServer(":80",nil)
*/
func main() {
	http.HandleFunc("/web", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome to my website")
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":80", nil)
}
