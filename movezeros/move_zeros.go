package movezeros

func MoveZeros(arr []int) []int {
	nonZeros := []int{}
	zeros := []int{}

	for _, number := range arr {
		if number == 0 {
			zeros = append(zeros, number)
		} else {
			nonZeros = append(nonZeros, number)
		}
	}

	return append(nonZeros, zeros...)
}
