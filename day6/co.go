package main

import "fmt"

func main() {
	test(3)

	// test1()

	test2(1)
}

/*
	go语言递归
*/

func test(num int) {
	num++
	if num < 10 {
		fmt.Printf("num: %v\n", num)
		test(num)
	}
}

/*
	阶层
*/

func test1() {
	s := 1
	for i := 1; i <= 5; i++ {
		s = s * i
	}
	fmt.Printf("s: %v\n", s)
}
