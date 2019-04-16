package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := "aeiou"

	for _, c := range str {
		if strings.Contains(vowels, string(c)) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(GetCount("abracadabrao"))
}

// Return the number (count) of vowels in the given string.
// We will consider a, e, i, o, and u as vowels for this Kata.
// The input string will only consist of lower case letters and/or spaces.
