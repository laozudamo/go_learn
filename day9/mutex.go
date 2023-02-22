package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		go add()
		go dec()
	}
	fmt.Printf("\"end\": %v\n", "end")

}

var i int = 100

func add() {
	i += 1
	fmt.Printf("i++: %v\n", i)
}

func dec() {
	i -= 1
	fmt.Printf("i--: %v\n", i)
}
