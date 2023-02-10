package main

import "fmt"

func main() {
	f1("jack", sayHello)

	ff := calc("+")
	v := ff(4, 5)
	fmt.Printf("v: %v\n", v)
}

func sayHello(name string) {

	fmt.Printf("\"hello\": %v\n", name)
}

func f1(name string, fn func(string)) {
	fn(name)
}

func add(x int, y int) int {
	return x + y
}

func dec(x int, y int) int {
	return x - y
}

func calc(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return dec
	default:
		return nil
	}
}
