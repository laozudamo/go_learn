package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string)
var intChan = make(chan int)

/*
select是Go中的一个控制结构，类似于 switch 语向，用于处理异步10操作。
select 会监听case语向中channe的读写操作，当case中channel读写操作为非阻塞状态（即能读亏）时，将会触发相应的动作。
	2. 如果有多个case 都可以运行，select 会随机公平地选出一个执行，其他不会执行。
3，如果没有可运行的c日e语句，且有default语句，那么就会执行default的动作。
4. 如果没有可运行的case 语句，且没有default 语句， select 将阳塞，直到某个 case 通信可以运行
*/

func main() {
	go test()

	for {
		select {
		case r := <-strChan:
			fmt.Printf("r: %v\n", r)
		case r := <-intChan:
			fmt.Printf("r: %v\n", r)
		default:
			fmt.Printf("\"没有值\": %v\n", "没有值")
		}
		time.Sleep(time.Second)
	}

}

func test() {
	intChan <- 100
	strChan <- "hello world"
	defer close(intChan)
	defer close(strChan)
}
