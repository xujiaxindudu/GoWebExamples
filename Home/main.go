package main

import (
	"fmt"
	"net/http"
)

/*
首先需要创建一个Handler接收browser、HTTP client或者API请求的所有传入HTTP连接
格式
func(w http.ResponseWriter,r *http.Request){}
http.Response：编写text/html相应
http.Request:包含请求的所有信息，包括URL或标头字段信息
单独的请求处理程序不能接受来自外部的任何 HTTP 连接。HTTP 服务器必须侦听端口才能将连接传递给请求处理程序。
因为端口 80 在大多数情况下是 HTTP 流量的默认端口，所以该服务器也会监听它。
http.ListenAndServer(":80",nil):
*/

func main() {
	http.HandleFunc("/websockets", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("dudu\n"))
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.Host)
	})

	http.ListenAndServe(":80", nil)
}
