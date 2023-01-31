package main

import (
	"fmt"
)

func main() {
	age := 10
	name := "jack"

	ad := &age

	v := *ad

	fmt.Printf("age: %T\n", age)
	fmt.Printf("name: %T\n", name)
	// fmt.Printf("v: %T\n", v)
	fmt.Printf("v: %v\n", v)
}
