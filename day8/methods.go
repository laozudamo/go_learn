package main

import "fmt"

func main() {
	p := Person{
		name: "pppppp",
	}
	showPerson(p)
}

/*
	传递struct
*/

type Person struct {
	name string
	age  int
}

func showPerson(per Person) {
	per.name = "jack"
	fmt.Printf("per: %v\n", per)
}

func showPerson1(per *Person) {
	per.name = "luackF"
}
