package kata2_test

import (
	"golang-kata/kata2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKata1(t *testing.T) {
	t.Run("should return 2 if it prime number", func(t *testing.T) {
		result := kata2.Divisors(5)
		assert.Equal(t, 2, result)
	})

	t.Run("should return 1 if input is 1", func(t *testing.T) {
		result := kata2.Divisors(1)
		assert.Equal(t, 1, result)
	})

	t.Run("should return number of divisor of input", func(t *testing.T) {
		result := kata2.Divisors(12)
		assert.Equal(t, 6, result)
	})

	t.Run("should return number of divisor of input (2)", func(t *testing.T) {
		result := kata2.Divisors(30)
		assert.Equal(t, 8, result)
	})
}
