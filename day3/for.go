package main

import "fmt"

func main() {

	// getAge()
	// for i := 1; i < 20; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }

	getNum()
}

// 初始条件写到外面

func getAge() {
	age := 18
	for ; age < 20; age++ {
		fmt.Printf("age: %v\n", age)
	}
}

func getNum() {
	age := 3
	for age < 19 {
		fmt.Printf("玩摇摇马%v\n", age)
		age++
	}
}
