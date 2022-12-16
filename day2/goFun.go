package main

import "fmt"

func main() {
	a, b := 10, 20
	resMax := max(a, b)
	fmt.Println("resMax", resMax)

	m, n := swap("Google", "Runoob")
	fmt.Println(m, n)
}

func max(num1, num2 int) int {
	var res int
	if num1 > num2 {
		res = num2
	} else {
		res = num1
	}
	return res
}

// 返回多个函数值

func swap(x, y string) (string, string) {
	return x, y
}
