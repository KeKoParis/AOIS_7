package find_functions

func FindInterval(arr [numWords][wordLen]int32, top [wordLen]int32, bottom [wordLen]int32) [][wordLen]int32 {
	var g, l int32 = 0, 0
	var interval [][wordLen]int32
	var outOfInterval [numWords]bool

	for i := range arr {
		g, l = find(arr[i], top, g, l, wordLen-1)
		if g == 1 && l == 0 {
			outOfInterval[i] = true
		}
		g, l = 0, 0
	}

	for i := range arr {
		g, l = find(arr[i], bottom, g, l, wordLen-1)
		if g == 0 && l == 1 {
			outOfInterval[i] = true
		}
		g, l = 0, 0
	}

	for i := range arr {
		if outOfInterval[i] == false {
			interval = append(interval, arr[i])
		}
	}
	return interval
}
