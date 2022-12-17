package main

import "fmt"

func main() {
	a, b := 10, 20
	resMax := max(a, b)
	fmt.Println("resMax", resMax)

	m, n := swap("Google", "Runoob")
	fmt.Println(m, n)

	var grades = [5]int{1, 3, 5, 6, 20}

	theGrade := getAverageGrade(grades)

	fmt.Println("theGrade", theGrade)
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

// 传递数组
func getAverageGrade(arr [5]int) int {
	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	// fmt.Println("sum", sum)
	sumAvarageGrade := sum / len(arr)
	return sumAvarageGrade
}
