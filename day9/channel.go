package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("value: %v\n", value)
}

/*
	通道
	通道用于携程之间的通信
	g1
	g2
	两个要相互发消息

	chan 类型的空值是 nil，声明后需要配合 make 后才能使用

	Unbuffered := make (chan int） //整型无緩冲通道
	buffered := make (chan int, 10) //整型有緩沖通道
*/

var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}
