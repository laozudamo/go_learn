package main

import "fmt"

func main() {
	isWorkDay()

	isfunDay()
}

func isWorkDay() {
	day := 4
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Printf("\"工作日期\": %v\n", "工作日期")

	case 6, 7:
		fmt.Printf("\"周末\": %v\n", "周末")
	}
}
func isfunDay() {
	day := 1
	switch {
	case day == 1:
		fmt.Printf("\"摇摇马\": %v\n", "摇摇马")
	case day == 2:
		fmt.Printf("\"2222\": %v\n", "2222")
	case day == 3, day == 4, day == 5, day == 6:
		fmt.Printf("\"工作日期\": %v\n", "工作日期")
	}
}
