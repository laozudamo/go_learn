package main

import "fmt"

func main() {
	var a, b, c = 1, false, "str"
	a = 2
	fmt.Println(a, b, c)

	x, y, z := 2, 3, "hahhah"
	fmt.Println(x, y, z)

	const PI float64 = 3.14
	const PI2 = 3.1415

	const (
		weight = 32
		height = 32
	)

	const i, j = 1, 2

}
