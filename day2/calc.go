package main

import "fmt"

func main() {
	a := 1

	a += a

	c := 2
	c *= a

	d := &c

	e := *d

	fmt.Println("aa", a, c, d, e)
}
