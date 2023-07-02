package helpers

import "fmt"

func ExampleCombinations() {
	nums := []int{1, 2, 3}
	for combo := range Combinations(nums, 2) {
		fmt.Println(combo)
	}
	// Output:
	// [1 2]
	// [1 3]
	// [2 3]
}

func ExampleGeneratorToArray() {
	nums := []int{1, 2, 3}
	combos := Combinations(nums, 2)
	fmt.Println(GeneratorToArray(combos))
	// Output:
	// [[1 2] [1 3] [2 3]]
}

func ExampleNumPermutations() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(NumPermutations(nums))
	// Output: 120
}

func ExamplePermutations() {
	nums := []int{1, 2, 3}
	for p := range Permutations(nums, -1) {
		fmt.Println(p)
	}
	// Output:
	// [1 2 3]
	// [1 3 2]
	// [2 1 3]
	// [2 3 1]
	// [3 2 1]
	// [3 1 2]
}
