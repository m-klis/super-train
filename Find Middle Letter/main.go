package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FindMiddleLetter("Q", "U"))
}

func FindMiddleLetter(a, b string) string {
	letter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var w, x, y, z int
	for i := range letter {
		if a == string(letter[i]) {
			x = i
		}
		if b == string(letter[i]) {
			y = i
		}
	}

	if x > y {
		w = x
	} else {
		w = y
	}

	z = int(math.Abs(float64(x-y))) / 2

	return string(letter[w-z])
}
