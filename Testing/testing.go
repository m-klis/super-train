package testing

import (
	"math"
	"strconv"
	"strings"
)

func FactorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

func Factorial(n int) int {
	x := 1
	for i := 1; i <= n; i++ {
		x = x * i
	}
	return x
}

func FindMiddleLetter(a, b string) string {
	if len(a) > 1 || len(b) > 1 {
		return "Single string input (\"A\",\"Z\")"
	}
	a, b = strings.ToUpper(a), strings.ToUpper(b)
	letter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	arrLetter := strings.Split(letter, "")
	var i_a, i_b int
	for i, l := range arrLetter {
		if a == l {
			i_a = i
		}
		if b == l {
			i_b = i
		}
	}
	x := math.Abs(float64(i_a)-float64(i_b)) / 2
	isBigger := 0
	if i_a > i_b {
		isBigger = i_a
	} else {
		isBigger = i_b
	}
	return arrLetter[isBigger-int(x)]
}

func FizzBuzz(n int) []string {
	x := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			x = append(x, "FizzBuzz")
		} else if i%5 == 0 {
			x = append(x, "Buzz")
		} else if i%3 == 0 {
			x = append(x, "Fizz")
		} else {
			str := strconv.Itoa(i)
			x = append(x, str)
		}
	}
	return x
}

func ReverseWord(str string) string {
	x := strings.Split(str, " ")
	y := make([]string, 0)
	for _, l := range x {
		rl := ReverseLetter(l)
		y = append(y, rl)
	}
	str = strings.Join(y, " ")
	return str
}

func ReverseLetter(str string) string {
	x := strings.Split(str, "")
	rw := make([]string, 0)
	for i := 0; i < len(x); i++ {
		isUpper := false
		if strings.ToUpper(x[i]) == x[i] {
			isUpper = true
		}

		if isUpper {
			cekUp := false
			if strings.ToUpper(x[len(x)-i-1]) == x[len(x)-i-1] {
				cekUp = true
			}
			if cekUp {
				rw = append(rw, x[len(x)-i-1])
			} else {
				rw = append(rw, strings.ToUpper(x[len(x)-i-1]))
			}
		} else {
			cekUp := false
			if strings.ToUpper(x[len(x)-i-1]) == x[len(x)-i-1] {
				cekUp = true
			}
			if cekUp {
				rw = append(rw, strings.ToLower(x[len(x)-i-1]))
			} else {
				rw = append(rw, x[len(x)-i-1])
			}
		}
	}

	str = strings.Join(rw, "")

	return str
}

func LeapYear(a, b int) []int {
	low := 0
	high := 0
	if a > b {
		high, low = a, b
	} else {
		high, low = b, a
	}
	arr := make([]int, 0)
	for i := low + 1; i <= high; i++ {
		if i%4 == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

func Palindrome(str string) string {
	x := strings.ToLower(str)
	y := false
	for i := range x {
		if x[i] == x[len(x)-i-1] {
			y = true
		} else {
			y = false
			break
		}
	}

	if y {
		str = "palindrome"
	} else {
		str = "not palindrome"
	}

	return str
}

func NearestFibonacci(arr []int) int {
	var x int = 0
	for i := 0; i < len(arr); i++ {
		x = x + arr[i]
	}

	var y int = 0
	z := x
	for {
		f := Fibonacci(y)

		if f > x*2 {
			break
		}

		cek := int(math.Abs(float64(f - x)))

		if cek < z {
			z = cek
		}

		y = y + 1
	}
	return z
}

func Fibonacci(n int) int {
	x := make([]int, 3)
	x[0] = 0
	x[1] = 1
	y := 0
	for i := 2; i <= n; i++ {
		if i == 2 {
			x[2] = x[0] + x[1]
		} else {
			y = x[i-1] + x[i-2]
			x = append(x, y)
		}
	}
	return x[n]
}
