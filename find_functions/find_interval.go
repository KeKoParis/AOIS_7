package find_functions

func FindInterval(arr [numWords][wordLen]int32, top [wordLen]int32, bottom [wordLen]int32) [][wordLen]int32 {
	var interval [][wordLen]int32
	var outOfInterval [numWords]bool

	outOfInterval = findOutOfInterval(arr, top, bottom)

	for i := range arr {
		if outOfInterval[i] == false {
			interval = append(interval, arr[i])
		}
	}
	return interval
}

func findOutOfInterval(arr [numWords][wordLen]int32, top [wordLen]int32, bottom [wordLen]int32) [numWords]bool {
	var gTop, lTop, gBottom, lBottom int32 = 0, 0, 0, 0
	var outOfInterval [numWords]bool

	for i := range arr {
		gBottom, lBottom = find(arr[i], top, gTop, lTop, wordLen-1)
		gTop, lTop = find(arr[i], bottom, gTop, lTop, wordLen-1)
		if gTop == 0 && lTop == 1 {
			outOfInterval[i] = true
		}
		if gBottom == 1 && lBottom == 0 {
			outOfInterval[i] = true
		}
		gTop, lTop, gBottom, lBottom = 0, 0, 0, 0
	}
	return outOfInterval
}
