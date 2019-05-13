// +build ignore

/*
Go 语言中的通道（channel）是一种特殊的类型。在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信。

通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。

① 通道的收发操作在不同的两个 goroutine 间进行。

由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。

② 接收将持续阻塞直到发送方发送数据。

如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。

③ 每次接收一个元素。
通道一次只能接收一个数据元素。
*/

package main

import (
	"fmt"
)

func main() {

	// 通道实例 := make(chan 通道类型, 缓冲大小)
	ch := make(chan int)

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine,使用<-向通道发送数据
		ch <- 0
		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")
	// 等待匿名goroutine
	data := <-ch

	// 这一句会一直阻塞直到 main 的 goroutine 接收为止
	fmt.Println("all done", data)

	// 循环接收
	// ch := make(chan int)
	// go func() {
	// 	for i := 3; i >= 0; i-- {
	// 		ch <- i
	// 		time.Sleep(time.Second)
	// 	}
	// }()
	// for data := range ch {
	// 	fmt.Println(data)
	// 	if data == 0 {
	// 		break
	// 	}
	// }
}
