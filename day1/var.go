package main

import "fmt"

// 指定变量类型

func main() {
	var num1 int = 10

	// 不同变量相同类型的声明
	var num2, num3 int = 20, 8
	fmt.Println(num1, num2, num3)

	// 多个变量不同类型的声明
	var (
		a int    = 10
		b string = "hahha"
	)
	fmt.Println(a, b)

	// 采用类型推断
	x := "已知数"
	y := 999

	fmt.Println(x, y)
}

// 类型推断声明
// func main() {
// 	num1 := 1
// 	num2 := 2
// 	fmt.Println(num1, num2)
// }
