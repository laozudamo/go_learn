package main

import "fmt"

func main() {
	type myInt = int
	list := []myInt{1, 2, 3, 4}
	fmt.Printf("list: %v\n", list)
}

/*
	type newtype = Type
*/
