package main

import "fmt"

func main() {

	/*
		go语言特性:
			1. 静态语言[强类型]
	*/

	// 先声明变量再赋值
	var num1 int
	num1 = 1

	// 类型推断
	num2 := 2

	// 类型声明并赋值
	var num3 int = 3

	fmt.Println("num1的值是:", num1)
	fmt.Println("num1的值是:", num2)
	fmt.Println("num1的值是:", num3)

	// 同时命名多个变量
	var a, b, c, d int
	a = 1
	b = 2
	c = 3
	d = 4

	fmt.Println(a, b, c, d)

	f, g := 0, 1

	fmt.Println(f, g)

	var m, n int = 0, 0

	fmt.Println(m, n)

}
