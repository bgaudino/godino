package godino

import "golang.org/x/exp/maps"

// A wrapper around a map with several convenience methods
type Dict[K comparable, V any] map[K]V

type DictItem[K comparable, V any] struct {
	Key   K
	Value V
}

// Removes all elements from the dictionary
func (dict Dict[K, V]) Clear() {
	maps.Clear(dict)
}

// Returns a copy of the dictionary
func (dict Dict[K, V]) Copy() Dict[K, V] {
	copy := make(Dict[K, V])
	maps.Copy(copy, dict)
	return copy
}

// Returns the value associated with given key.
// In the case the the key is not present, a fallback value is returned if provided.
// Otherwise the zero-value for the value type is returned.
func (dict Dict[K, V]) Get(key K, fallback ...V) V {
	value, ok := dict[key]
	if !ok && len(fallback) >= 1 {
		return fallback[0]
	}
	return value
}

// Returns true if the dictionary contains the given key.
func (dict Dict[K, V]) Has(key K) bool {
	_, ok := dict[key]
	return ok
}

// Returns an array of dictionary items (a struct with Key and Value fields)
func (dict Dict[K, V]) Items() []DictItem[K, V] {
	items := make([]DictItem[K, V], len(dict))
	i := 0
	for k, v := range dict {
		items[i] = DictItem[K, V]{Key: k, Value: v}
		i++
	}
	return items
}

// Returns an array of keys present in the dictionary
func (dict Dict[K, V]) Keys() []K {
	keys := make([]K, len(dict))
	i := 0
	for k := range dict {
		keys[i] = k
		i++
	}
	return keys
}

// Removes the given key from the dictionary and returns it's associated value.
// If the key is not present in the dictionary, a fallback is returned if provided.
// Otherwise a zero-value is returned.
// Returns a second boolean value which is true if the key was present in the dictionary.
func (dict Dict[K, V]) Pop(key K, fallback ...V) (V, bool) {
	value, ok := dict[key]
	if ok {
		delete(dict, key)
	} else if len(fallback) >= 1 {
		value = fallback[0]
	}
	return value, ok
}

// Sets the value for the given key is the key is not already present.
// Returns the value for the given key.
func (dict Dict[K, V]) SetDefault(key K, d V) V {
	_, ok := dict[key]
	if !ok {
		dict[key] = d
	}
	return dict[key]
}

// Updates the keys and values from the given dictionary
func (dict1 Dict[K, V]) Update(dict2 Dict[K, V]) {
	for key, value := range dict2 {
		dict1[key] = value
	}
}

// Returns an array of the values of the dictionary
func (dict Dict[K, V]) Values() []V {
	values := make([]V, len(dict))
	i := 0
	for _, v := range dict {
		values[i] = v
		i++
	}
	return values
}
