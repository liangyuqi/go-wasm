// +build ignore

/*
只能发送的通道类型为chan<-，只能接收的通道类型为<-chan
单向通道有利于代码接口的严谨性
*/

package main

func main() {

	ch := make(chan int)
	// 声明一个只能发送的通道类型, 并赋值为ch
	var chSendOnly chan<- int = ch

	//声明一个只能接收的通道类型, 并赋值为ch
	var chReadOnly <-chan int = ch

	// 当然，使用 make 创建通道时，也可以创建一个只发送或只读取的通道：
	// ch := make(<-chan int)
	// var chReadOnly <-chan int = ch
	// <-chReadOnly
}
