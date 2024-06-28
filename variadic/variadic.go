package variadic

import "fmt"

func Variadic() {
	// numbers := []int{1, 10, 15}
	// fmt.Println(sumUp(numbers))

	fmt.Println(sumUp(1, 10, 15, 49))

	numbers := []int{1, 10, 15}
	fmt.Println(sumUp(numbers...))
}

// variadic function
func sumUp(numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}
