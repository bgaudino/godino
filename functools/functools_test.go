package functools

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestEvery(t *testing.T) {
	// Test case 1: Every element satisfies the condition
	arr1 := []int{2, 4, 6, 8, 10}
	every1 := Every(arr1, func(num int) bool {
		return num%2 == 0
	})
	assert.True(t, every1, "failed to evaluate correctly for all elements satisfying the condition")

	// Test case 2: Not every element satisfies the condition
	arr2 := []int{2, 4, 6, 7, 10}
	every2 := Every(arr2, func(num int) bool {
		return num%2 == 0
	})
	assert.False(t, every2, "Everyfailed to evaluate correctly for not all elements satisfying the condition")

	// Test case 3: Empty list
	emptyArr := []int{}
	arr3 := Every(emptyArr, func(num int) bool {
		return num%2 == 0
	})
	assert.True(t, arr3, "failed to evaluate correctly for an empty list")
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	filtered := Filter(arr, func(num int) bool {
		return num%2 == 0
	})
	expected := []int{2, 4}
	assert.Equal(t, expected, filtered)
}

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	t.Run("should find value", func(t *testing.T) {
		value, found := Find(arr, func(num int) bool {
			return num%2 == 0
		})
		assert.True(t, found, "failed to find value")
		assert.Equal(t, 2, value, "returned incorrect value")
	})

	t.Run("should not find value", func(t *testing.T) {
		value, found := Find(arr, func(num int) bool {
			return num > 10
		})
		assert.False(t, found, "incorrectly found a value")
		assert.Equal(t, 0, value, "returned non-zero value for non-existing condition")
	})
}

func TestForEach(t *testing.T) {
	arr := []int{1, 2, 3}
	num := 0
	f := func(i int, n int) {
		num += n + i
	}
	ForEach(arr, f)
	assert.Equal(t, 9, num)

}

func TestForEachRef(t *testing.T) {
	arr := []int{1, 2, 3}
	f := func(i int, n *int) {
		*n += i
	}
	ForEachRef(arr, f)
	assert.Equal(t, []int{1, 3, 5}, arr)
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	mapped := Map(arr, func(num int) int {
		return num * num
	})
	expected := []int{1, 4, 9, 16, 25}
	assert.Equal(t, expected, mapped, "failed to map numbers to their squares")

	mappedString := Map(arr, func(num int) string {
		return "Number: " + strconv.Itoa(num)
	})
	expectedStrings := []string{"Number: 1", "Number: 2", "Number: 3", "Number: 4", "Number: 5"}
	assert.Equal(t, expectedStrings, mappedString, "failed to map numbers to strings")
}

func TestReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	sum := Reduce(arr, func(acc, num int) int {
		return acc + num
	}, 0)
	expectedSum := 15
	assert.Equal(t, expectedSum, sum, "failed to calculate the sum")

	product := Reduce(arr, func(acc, num int) int {
		return acc * num
	}, 1)
	expectedProduct := 120
	assert.Equal(t, expectedProduct, product, "failed to calculate the product")

	m := Reduce(arr, func(acc map[int]int, num int) map[int]int {
		acc[num] = num * num
		return acc
	}, make(map[int]int))
	expectedMap := map[int]int{1: 1, 2: 4, 3: 9, 4: 16, 5: 25}
	assert.True(t, maps.Equal(m, expectedMap), "failed to make square map")
}

func TestSome(t *testing.T) {
	// Test case 1: Some elements satisfy the condition
	arr1 := []int{1, 2, 3, 4, 5}
	some1 := Some(arr1, func(num int) bool {
		return num%2 == 0
	})
	assert.True(t, some1, "failed to evaluate correctly for some elements satisfying the condition")

	// Test case 2: No element satisfies the condition
	arr2 := []int{1, 3, 5, 7, 9}
	some2 := Some(arr2, func(num int) bool {
		return num%2 == 0
	})
	assert.False(t, some2, "failed to evaluate correctly for no elements satisfying the condition")

	// Test case 3: Empty list
	emptyArr := []int{}
	some3 := Some(emptyArr, func(num int) bool {
		return num%2 == 0
	})
	assert.False(t, some3, "failed to evaluate correctly for an empty list")
}
