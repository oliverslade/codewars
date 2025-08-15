package secondstostring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecondsToString(t *testing.T) {
	t.Run("zero seconds", func(t *testing.T) {
		assert.Equal(t, "now", secondsToString(0))
	})

	t.Run("single second", func(t *testing.T) {
		assert.Equal(t, "1 second", secondsToString(1))
	})

	t.Run("multiple seconds", func(t *testing.T) {
		assert.Equal(t, "15 seconds", secondsToString(15))
	})

	t.Run("single minute and seconds", func(t *testing.T) {
		assert.Equal(t, "1 minute and 2 seconds", secondsToString(62))
	})

	t.Run("exact minute", func(t *testing.T) {
		assert.Equal(t, "1 minute", secondsToString(60))
	})

	t.Run("multiple minutes", func(t *testing.T) {
		assert.Equal(t, "2 minutes", secondsToString(120))
	})

	t.Run("multiple minutes with seconds", func(t *testing.T) {
		assert.Equal(t, "2 minutes and 5 seconds", secondsToString(125))
	})

	t.Run("hours with minutes and seconds", func(t *testing.T) {
		assert.Equal(t, "1 hour, 1 minute and 2 seconds", secondsToString(3662))
	})

	t.Run("years with days hours and minutes", func(t *testing.T) {
		assert.Equal(t, "4 years, 68 days, 3 hours and 4 minutes", secondsToString(132030240))
	})
}
