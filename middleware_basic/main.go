package main

import (
	"fmt"
	"log"
	"net/http"
)
//中间件：通常会接受一个请求，对其进行处理，每个中间件只处理一件事情，完成后将其传递给另一个中间件或最终处理程序，做到程序的解耦。
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		fmt.Println(r)
	}
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}
func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))
	http.ListenAndServe(":8080", nil)
}
