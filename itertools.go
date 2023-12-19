package godino

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

// Returns an iterable containing all possible subsets of the array of the given length
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

// Returns an iterable containing all possible orderings of subsets of array of the given length.
// If -1 is provided for the length, the length of the array is used.
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

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Returns the number of permutations of the array
func NumPermutations[T any](arr []T) int {
	return factorial(len(arr))
}

// Converts a channel used as a generator to an array
func GeneratorToArray[T any](c <-chan []T) [][]T {
	arr := [][]T{}
	for item := range c {
		arr = append(arr, item)
	}
	return arr
}
