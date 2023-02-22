package main

import (
	"fmt"
	"runtime"
)

/*
	runtime.Gosched()
*/
func main() {
	go showMsg("java")

	runtime.GOMAXPROCS(1)
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	for i := 0; i < 2; i++ {

		// 让出资源 等待
		runtime.Gosched()
		fmt.Printf("\"goalng\": %v\n", "goalng")
	}
}

/*
	Goexit 推出
*/
func showMsg(s string) {
	count := 5
	for i := 0; i < count; i++ {
		if i >= 3 {
			runtime.Goexit()
		}
		fmt.Printf("i: %v\n", s)
	}
}
