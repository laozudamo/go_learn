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

	jack := "已知数"

	fmt.Printf("jack: %v\n", jack)

	fmt.Println(x, y)

	// 下划线 这个变量并没有用上
	_, age := getInfo()

	// fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}

// 类型推断声明
// 短变量变量声明 只能用于函数内部

// func main() {
// 	num1 := 1
// 	num2 := 2
// 	fmt.Println(num1, num2)
// }

// 批量初始化
func initVars() {
	var age, sex, height = 18, "男", 188
	fmt.Printf("age: %v\n", age)
	fmt.Printf("sex: %v\n", sex)
	fmt.Printf("height: %v\n", height)
}

// 匿名变量

func getInfo() (name string, age int) {
	return "tom", 20
}
