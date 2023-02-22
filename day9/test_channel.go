package main

import "fmt"

var c = make(chan int)

func main() {
	go test()
	for i := 0; i < 10; i++ {
		r := <-c
		fmt.Printf("r: %v\n", r)
	}

}

func test() {
	count := 10
	for i := 0; i < count; i++ {
		c <- i
	}
}
