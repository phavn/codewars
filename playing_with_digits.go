package main

import (
	"fmt"
	"math"
	"strconv"
)

func DigPow(n, p int) int {

	var total float64

	for _, num := range strconv.Itoa(n) {
		base, err := strconv.Atoi(string(num))
		if err != nil {
			fmt.Println(err)
		}

		total = total + math.Pow(float64(base), float64(p))
		p++
	}

	retValue := int(total) / n
	remainder := int(total) % n

	if remainder == 0 {
		return retValue
	} else {
		return -1
	}

}

func main() {
	fmt.Println(DigPow(3456789, 5))
}

// https://www.codewars.com/kata/playing-with-digits/go

// Some numbers have funny properties. For example:
//     89 --> 8¹ + 9² = 89 * 1
//     695 --> 6² + 9³ + 5⁴= 1390 = 695 * 2
//     46288 --> 4³ + 6⁴+ 2⁵ + 8⁶ + 8⁷ = 2360688 = 46288 * 51

// Given a positive integer n written as abcd... (a, b, c, d... being digits) and a positive integer p
//     we want to find a positive integer k, if it exists, such as the sum of the digits of n taken to the successive powers of p is equal to k * n.

// In other words:
//     Is there an integer k such as : (a ^ p + b ^ (p+1) + c ^(p+2) + d ^ (p+3) + ...) = n * k

// If it is the case we will return k, if not return -1.
// Note: n and p will always be given as strictly positive integers.
