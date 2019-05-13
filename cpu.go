/*
Go 1.5 版本之前，默认使用的是单核心执行。从 Go 1.5 版本开始，默认执行上面语句以便让代码并发执行，最大效率地利用 CPU。

GOMAXPROCS 同时也是一个环境变量，在应用程序启动前设置环境变量也可以起到相同的作用。
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 可以使用 runtime.NumCPU() 查询 CPU 数量
	fmt.Print(runtime.NumCPU())

	// 使用 runtime.GOMAXPROCS() 函数进行设置
	//runtime.GOMAXPROCS(逻辑CPU数量)
	// 这里的逻辑CPU数量可以有如下几种数值：
	// <1：不修改任何数值。
	// =1：单核心执行。
	// >1：多核并发执行
	runtime.GOMAXPROCS(runtime.NumCPU())
}
