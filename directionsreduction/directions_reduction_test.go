package directionsreduction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirReduc(t *testing.T) {
	t.Run("should cancel all directions when they form opposite pairs", func(t *testing.T) {
		input := []string{"NORTH", "SOUTH", "EAST", "WEST"}
		expected := []string{}
		assert.Equal(t, expected, DirReduc(input))
	})

	t.Run("should handle cascading reductions", func(t *testing.T) {
		input := []string{"NORTH", "EAST", "WEST", "SOUTH", "WEST", "WEST"}
		expected := []string{"WEST", "WEST"}
		assert.Equal(t, expected, DirReduc(input))
	})

	t.Run("should return non-reducible path unchanged", func(t *testing.T) {
		input := []string{"NORTH", "WEST", "SOUTH", "EAST"}
		expected := []string{"NORTH", "WEST", "SOUTH", "EAST"}
		assert.Equal(t, expected, DirReduc(input))
	})

	t.Run("should return single direction unchanged", func(t *testing.T) {
		assert.Equal(t, []string{"NORTH"}, DirReduc([]string{"NORTH"}))
	})

	t.Run("should cancel simple opposite pairs", func(t *testing.T) {
		assert.Equal(t, []string{}, DirReduc([]string{"NORTH", "SOUTH"}))
		assert.Equal(t, []string{}, DirReduc([]string{"SOUTH", "NORTH"}))
		assert.Equal(t, []string{}, DirReduc([]string{"EAST", "WEST"}))
		assert.Equal(t, []string{}, DirReduc([]string{"WEST", "EAST"}))
	})

	t.Run("should leave single direction after canceling pair", func(t *testing.T) {
		assert.Equal(t, []string{"WEST"}, DirReduc([]string{"NORTH", "SOUTH", "WEST"}))
	})

	t.Run("main example from problem description", func(t *testing.T) {
		input := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}
		expected := []string{"WEST"}
		assert.Equal(t, expected, DirReduc(input))
	})

	t.Run("should handle multiple consecutive opposite pairs", func(t *testing.T) {
		assert.Equal(t, []string{}, DirReduc([]string{"NORTH", "SOUTH", "EAST", "WEST"}))
	})

	t.Run("should handle complex cascading scenarios", func(t *testing.T) {
		assert.Equal(t, []string{}, DirReduc([]string{"NORTH", "EAST", "WEST", "SOUTH"}))
	})

	t.Run("should handle mixed reducible and non-reducible parts", func(t *testing.T) {
		assert.Equal(t, []string{"WEST", "SOUTH"}, DirReduc([]string{"NORTH", "SOUTH", "WEST", "SOUTH"}))
	})
}
