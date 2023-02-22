package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//100 万你-10你老婆90-10
// var a int = 100

// var lock sync.Mutex

// func main() {
// 	for i := 0; i < 100; i++ {
// 		go add()
// 		go dec()
// 	}
// 	time.Sleep(time.Second * 2)
// 	fmt.Printf("i: %v\n", a)

// }

// func add() {
// 	lock.Lock()
// 	a++
// 	lock.Unlock()
// }

// func dec() {
// 	lock.Lock()
// 	a--
// 	lock.Unlock()
// }

var i int32 = 100

func main() {
	for i := 0; i < 100; i++ {
		go add()
		go dec()
	}
	time.Sleep(time.Second * 2)
	fmt.Printf("i: %v\n", i)
}

func add() {
	atomic.AddInt32(&i, 1)
}

func dec() {
	atomic.AddInt32(&i, -1)
}
