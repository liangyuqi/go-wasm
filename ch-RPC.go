// +build ignore

/*
在使用通道时，想同时接收多个通道的数据是一件困难的事情。通道在接收数据时，如果没有数据可以接收将会发生阻塞。虽然可以使用如下模式进行遍历，但运行性能会非常差。
for{
    尝试接收ch1通道
    data, ok := <-ch1
    尝试接收ch2通道
    data, ok := <-ch2
    接收后续通道
    …
}

Remote Procedure Call，远程过程调用

RPC 能有效地封装通信过程，让远程的数据收发通信过程看起来就像本地的函数调用一样。
*/

package main

import (
	"errors"
	"fmt"
	"time"
)

// 模拟RPC客户端的请求和接收消息封装
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req
	// 等待服务器返回
	select {
	case ack := <-ch: // 接收到服务器返回数据
		return ack, nil
	case nowTime := <-time.After(time.Second): // 超时，time.After 返回一个通道，这个通道在指定时间后，通过通道返回当前时间。
		return nowTime.String(), errors.New("Time out")
	}
}

// 模拟RPC服务器端接收客户端请求和回应
func RPCServer(ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch
		// 打印接收到的数据
		fmt.Println("server received:", data)

		// 模拟超时
		time.Sleep(time.Second * 5)
		// 反馈给客户端收到
		ch <- "tony"
	}
}
func main() {
	// 创建一个无缓冲字符串通道
	ch := make(chan string)
	// 并发执行服务器逻辑
	go RPCServer(ch)
	// 客户端请求数据和接收数据
	recv, err := RPCClient(ch, "hi")
	if err != nil {
		// 发生错误打印
		fmt.Println(err)
	} else {
		// 正常接收到数据
		fmt.Println("client received", recv)
	}
}
