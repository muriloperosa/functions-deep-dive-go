package functionsasvalueandtypes

import "fmt"

type transformFn func(int) int

func FunctionsAsValueAndTypes() {
	numbers := []int{1, 2, 3, 4}

	// function as param
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled, tripled)

	// function as param -> returned as value
	doubledF := transformNumbers(&numbers, getTransformer("double"))
	tripledF := transformNumbers(&numbers, getTransformer("triple"))
	testF := transformNumbers(&numbers, getTransformer("triplex"))

	fmt.Println(doubledF, tripledF, testF)
}

// func transformNumbers(numbers *[]int, transform func(int) int) []int {
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformer(calc string) func(int) int {
	transformers := map[string]func(int) int{
		"double": double,
		"triple": triple,
	}

	if transformers[calc] == nil {
		return transformers["double"]
	}

	return transformers[calc]
}
