package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {

	var returnString string
	catValues := make([]int, len(listCat))

	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	for catIndex, cat := range listCat {
		for _, art := range listArt {
			if art[0:1] == cat {
				tmpVal, err := strconv.Atoi(strings.Fields(art)[1])
				if err != nil {
					fmt.Println(err)
				}
				catValues[catIndex] = catValues[catIndex] + tmpVal
			}
		}
	}

	for catIndex := range listCat {
		returnString = returnString + "(" + listCat[catIndex] + " : " + strconv.Itoa(catValues[catIndex]) + ")"

		if catIndex < len(listCat)-1 {
			returnString = returnString + " - "
		}
	}

	return returnString
}

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}

	fmt.Println(StockList(b, c))
}
