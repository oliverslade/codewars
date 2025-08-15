package movezeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZero(t *testing.T) {
	t.Run("codewar example", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0},
			MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}),
		)
	})

	t.Run("empty array", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{},
			MoveZeros([]int{}),
		)
	})

	t.Run("array with no zeros", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1, 2, 3, 4, 5},
			MoveZeros([]int{1, 2, 3, 4, 5}),
		)
	})

	t.Run("array with only zeros", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{0, 0, 0, 0},
			MoveZeros([]int{0, 0, 0, 0}),
		)
	})

	t.Run("single zero", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{0},
			MoveZeros([]int{0}),
		)
	})

	t.Run("single non-zero", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{5},
			MoveZeros([]int{5}),
		)
	})

	t.Run("zeros at the beginning", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1, 2, 3, 0, 0, 0},
			MoveZeros([]int{0, 0, 0, 1, 2, 3}),
		)
	})

	t.Run("zeros at the end", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1, 2, 3, 0, 0, 0},
			MoveZeros([]int{1, 2, 3, 0, 0, 0}),
		)
	})

	t.Run("zeros in the middle", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1, 2, 4, 5, 0, 0},
			MoveZeros([]int{1, 2, 0, 0, 4, 5}),
		)
	})

	t.Run("negative numbers", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{-1, -2, 3, 4, 0, 0},
			MoveZeros([]int{-1, 0, -2, 0, 3, 4}),
		)
	})

	t.Run("mixed positive and negative with zeros", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1, -2, 3, -4, 5, 0, 0, 0},
			MoveZeros([]int{1, 0, -2, 0, 3, -4, 5, 0}),
		)
	})

	t.Run("large numbers", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1000, 2000, 3000, 0, 0},
			MoveZeros([]int{1000, 0, 2000, 0, 3000}),
		)
	})

	t.Run("alternating zeros and non-zeros", func(t *testing.T) {
		assert.Equal(
			t,
			[]int{1, 3, 5, 7, 0, 0, 0, 0},
			MoveZeros([]int{1, 0, 3, 0, 5, 0, 7, 0}),
		)
	})
}
