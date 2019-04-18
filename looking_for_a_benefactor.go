package main

import (
	"fmt"
	"math"
)

func NewAvg(arr []float64, navg float64) int64 {

	var arrSum float64

	for _, val := range arr {
		arrSum = arrSum + val
	}

	retVal := math.Round((float64(len(arr)+1) * navg) - arrSum)

	if retVal <= 0 {
		return -1
	}

	return int64(retVal)
}

func main() {

	//var dons = []float64{14, 30, 5, 7, 9, 11, 15}
	var dons = []float64{1400.25, 30000.76, 5.56, 7, 9, 11, 15.48, 120.98}
	fmt.Println(NewAvg(dons, 10000))

}

// https://www.codewars.com/kata/looking-for-a-benefactor

// The accounts of the "Fat to Fit Club (FFC)" association are supervised by John as a volunteered accountant. The association is funded through financial donations from generous benefactors. John has a list of the first n donations: [14, 30, 5, 7, 9, 11, 15] He wants to know how much the next benefactor should give to the association so that the average of the first n + 1 donations should reach an average of 30. After doing the math he found 149. He thinks that he made a mistake. Could you help him?
// if dons = [14, 30, 5, 7, 9, 11, 15] then new_avg(dons, 30) --> 149

// The function new_avg(arr, navg) should return the expected donation (rounded up to the next integer) that will permit to reach the average navg.

// Should the last donation be a non positive number (<= 0) John wants us to throw (or raise) an error
// (return Nothing in Haskell, return None in F# and Ocaml, return -1 in C, Fortran, Nim, PowerShell, Go; echo ERRORin Shell, raise-argument-error in Racket)
// so that he clearly sees that his expectations are not great enough.

// Notes:

//     all donations and navg are numbers (integers or floats), arr can be empty.
//     See examples below and "Test Samples" to see which error is to be thrown.
