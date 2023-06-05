package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("should return the sum of numbers", func(t *testing.T) {
		sum := Sum(1, 2, 3, 4)

		assert.Equal(t, 10, sum)
	})
}

func TestProd(t *testing.T) {
	t.Run("should return the product of numbers", func(t *testing.T) {
		product := Prod(2, 3, 4)

		assert.Equal(t, 24, product)
	})
}

func TestMax(t *testing.T) {
	t.Run("should return the maximum element and true for a non-empty list", func(t *testing.T) {
		max, ok := Max(2, 5, 1, 4)

		assert.Equal(t, 5, max)
		assert.True(t, ok)
	})

	t.Run("should return zero value and false for an empty list", func(t *testing.T) {
		max, ok := Max[int]()

		assert.Equal(t, 0, max)
		assert.False(t, ok)
	})
}

func TestMin(t *testing.T) {
	t.Run("should return the minimum element and true for a non-empty list", func(t *testing.T) {
		min, ok := Min(2, 5, 1, 4)

		assert.Equal(t, 1, min)
		assert.True(t, ok)
	})

	t.Run("should return zero value and false for an empty list", func(t *testing.T) {
		min, ok := Min[int]()

		assert.Equal(t, 0, min)
		assert.False(t, ok)
	})
}
