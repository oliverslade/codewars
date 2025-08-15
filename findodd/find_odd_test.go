package findodd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOdd(t *testing.T) {
	t.Run("single element array", func(t *testing.T) {
		assert.Equal(t, 7, FindOdd([]int{7}))
		assert.Equal(t, 0, FindOdd([]int{0}))
	})

	t.Run("short multiple element array", func(t *testing.T) {
		assert.Equal(t, 2, FindOdd([]int{1, 1, 2}))
	})
}
