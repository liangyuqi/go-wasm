// +build ignore

/*
所有 goroutine 在 main() 函数结束时会一同结束。

goroutine 虽然类似于线程概念，但是从调度性能上没有线程细致，而细致程度取决于 Go 程序的 goroutine 调度器的实现和运行环境。

终止 goroutine 的最好方法就是自然返回 goroutine 对应的函数。虽然可以用 golang.org/x/net/context 包进行 goroutine 生命期深度控制，但这种方法仍然处于内部试验阶段，并不是官方推荐的特性。
*/

package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
func main() {
	// 并发执行程序
	go running()
	// 使用匿名函数创建goroutine
	// go func() {
	//     var times int
	//     for {
	//         times++
	//         fmt.Println("tick", times)
	//         time.Sleep(time.Second)
	//     }
	// }()

	// 接受命令行输入, 不做任何事情，Scanln回车直接结束
	var input string
	fmt.Scanln(&input)
}
