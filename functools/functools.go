package functools

func Every[L ~[]T, T any](arr L, f func(T) bool) bool {
	for _, v := range arr {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter[L ~[]T, T any](arr L, condition func(T) bool) []T {
	filtered := []T{}
	for _, v := range arr {
		if condition(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Find[L ~[]T, T any](arr L, condition func(T) bool) (value T, found bool) {
	for _, v := range arr {
		if condition(v) {
			return v, true
		}
	}
	return value, false
}

func ForEach[L ~[]T, T any](arr L, f func(index int, value T)) {
	for i := 0; i < len(arr); i++ {
		f(i, arr[i])
	}
}

func ForEachRef[L ~[]T, T any](arr L, f func(index int, value *T)) {
	for i := 0; i < len(arr); i++ {
		f(i, &arr[i])
	}
}

func Map[L ~[]T, T any, V any](arr L, f func(T) V) []V {
	mapped := []V{}
	for _, v := range arr {
		mapped = append(mapped, f(v))
	}
	return mapped
}

func Reduce[L ~[]T, T any, V any](arr L, f func(V, T) V, acc V) V {
	for _, v := range arr {
		acc = f(acc, v)
	}
	return acc
}

func Some[L ~[]T, T any](arr L, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}
	return false
}
