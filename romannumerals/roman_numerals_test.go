package romannumerals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumerals(t *testing.T) {
	t.Run("single digit cases", func(t *testing.T) {
		assert.Equal(t, "I", RomanNumeral(1))
		assert.Equal(t, "V", RomanNumeral(5))
		assert.Equal(t, "X", RomanNumeral(10))
	})

	t.Run("codewars examples", func(t *testing.T) {
		assert.Equal(t, "CLXXXII", RomanNumeral(182))
		assert.Equal(t, "MCMXC", RomanNumeral(1990))
		assert.Equal(t, "MDCCCLXXV", RomanNumeral(1875))
	})

	t.Run("complex cases", func(t *testing.T) {
		assert.Equal(t, "M", RomanNumeral(1000))
		assert.Equal(t, "MDCLXVI", RomanNumeral(1666))
		assert.Equal(t, "MMVIII", RomanNumeral(2008))
	})

	t.Run("subtractive cases", func(t *testing.T) {
		assert.Equal(t, "IV", RomanNumeral(4))
		assert.Equal(t, "IX", RomanNumeral(9))
		assert.Equal(t, "XL", RomanNumeral(40))
		assert.Equal(t, "XC", RomanNumeral(90))
		assert.Equal(t, "CD", RomanNumeral(400))
		assert.Equal(t, "CM", RomanNumeral(900))
	})

	t.Run("edge cases", func(t *testing.T) {
		assert.Equal(t, "I", RomanNumeral(1))
		assert.Equal(t, "MMMCMXCIX", RomanNumeral(3999))
	})
}
