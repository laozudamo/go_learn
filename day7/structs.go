package main

import "fmt"

func main() {
	test()
}

/*
	结构体的嵌套
*/
type Person struct {
	name string
	age  int
	dog  Dog
}

type Dog struct {
	name, color string
}

func test() {
	jack := Person{
		name: "jack",
		age:  18,
		dog: Dog{
			name:  "阿黄",
			color: "yellow",
		},
	}
	fmt.Printf("jack: %v\n", jack)
}
