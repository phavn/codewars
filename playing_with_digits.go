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
