package multiple3and5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiple3And5(t *testing.T) {
	t.Run("simple calculation", func(t *testing.T) {
		assert.Equal(t, 23, Multiple3And5(10))
	})

	t.Run("number below 3 should return 0", func(t *testing.T) {
		assert.Equal(t, 0, Multiple3And5(2))
	})

	t.Run("number 3 should return 0 (below 3)", func(t *testing.T) {
		assert.Equal(t, 0, Multiple3And5(3))
	})

	t.Run("larger number 20", func(t *testing.T) {
		assert.Equal(t, 78, Multiple3And5(20))
	})

	t.Run("number 15 includes multiple of both 3 and 5", func(t *testing.T) {
		assert.Equal(t, 45, Multiple3And5(15))
	})

	t.Run("number 16 includes 15 which is multiple of both", func(t *testing.T) {
		assert.Equal(t, 60, Multiple3And5(16))
	})

	t.Run("large number 100", func(t *testing.T) {
		assert.Equal(t, 2318, Multiple3And5(100))
	})

	t.Run("edge case 1 should return 0", func(t *testing.T) {
		assert.Equal(t, 0, Multiple3And5(1))
	})

	t.Run("edge case 0 should return 0", func(t *testing.T) {
		assert.Equal(t, 0, Multiple3And5(0))
	})
}
