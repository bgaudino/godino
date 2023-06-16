package collections

import "golang.org/x/exp/maps"

type Dict[K comparable, V any] map[K]V

type DictItem[K comparable, V any] struct {
	key   K
	value V
}

func (dict Dict[K, V]) Clear() {
	maps.Clear(dict)
}

func (dict Dict[K, V]) Copy() Dict[K, V] {
	copy := make(Dict[K, V])
	maps.Copy(copy, dict)
	return copy
}

func (dict Dict[K, V]) Get(key K, fallback V) V {
	value, ok := dict[key]
	if ok {
		return value
	}
	return fallback
}

func (dict Dict[K, V]) Has(key K) bool {
	_, ok := dict[key]
	return ok
}

func (dict Dict[K, V]) Items() []DictItem[K, V] {
	items := make([]DictItem[K, V], len(dict))
	i := 0
	for k, v := range dict {
		items[i] = DictItem[K, V]{key: k, value: v}
		i++
	}
	return items
}

func (dict Dict[K, V]) Keys() []K {
	keys := make([]K, len(dict))
	i := 0
	for k := range dict {
		keys[i] = k
		i++
	}
	return keys
}

func (dict Dict[K, V]) Pop(key K, fallback V) (V, bool) {
	value, ok := dict[key]
	if ok {
		delete(dict, key)
		return value, true
	}
	return fallback, false
}

func (dict Dict[K, V]) SetDefault(key K, d V) V {
	_, ok := dict[key]
	if !ok {
		dict[key] = d
	}
	return dict[key]
}

func (dict1 Dict[K, V]) Update(dict2 Dict[K, V]) {
	for key, value := range dict2 {
		dict1[key] = value
	}
}

func (dict Dict[K, V]) Values() []V {
	values := make([]V, len(dict))
	i := 0
	for _, v := range dict {
		values[i] = v
		i++
	}
	return values
}
