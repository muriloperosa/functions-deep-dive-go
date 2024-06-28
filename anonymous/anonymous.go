package anonymous

import "fmt"

func Anonymous() {
	numbers := []int{1, 2, 3}

	// anonymous functions
	transformed := transformNumbers(&numbers, func(i int) int {
		return i * 2
	})

	fmt.Println(transformed)

	// closure
	transformedFactor4 := transformNumbers(&numbers, createTransformer(4))
	fmt.Println(transformedFactor4)

	transformedFactor3 := transformNumbers(&numbers, createTransformer(3))
	fmt.Println(transformedFactor3)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}
