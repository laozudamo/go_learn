package main

import "fmt"

func main() {
	var a int = 20

	if (a < 10) {
		fmt.Println("a小于10")
	} else if (a <= 15) {
		fmt.Println("a < 15")
	} else {
		fmt.Println("a > 20")
	}
}