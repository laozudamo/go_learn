package main

import "fmt"

func main() {
	f1 := add()
	a := f1(2)
	b := f1(10)
	fmt.Printf("f1: %v\n", a)
	fmt.Printf("f1: %v\n", b)
}

/*
	go语言闭包
	函数内部的函数
*/
func add() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x + y
	}
}
