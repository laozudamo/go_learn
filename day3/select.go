package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "111111"
	}()
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "333333"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("msg1: %v\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("msg2: %v\n", msg2)
		}
	}

}

func setArr() {
	list := [5]int{1, 2, 3, 4, 5}

	arr1 := [...]int{1, 2, 3}

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("list: %v\n", list)
}
