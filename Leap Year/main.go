package main

import "fmt"

func main() {
	fmt.Println(LeapYear(1900, 2020))
}

func LeapYear(a, b int) []int {
	var x []int

	for i := a + 1; i <= b; i++ {
		if i%4 == 0 {
			x = append(x, i)
		}
	}

	return x
}
