package main

import "fmt"

type testInt func(int) bool

func isOdd(integer int) bool {
	if 	integer % 2 == 0 {
		return true
	}
	return false
}

func isEven(integer int) bool {
	if integer % 2 == 0 {
		return false
	}

	return true
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}

	return result
}
func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice, "\n")
	even := filter(slice, isEven)
	odd := filter(slice, isOdd)

	fmt.Println(even, "\n")
	fmt.Println(odd, "\n")
}
