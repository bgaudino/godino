package godino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	// Test case 1: All conditions are true
	result1 := All(true, true, true, true)
	assert.True(t, result1, "failed to evaluate correctly for all true conditions")

	// Test case 2: Some conditions are false
	result2 := All(true, false, true, true)
	assert.False(t, result2, "failed to evaluate correctly for some false conditions")

	// Test case 3: All conditions are false
	result3 := All(false, false, false)
	assert.False(t, result3, "failed to evaluate correctly for all false conditions")

	// Test case 4: No conditions provided
	result4 := All()
	assert.True(t, result4, "failed to evaluate correctly for no conditions provided")
}

func TestAny(t *testing.T) {
	// Test case 1: Some conditions are true
	result1 := Any(true, false, false, true)
	assert.True(t, result1, "failed to evaluate correctly for some true conditions")

	// Test case 2: All conditions are false
	result2 := Any(false, false, false)
	assert.False(t, result2, "failed to evaluate correctly for all false conditions")

	// Test case 3: No conditions provided
	result3 := Any()
	assert.False(t, result3, "failed to evaluate correctly for no conditions provided")
}

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
