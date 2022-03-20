package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{15, 3, 1}
	fmt.Println(NearestFibonacci(arr))
}

func NearestFibonacci(arr []int) int {
	x := 0
	for _, l := range arr {
		x = x + l
	}

	y := 0
	z := x
	for {
		f := Fibonacci(y)

		if f > x*2 {
			break
		}

		check := math.Abs(float64(f - x))

		if z > int(check) {
			z = int(check)
		}

		y = y + 1
	}

	return z
}

func Fibonacci(n int) int {
	ar := make([]int, 3)
	ar[0] = 0
	ar[1] = 1

	for i := 2; i <= n; i++ {
		if i == 2 {
			ar[2] = ar[0] + ar[1]
		} else {
			x := ar[i-1] + ar[i-2]
			ar = append(ar, x)
		}
	}

	return ar[n]
}
