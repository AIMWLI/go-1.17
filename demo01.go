package main

import (
	"fmt"
	"net/http"
)

//halo

//func main() {
//	fmt.Println("hello========")
//}

//web版halo
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "hello world")
	})
	fmt.Println("打开网页localhost:8888")
	// 打开端口
	http.ListenAndServe(":8888", nil)
}

//
//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprint(w, "Hello, World!")
//	})
//	fmt.Println("Please Visit -  http://localhost:8888/")
//	http.ListenAndServe(":8888", nil)
//}
