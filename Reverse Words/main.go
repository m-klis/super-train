package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Great"
	fmt.Println(ReverseWord(word))
	sentence := "I am A Great human"
	fmt.Println(ReverseSentence(sentence))
}

// Big 65 - 90, A - Z
// Small 97 - 122, a - z
func ReverseSentence(w string) string {
	x := strings.Split(w, " ")
	var y []string

	for i := range x {
		save := ReverseWord(x[i])
		y = append(y, save)
	}

	w = strings.Join(y, " ")
	return w
}

func ReverseWord(str string) string {
	var rw []string
	var x, y, z bool
	for i, l := range str {
		x = l >= 65 && l <= 90  // besar
		y = l >= 97 && l <= 122 // kecil
		if x {
			z = str[len(str)-i-1] >= 65 && str[len(str)-i-1] <= 90
			if z {
				rw = append(rw, string(str[len(str)-i-1]))
			} else {
				rw = append(rw, string(str[len(str)-i-1]-32))
			}
		}
		if y {
			z = str[len(str)-i-1] >= 65 && str[len(str)-i-1] <= 90
			if z {
				rw = append(rw, string(str[len(str)-i-1]+32))
			} else {
				rw = append(rw, string(str[len(str)-i-1]))

			}
		}
	}
	str = strings.Join(rw, "")
	return str
}
