package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	t.Run("should return the maximum element and no error for a non-empty list", func(t *testing.T) {
		max, err := Max(2, 5, 1, 4)

		assert.Equal(t, 5, max)
		assert.NoError(t, err)
	})

	t.Run("should return zero value and and an error for an empty list", func(t *testing.T) {
		max, err := Max[int]()

		assert.Equal(t, 0, max)
		assert.Error(t, err)
	})
}

func TestMin(t *testing.T) {
	t.Run("should return the minimum element and no error for a non-empty list", func(t *testing.T) {
		min, err := Min(2, 5, 1, 4)

		assert.Equal(t, 1, min)
		assert.NoError(t, err)
	})

	t.Run("should return zero value and an error for an empty list", func(t *testing.T) {
		min, err := Min[int]()

		assert.Equal(t, 0, min)
		assert.Error(t, err)
	})
}

func TestProd(t *testing.T) {
	t.Run("should return the product of numbers", func(t *testing.T) {
		product := Prod(2, 3, 4)

		assert.Equal(t, 24, product)
	})
}

func TestSum(t *testing.T) {
	t.Run("should return the sum of numbers", func(t *testing.T) {
		sum := Sum(1, 2, 3, 4)

		assert.Equal(t, 10, sum)
	})
}
