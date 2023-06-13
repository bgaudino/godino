package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	arr := []int{1, 2, 3}
	combos := Combinations(arr, 2)
	comboArr := GeneratorToArray(combos)
	expected := [][]int{{1, 2}, {1, 3}, {2, 3}}
	assert.Equal(t, expected, comboArr)
}

func TestPermutations(t *testing.T) {
	t.Run("permutations of the entire array", func(t *testing.T) {
		arr := []int{1, 2, 3}
		perms := Permutations(arr, -1)
		permArry := GeneratorToArray(perms)
		expected := [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 2, 1},
			{3, 1, 2},
		}
		assert.ElementsMatch(t, expected, permArry)
	})

	t.Run("permutations of a length less than the array", func(t *testing.T) {
		arr := []int{1, 2, 3}
		perms := Permutations(arr, 2)
		permArry := GeneratorToArray(perms)
		expected := [][]int{
			{1, 2},
			{1, 3},
			{2, 1},
			{2, 3},
			{3, 2},
			{3, 1},
		}
		assert.ElementsMatch(t, expected, permArry)
	})
}
