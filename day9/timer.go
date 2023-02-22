package main

import (
	"fmt"
	"time"
)

/*
	NewTimer timer的实栗化
*/
func main() {
	// test2()
	// test3()
	// test4()
}

/*
	单纯想等待 time.Sleep()
*/

/*
	time.After()
*/

func test() {
	timer := time.NewTimer(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())

	/*
		阻塞2s
	*/
	t1 := <-timer.C
	// 也可以
	// <-timer.C
	fmt.Printf("t1: %v\n", t1)
}

/*
	<-time.After(time.Second * 2)进行延迟处理
*/
func test2() {
	fmt.Printf("time.Now(): %v\n", time.Now())
	<-time.After(time.Second * 2)
	fmt.Printf("time.Now(): %v\n", time.Now())
}

/*
	<-time.After(time.Second * 2)进行延迟处理
*/
func test3() {
	fmt.Printf("time.Now(): %v\n", time.Now())
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Printf("time.Now(): %v\n", time.Now())
}

/*
	timr.Reset()
*/

/*
	timr.Stop()
*/

/*
	timer.Ticker()
*/

/*
	ticker.Stop()
*/
func test4() {
	ticker := time.NewTicker(time.Second)
	counter := 1
	for _ = range ticker.C {
		counter += 1
		fmt.Printf("\"ticker\": %v\n", "ticker")
		if counter >= 10 {
			ticker.Stop()
			break
		}
	}
}
