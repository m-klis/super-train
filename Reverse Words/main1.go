package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "cek Huruf Besar kecil"
	fmt.Println(reverseWord(data))
}

func IsUpper(s string) bool {
	if strings.ToLower(s) == s {
		return false
	}
	return true
}

func reverseWord(input string) string {
	var result string
	var output []string

	splitString := strings.Split(input, " ")

	for _, v := range splitString {
		result = ""
		for idx, w := range v {

			var a string
			a = string(w)

			if IsUpper(string(v[len(v)-idx-1])) {
				a = string(w)
			}

			if IsUpper(string(v[len(v)-idx-1])) {
				a = strings.ToUpper(a)
			} else {
				a = strings.ToLower(a)
			}

			result = a + result
		}
		output = append(output, result)
	}

	stringOutput := strings.Join(output, " ")
	return stringOutput
}
