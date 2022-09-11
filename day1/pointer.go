package main

import "fmt"

func main() {
	var value string = "hello world"
	var point *string = &value
	fmt.Println(point, *point)
}
