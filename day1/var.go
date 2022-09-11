package main

import "fmt"

func main() {
	fmt.Println("你好,你是个狗狗")

	// 打包声明多个变量
	var (
		age = 24
		sex = "男"
	)

	fmt.Println(age, sex)

}
