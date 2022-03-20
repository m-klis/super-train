package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Radar"
	fmt.Println(Palindrome(str))
	str = "Malam"
	fmt.Println(Palindrome(str))
	str = "Kasur ini rusak"
	fmt.Println(Palindrome(str))
	str = "Ibu ratna antar ubi"
	fmt.Println(Palindrome(str))
	str = "Malas"
	fmt.Println(Palindrome(str))
	str = "Makan nasi goreng"
	fmt.Println(Palindrome(str))
	str = "Balonku ada lima"
	fmt.Println(Palindrome(str))
}

func Palindrome(str string) bool {
	var x bool
	str = strings.ToLower(str)

	for i := 0; i < len(str); i++ {
		if str[i] == str[len(str)-i-1] {
			x = true
		} else {
			x = false
			break
		}
	}

	return x
}
