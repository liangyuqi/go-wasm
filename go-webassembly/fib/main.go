/*
运行静态文件服务器，这里不能使用普通的静态文件服务器，因为浏览器要求请求到的 WebAssemly 字节码文件的 Content-Type 必须是 application/wasm，很多静态文件服务器并不会因为扩展名是 wasm 就会自动使用这个 Content-Type。但是 Go 内置的 HTTP 服务器可以。
go run main.go
*/
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8000", mux))
}
