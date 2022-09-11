package main

import "fmt"

func main() {
	var str string = "bearer"
	var name = getName(str)
	fmt.Println(name)
}

func getName(str string) string {
	switch str {
	case "aplle":
		return "苹果"
	case "bearer":
		return "啤酒"
	default:
		return "不知道"
	}
}