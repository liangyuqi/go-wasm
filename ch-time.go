// +build ignore

/*
time.AfterFunc(time. fn) 函数是在 time.After 基础上增加了到时的回调

time.NewTicker(time) 打点器

time.NewTimer(time) 定时器
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// 声明一个退出用的通道
	// exit := make(chan int)
	// // 打印开始
	// fmt.Println("start")

	// nextSecond := time.After(time.Second)

	// // 过1秒后, 调用匿名函数
	// time.AfterFunc(time.Second, func() {
	// 	// 1秒后, 打印结果
	// 	fmt.Println("one second after")
	// 	// 通知main()的goroutine已经结束
	// 	exit <- 0
	// })
	// // 等待结束
	// <-exit

	// 创建一个打点器, 每500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)
	// 创建一个计时器, 2秒后触发
	stopper := time.NewTimer(time.Second * 2)
	// 声明计数变量
	var i int
	// 不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case <-stopper.C: // 计时器到时了
			fmt.Println("stop")
			// 跳出循环
			goto StopHere
		case <-ticker.C: // 打点器触发了
			// 记录触发了多少次
			i++
			fmt.Println("tick", i)
		}
	}

	// 退出的标签, 使用goto跳转
StopHere:
	fmt.Println("done")

}
