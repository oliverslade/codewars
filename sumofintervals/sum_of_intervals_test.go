package sumofintervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOfIntervals(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		assert.Equal(t, 1, SumOfIntervals([][2]int{{1, 2}}))
	})

	t.Run("simple negative case", func(t *testing.T) {
		assert.Equal(t, 20, SumOfIntervals([][2]int{{-40, -20}}))
	})

	t.Run("multiple simple arrays", func(t *testing.T) {
		assert.Equal(t, 8, SumOfIntervals([][2]int{{1, 5}, {6, 10}}))
		assert.Equal(t, 9, SumOfIntervals([][2]int{{1, 2}, {6, 10}, {11, 15}}))
		assert.Equal(t, 7, SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}}))
		assert.Equal(t, 19, SumOfIntervals([][2]int{{1, 5}, {10, 20}, {1, 6}, {16, 19}, {5, 11}}))
	})

	t.Run("totally overlapping array", func(t *testing.T) {
		assert.Equal(t, 4, SumOfIntervals([][2]int{{1, 5}, {1, 5}}))
	})

	t.Run("partially overlapping array", func(t *testing.T) {
		assert.Equal(t, 7, SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}}))
	})

	t.Run("partially overlapping negative array", func(t *testing.T) {
		assert.Equal(t, 40, SumOfIntervals([][2]int{{-50, -10}, {-40, -20}}))
	})

	t.Run("partially overlapping negative to positive array", func(t *testing.T) {
		assert.Equal(t, 70, SumOfIntervals([][2]int{{-50, 10}, {-40, 20}}))
	})

	t.Run("large intervals", func(t *testing.T) {
		assert.Equal(t, 100_000_030, SumOfIntervals([][2]int{{0, 20}, {-100_000_000, 10}, {30, 40}}))
	})

	t.Run("very large arrays", func(t *testing.T) {
		input := [][2]int{
			{79, 88}, {1, 2}, {6, 35}, {76, 83}, {-17, 38}, {72, 99}, {85, 95}, {48, 94},
			{64, 85}, {22, 74}, {4, 86}, {27, 61}, {-70, 15}, {14, 77}, {-93, -30}, {-31, 7},
			{72, 90}, {-86, 99}, {-71, 6}, {37, 61}, {64, 70}, {-35, 11}, {79, 96}, {-13, 39},
			{-73, 69}, {45, 66}, {14, 31}, {9, 22}, {-41, -26}, {-49, -30}, {10, 34}, {74, 91},
			{-57, 86}, {86, 95}, {50, 63}, {44, 51}, {-43, 57}, {6, 53}, {-87, -11}, {-55, 2},
			{18, 29}, {69, 87},
		}
		assert.Equal(t, 192, SumOfIntervals(input))
	})
}
