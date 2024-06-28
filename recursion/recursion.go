package recursion

import "fmt"

func Recursion() {
	fmt.Println(factorial(6))
}

func factorial(number int) int {

	// recursive solution
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)

	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result *= i
	// }
	// return result
}
