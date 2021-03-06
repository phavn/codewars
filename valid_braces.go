package main

import (
	"fmt"
)

func ValidBraces(str string) bool {

	var strSlice []string

	for _, c := range(str) {
		strSlice = append(strSlice, string(c))
	}

	i := 0
	l := len(strSlice)

	// If matching braces found, remove them both from slice and decrease slice length by 2
	for i < l {
		switch string(strSlice[i]) {
		case ")":
			if i > 0 && strSlice[i-1] == "(" {
				strSlice = append(strSlice[:i-1], strSlice[i+1:]...)
				l = l - 2
				i--
			} else {
				i++
			}
		case "]":
			if i > 0 && strSlice[i-1] == "[" {
				strSlice = append(strSlice[:i-1], strSlice[i+1:]...)
				l = l - 2
				i--
			} else {
				i++
			}
		case "}":
			if i > 0 && strSlice[i-1] == "{" {
				strSlice = append(strSlice[:i-1], strSlice[i+1:]...)
				l = l - 2
				i--
			} else {
				i++
			}
		default:
			i++
		}
	}

	// Return fqlse if non empty slice
	return len(strSlice) == 0

}

func main() {

	//fmt.Println(ValidBraces("({})"))
	//fmt.Println(ValidBraces("(){}[]"))
	fmt.Println(ValidBraces("]}"))
}



// https://www.codewars.com/kata/5277c8a221e209d3f6000b56/train/go

// Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

// This Kata is similar to the Valid Parentheses Kata, but introduces new characters: brackets [], and curly braces {}. Thanks to @arnedag for the idea!

// All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

// What is considered Valid?
// A string of braces is considered valid if all braces are matched with the correct brace.

// Examples
//    =>  True
// "([{}])"   =>  True
// "(}"       =>  False
// "[(])"     =>  False
// "[({})](]" =>  False