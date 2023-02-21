package main

import "fmt"

func main() {

	// test()
	test2()
}

/*
	相同类型可以定义到一行
*/
type Person struct {
	id, sex string
	age     int
}

func test() {
	var tom Person
	tom.id = "12"
	tom.sex = "男"
	tom.age = 18

}

func test2() {

	// var jack Person
	jack := Person{
		id:  "20",
		sex: "女",
		age: 20,
	}
	lucy := Person{
		"20",
		"nan",
		19,
	}
	fmt.Printf("jack: %v\n", jack)
	fmt.Printf("lucy: %v\n", lucy)
}

/*
	匿名结构体
*/

/*
	结果体初始化
*/
