package godino

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getDict() Dict[string, int] {
	return Dict[string, int]{
		"apple":  5,
		"banana": 3,
		"orange": 2,
	}
}
func TestDict(t *testing.T) {
	t.Run("should clear the dictionary", func(t *testing.T) {
		dict := getDict()
		dict.Clear()
		assert.Empty(t, dict)
	})

	t.Run("should return a copy of the dictionary", func(t *testing.T) {
		dict := getDict()
		copyDict := dict.Copy()
		assert.Equal(t, dict, copyDict)
	})

	t.Run("should retrieve the value for the key if it exists in the map else the giben fallback value", func(t *testing.T) {
		dict := getDict()
		value1 := dict.Get("apple", 0)
		assert.Equal(t, value1, 5, "failed to retrieve the existing key")

		value2 := dict.Get("mango", 10)
		assert.Equal(t, 10, value2, "failed to return the fallback value for non-existent key")
	})

	t.Run("should return whether the key exists in the dictionary", func(t *testing.T) {
		dict := getDict()
		has1 := dict.Has("apple")
		assert.True(t, has1, "failed to find the existing key")

		has2 := dict.Has("mango")
		assert.False(t, has2, "found the non-existent key")
	})

	t.Run("should return the dictionary's keys and values", func(t *testing.T) {
		dict := getDict()
		items := dict.Items()
		expectedItems := []DictItem[string, int]{
			{Key: "apple", Value: 5},
			{Key: "banana", Value: 3},
			{Key: "orange", Value: 2},
		}
		assert.ElementsMatch(t, expectedItems, items)

	})

	t.Run("should return the dictionary's keys", func(t *testing.T) {
		dict := getDict()
		keys := dict.Keys()
		expectedKeys := []string{"apple", "banana", "orange"}
		assert.ElementsMatch(t, expectedKeys, keys)
	})

	t.Run("should remove the item for the dictionary if the key exists and return the value", func(t *testing.T) {
		dict := getDict()
		value, ok := dict.Pop("apple", 0)
		assert.Equal(t, 5, value, "failed to retrieve the existing key")
		assert.True(t, ok, "incorrectly reported the key as non-existent")
		_, ok = dict["apple"]
		assert.False(t, ok, "failed to remove the key from the dictionary")

		value, ok = dict.Pop("mango", 10)
		assert.Equal(t, 10, value, "failed to return the fallback value for non-existent key")
		assert.False(t, ok, "incorrectly reported the key as existent")
	})

	t.Run("should set the value of the dictionary if the key is not already present, then return the associated value", func(t *testing.T) {
		dict := getDict()
		defaultValue1 := dict.SetDefault("apple", 0)
		assert.Equal(t, 5, defaultValue1, "failed to return the existing key's value")

		defaultValue2 := dict.SetDefault("mango", 10)
		assert.Equal(t, 10, defaultValue2, "failed to set the fallback value for non-existent key")
		assert.Equal(t, 10, dict["mango"], "failed to add the non-existent key")
	})

	t.Run("should update the dictionary with the keys and values of the given dictionary", func(t *testing.T) {
		dict := getDict()
		anotherDict := Dict[string, int]{
			"mango":  8,
			"banana": 4,
		}
		dict.Update(anotherDict)
		expectedDict := Dict[string, int]{
			"apple":  5,
			"banana": 4,
			"orange": 2,
			"mango":  8,
		}
		assert.Equal(t, expectedDict, dict)
	})

	t.Run("should return the values of the dictionary", func(t *testing.T) {
		dict := getDict()
		values := dict.Values()
		expectedValues := []int{5, 3, 2}
		assert.ElementsMatch(t, expectedValues, values)
	})
}
