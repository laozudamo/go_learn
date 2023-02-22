package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// load()
	ces()
}

/*
	载入
*/

func load() {
	var i int32 = 100
	// 写
	atomic.LoadInt32(&i)

	fmt.Printf("i: %v\n", i)

	// 存储
	atomic.StoreInt32(&i, 200)

	fmt.Printf("i: %v\n", i)
}

/*
	交换
*/

func ces() {
	var i int32 = 64

	// 返回布尔值
	b := atomic.CompareAndSwapInt32(&i, 64, 300)
	fmt.Printf("b: %v\n", b)
}
