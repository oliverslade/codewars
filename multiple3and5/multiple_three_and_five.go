package multiple3and5

func Multiple3And5(number int) int {
	result := 0

	for i := number - 1; i > 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			result = result + i
		}
	}

	return result
}
