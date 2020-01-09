// Go1.11开始支持WebAssembly，
// 对应的操作系统名为js，对应的CPU类型为wasm。
// 目前还无法通过 go run 的方式直接运行输出的wasm文件，
// 因此我们需要通过 go build 的方式生成wasm目标文件，
// cd go-webassembly/hello
// GOARCH=wasm GOOS=js go build -o hello.wasm hello.go
// 然后通过Node环境执行。
// node ../lib/wasm_exec.js hello.wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello WebAssembly!")
	js.Global().Get("console").Call("log", "Hello World Go/wasm!")
	// js.Global().Get("document").Call("getElementById", "app").Set("innerText", time.Now().String())
}
