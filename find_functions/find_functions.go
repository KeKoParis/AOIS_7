package find_functions

const numWords, wordLen = 30, 8

func FindLess(arr [numWords][wordLen]int32, attribute [wordLen]int32) [wordLen]int32 {
	var g, l int32 = 0, 0
	var arrLess [][wordLen]int32

	for i := range arr {
		g, l = find(arr[i], attribute, g, l, wordLen-1)
		if g == 0 && l == 1 {
			arrLess = append(arrLess, arr[i])
		}
		g, l = 0, 0
	}
	//for i := range arrLess {
	//	fmt.Println(arrLess[i])
	//}

	if len(arrLess) == 0 {
		var a [wordLen]int32
		return a
	}

	currBiggest := arrLess[0]
	for i := range arrLess {
		g, l = 0, 0
		g, l = find(arrLess[i], currBiggest, g, l, wordLen-1)
		if g == 1 && l == 0 {
			currBiggest = arrLess[i]
		}
	}

	return currBiggest
}

func FindMore(arr [numWords][wordLen]int32, attribute [wordLen]int32) [wordLen]int32 {
	var g, l int32 = 0, 0
	var arrLess [][wordLen]int32

	for i := range arr {
		g, l = find(arr[i], attribute, g, l, wordLen-1)
		if g == 1 && l == 0 {
			arrLess = append(arrLess, arr[i])
		}
		g, l = 0, 0
	}
	//for i := range arrLess {
	//	fmt.Println(arrLess[i])
	//}

	if len(arrLess) == 0 {
		var a [wordLen]int32
		return a
	}

	currBiggest := arrLess[0]
	for i := range arrLess {
		g, l = 0, 0
		g, l = find(arrLess[i], currBiggest, g, l, wordLen-1)
		if g == 0 && l == 1 {
			currBiggest = arrLess[i]
		}
	}

	return currBiggest
}

func find(arr [wordLen]int32, attribute [wordLen]int32, g int32, l int32, j int) (int32, int32) {
	g1 := calcG(g, attribute[j], l, arr[j])
	l1 := calcL(g, attribute[j], l, arr[j])

	if j == 0 {
		return g1, l1
	} else {
		return find(arr, attribute, g1, l1, j-1)
	}
}

func calcG(g int32, a int32, l int32, s int32) int32 {
	if a == 0 && l == 0 && s == 1 {
		return 1
	} else {
		return g
	}
}

func calcL(g int32, a int32, l int32, s int32) int32 {
	if a == 1 && s == 0 && g == 0 {
		return 1
	} else {
		return l
	}
}
