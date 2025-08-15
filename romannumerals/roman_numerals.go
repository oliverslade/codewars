package romannumerals

func RomanNumeral(number int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	// foreach value to symbol pair, add as many symbols as possible
	for i := 0; i < len(values); i++ {
		count := number / values[i]
		if count > 0 {
			for j := 0; j < count; j++ {
				result += symbols[i]
			}
			number -= count * values[i]
		}
	}

	return result
}
