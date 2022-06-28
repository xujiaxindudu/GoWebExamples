package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//html/template 为HTML提供了丰富的模板语言。它主要用于web应用程序中，以在客户端浏览器中以结构化方式显示数据。Go的模板语言的一大好处是数据的自动转意
//无需担心 XSS 攻击，因为 Go 在将其显示给浏览器之前会解析 HTML 模板并转义所有输入。
type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func handle(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
		fmt.Println(data,w)
	}
}
func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.Handle("/", handle(tmpl))
	http.ListenAndServe(":80", nil)
}
