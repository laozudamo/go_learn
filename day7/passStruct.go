package main

import "fmt"

type Person struct {
	age, id int
	name    string
}

func main() {
	person := Person{19, 22, "jack"}
	fmt.Printf("person: %v\n", person)
	test(person)
}

/*
	go结构体可以像昔通变量一样，作为函数的参数，传递给函数，这里分为两种情况：
1，直接传递结构体，这是是一个副本（持贝），在函数内部不会政变外面结构体内容。
2，传递结构体指针，这时在函数内部，能够改变外部结构体内容。
*/

// 拷贝了副本 进行值传递
func test(person Person) {
	person.id = 2222
	// person.age = 20
	// person.name = "jack"
	fmt.Printf("person: %v\n", person)
}
