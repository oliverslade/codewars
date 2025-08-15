package findodd

func FindOdd(seq []int) int {
	countMap := make(map[int]int)

	for i := 0; i < len(seq); i++ {
		countMap[seq[i]]++
	}

	for number, count := range countMap {
		// Crucially there will always be only one integer that appears an odd number of times
		if count%2 != 0 {
			return number
		}
	}

	panic("should never reach this point")
}
