package main

import "fmt"

func main() {
	fmt.Println(Factorial(2))
}

func Factorial(n int) int {
	x := 1
	if n == 0 {
		return 0
	}
	for i := n; i >= 1; i-- {
		x = x * i
	}
	return x
}
