package helpers

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestAppend(t *testing.T) {
	arr := []int{1, 2, 3}
	Append(&arr, 4)
	assert.Equal(t, []int{1, 2, 3, 4}, arr)
}

func TestClearArray(t *testing.T) {
	arr := []int{1, 2, 3}
	Clear(&arr)
	assert.Empty(t, arr)
}

func TestCopyArray(t *testing.T) {
	arr := []int{1, 2, 3}
	assert.Equal(t, Copy(arr), arr)
}

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3}
	assert.True(t, Contains(arr, 3))
	assert.False(t, Contains(arr, 0))
}

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

func TestExtend(t *testing.T) {
	arr := []int{1, 2, 3}
	Extend(&arr, []int{4, 5, 6})
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, arr)
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

func TestIndex(t *testing.T) {
	t.Run("should return value", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		index := Index(arr, 3)
		expectedIndex := 2
		assert.Equal(t, expectedIndex, index)
	})

	t.Run("should return -1 (value is not present)", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		index := Index(arr, 6)
		expectedIndex := -1
		assert.Equal(t, expectedIndex, index)
	})
}

func TestInsert(t *testing.T) {
	arr := []int{1, 2, 3}

	Insert(&arr, 0, 0)
	expected := []int{0, 1, 2, 3}
	assert.Equal(t, expected, arr)

	Insert(&arr, 99, 2)
	expected = []int{0, 1, 99, 2, 3}
	assert.Equal(t, expected, arr)

	Insert(&arr, 100, len(arr))
	expected = []int{0, 1, 99, 2, 3, 100}
	assert.Equal(t, expected, arr)
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	mapped := Map(arr, func(num int) int {
		return num * num
	})
	expected := []int{1, 4, 9, 16, 25}
	assert.Equal(t, expected, mapped, "failed to map numbers to their squares")

	mappedString := Map[int, string](arr, func(num int) string {
		return "Number: " + strconv.Itoa(num)
	})
	expectedStrings := []string{"Number: 1", "Number: 2", "Number: 3", "Number: 4", "Number: 5"}
	assert.Equal(t, expectedStrings, mappedString, "failed to map numbers to strings")
}

func TestPopArray(t *testing.T) {
	arr := []int{1, 2, 3}
	popped := Pop(&arr)
	expectedItem := 3
	assert.Equal(t, expectedItem, popped, "Failed to return value")
	expected := []int{1, 2}
	assert.Equal(t, expected, arr, "Failed to modify list")
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

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	Reverse(&arr)
	expected := []int{4, 3, 2, 1}
	assert.Equal(t, expected, arr)
}

func TestShift(t *testing.T) {
	arr := []int{1, 2, 3}
	shifted := Shift(&arr)
	expectedItem := 1
	assert.Equal(t, expectedItem, shifted, "Failed to return value")
	expected := []int{2, 3}
	assert.Equal(t, expected, arr, "Failed to modify list")
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

func TestValueAt(t *testing.T) {
	t.Run("should return the value at the given index", func(t *testing.T) {
		arr := []int{1, 2, 3}
		item := ValueAt(arr, 2)
		assert.Equal(t, 3, item)
	})

	t.Run("should support negative index", func(t *testing.T) {
		arr := []int{1, 2, 3}
		item := ValueAt(arr, -2)
		assert.Equal(t, 2, item)
	})
}

func TestZip(t *testing.T) {
	// Test case 1: Valid input slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	slice3 := []int{7, 8, 9}
	expectedResult := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	result, err := Zip(slice1, slice2, slice3)
	assert.Nil(t, err, "Zip() returned an unexpected error")
	assert.Equal(t, expectedResult, result, "Zip() returned an incorrect result")

	// Test case 2: Slices with different lengths
	slice4 := []int{10, 11}
	_, err = Zip(slice1, slice2, slice3, slice4)
	expectedError := errors.New("Zip() received slices of different lengths")
	assert.EqualError(t, err, expectedError.Error(), "Zip() did not return the expected error")

	// Test case 3: No input slices
	_, err = Zip[any]()
	expectedError = errors.New("Zip() expected at least 1 argument but got 0")
	assert.EqualError(t, err, expectedError.Error(), "Zip() did not return the expected error")
}
