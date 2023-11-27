package kata2_test

import (
	"golang-kata/kata2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKata1(t *testing.T) {
	t.Run("", func(t *testing.T) {
		result := kata2.Divisors(5)
		assert.Equal(t, 2, result)
	})
}
