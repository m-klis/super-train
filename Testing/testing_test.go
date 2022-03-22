package testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMiddleLetter(t *testing.T) {
	fml := FindMiddleLetter("Q", "U")
	fmt.Println(fml)
	assert.Equal(t, "S", fml)
	assert.NotEqual(t, "N", fml)
}

func TestFizzBuzz(t *testing.T) {
	fb := FizzBuzz(3)
	fmt.Println(fb)
	testCase := []string{"1", "2", "Fizz"}
	assert.Equal(t, testCase, fb)
	fb = FizzBuzz(5)
	fmt.Println(fb)
	testCase = []string{"1", "2", "Fizz", "4", "Buzz"}
	assert.Equal(t, testCase, fb)
	fb = FizzBuzz(15)
	fmt.Println(fb)
	testCase = []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	assert.Equal(t, testCase, fb)
}

func TestReverseWord(t *testing.T) {
	rw := ReverseWord("I am A Great human")
	fmt.Println(rw)
	assert.Equal(t, "I ma A Taerg namuh", rw)
}

func TestReverseLetter(t *testing.T) {
	rl := ReverseLetter("Great")
	fmt.Println(rl)
	assert.Equal(t, "Taerg", rl)
}

func TestLeapYear(t *testing.T) {
	ly := LeapYear(1900, 2020)
	arr := []int{1904, 1908, 1912, 1916, 1920, 1924, 1928, 1932, 1936, 1940, 1944, 1948, 1952, 1956, 1960, 1964, 1968, 1972, 1976, 1980, 1984, 1988, 1992, 1996, 2000, 2004, 2008, 2012, 2016, 2020}
	assert.Equal(t, arr, ly)
}

func TestNotPalindrome(t *testing.T) {
	p := Palindrome("Malas")
	assert.Equal(t, "not palindrome", p)
	assert.NotEqual(t, "palindrome", p)
}

func TestPalindrome(t *testing.T) {
	p := Palindrome("Radar")
	assert.Equal(t, "palindrome", p)
	assert.NotEqual(t, "not palindrome", p)
}

func TestFibonacci(t *testing.T) {
	f := Fibonacci(2)
	assert.Equal(t, 1, f)
}

func TestNearestFibonacci(t *testing.T) {
	f := NearestFibonacci([]int{15, 3, 1})
	assert.Equal(t, 2, f)
}
