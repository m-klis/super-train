package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FizzBuzz(3))
	fmt.Println(FizzBuzz(5))
	fmt.Println(FizzBuzz(15))
}

func FizzBuzz(n int) []string {
	var arr []string
	var x string
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			x = strconv.Itoa(i)
			arr = append(arr, x)
		}
	}
	return arr
}
