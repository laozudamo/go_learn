package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		go showMag(i)

		//多一个携程
		wp.Add(1)
	}

	//主线程 这里需要等待 执行完
	wp.Wait()

	fmt.Printf("\"end\": %v\n", "end")
}

/*
	启动多个写成
*/

func showMag(i int) {
	fmt.Printf("i: %v\n", i)
	//执行一次 减少一个
	defer wp.Done()
}
