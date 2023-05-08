package main

import (
	"fmt"
	findC "lab7/closest_value"
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

	//for i := range arr {
	//	fmt.Println(i, ": ", arr[i])
	//}

	var find [8]int32
	find = arr[2]

	fmt.Println("Nearest-bottom")
	fmt.Println("Find ", find)

	fmt.Println("     ", findC.FindLess(arr, find))
}
