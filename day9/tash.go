package main

import (
	"fmt"
	"time"
)

/*
	golang 中的并发是西数相互独立运行的能力。Goroutines 是并发运行的函数。Golang 提供了 Goroutines 作为并发处理操作的一种方式。

	go task ()
*/
func main() {
	go showMsg("java") // 启动携程
	go showMsg("golang")
	fmt.Println("main end...") /// 主函数退出 程序就退出
	time.Sleep(time.Microsecond * 2000)
}

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Microsecond * 100)
	}
}
