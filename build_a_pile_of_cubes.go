package main

import (
	"fmt"
	"math"
)

func FindNb(m int) int {

	var sum float64
	var n int

	for n = 0; int(sum) < m; n++ {
		sum = math.Pow(float64(n*(n+1)), float64(2)) / 4
	}

	if int(sum) == m {
		return n - 1
	} else {
		return -1
	}

}

func main() {
	fmt.Println(FindNb(4183059834009))

}

// It("should handle basic cases", func() {
// 	dotest(4183059834009, 2022)
// 	dotest(24723578342962, -1)
// })
