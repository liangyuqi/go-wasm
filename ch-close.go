// +build ignore

/*
如果尝试对已经关闭的通道进行发送，将会触发宕机。
缓冲通道在关闭后依然可以访问内部的数据。
*/

package main

import (
	"fmt"
)

func main() {
	// 创建一个整型的通道
	ch := make(chan int)
	// 关闭通道
	close(ch)
	// 打印通道的指针, 容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))
	// ! 给关闭的通道发送数据,将会panic
	ch <- 1

	// 创建一个整型带两个缓冲的通道
	ch2 := make(chan int, 2)

	// 给通道放入两个数据
	ch2 <- 0
	ch2 <- 1

	// 关闭缓冲
	close(ch2)
	// 遍历缓冲所有数据, 且多遍历1个
	for i := 0; i < cap(ch2)+1; i++ {

		// 从通道中取出数据
		v, ok := <-ch2

		// 打印取出数据的状态
		fmt.Println(v, ok)
	}

}
