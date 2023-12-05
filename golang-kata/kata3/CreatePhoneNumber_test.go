package kata3_test

import (
	"golang-kata/kata3"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKata3(t *testing.T) {
	t.Run("should return telephone number form", func(t *testing.T) {
		actual := kata3.CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
		assert.Equal(t, "(123) 456-7890", actual)
	})
}
