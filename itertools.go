package helpers

func generateCombinations[T any](arr []T, length int, combination []T, results chan []T) {
	if length == 0 {
		temp := make([]T, len(combination))
		copy(temp, combination)
		results <- temp
	} else {
		for i := 0; i <= len(arr)-length; i++ {
			combination[len(combination)-length] = arr[i]
			generateCombinations(arr[i+1:], length-1, combination, results)
		}
	}
}

func Combinations[T any](arr []T, length int) <-chan []T {
	results := make(chan []T)
	go func() {
		defer close(results)
		combination := make([]T, length)
		generateCombinations(arr, length, combination, results)
	}()
	return results
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func permutations[T any](arr []T, start, end int, length int, results chan []T) {
	if start == length {
		temp := make([]T, length)
		copy(temp, arr[:length])
		results <- temp
	} else {
		for i := start; i <= end; i++ {
			swap(arr, start, i)
			permutations(arr, start+1, end, length, results)
			swap(arr, start, i) // backtrack
		}
	}
}

func Permutations[T any](arr []T, length int) <-chan []T {
	if length == -1 {
		length = len(arr)
	}
	results := make(chan []T)
	go func() {
		defer close(results)
		permutations(arr, 0, len(arr)-1, length, results)
	}()
	return results
}

func PermutationsList[T any](arr []T, length int) [][]T {
	perms := [][]T{}
	for perm := range Permutations(arr, length) {
		perms = append(perms, perm)
	}
	return perms
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func NumPermutations[T any](arr []T) int {
	return factorial(len(arr))
}

func GeneratorToArray[T any](c <-chan []T) [][]T {
	arr := [][]T{}
	for item := range c {
		arr = append(arr, item)
	}
	return arr
}
