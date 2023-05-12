package main

import (
	"fmt"
	findF "lab7/find_functions"
	"math/rand"
)

func wordsArray() [30][8]int32 {
	var words [30][8]int32

	for i := range words {
		for j := range words[i] {
			words[i][j] = rand.Int31n(2)
		}
	}

	return words
}

func main() {
	rand.Seed(12)
	// rand.Seed(time.Now().UnixNano())

	arr := wordsArray()

	for i := range arr {
		fmt.Println(i, ": ", arr[i])
	}

	var find, top, bottom [8]int32
	find = arr[2]
	top = arr[5]
	bottom = arr[2]

	fmt.Println("Nearest-bottom")
	fmt.Println("Find ", find)

	fmt.Println("Less ", findF.FindLess(arr, find))
	fmt.Println("More ", findF.FindMore(arr, find))

	fmt.Println("\nInterval\nTop ", top, "  Bottom ", bottom)
	fmt.Println("Interval")
	interval := findF.FindInterval(arr, top, bottom)
	for i := range interval {
		fmt.Println(interval[i])
	}
}
