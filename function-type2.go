package main

import "fmt"

type testInt func(int) bool

func isEven(num int) bool {
	if num % 2 == 0 {
		return false
	}
	return true
}

func isOdd(num int) bool {
	if num % 2 == 0 {
		return true
	}
	return  false
}

func filter(slice []int, f testInt) []int {
	var newSlice []int
	for _, v := range slice {
		if f(v) {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

func main()  {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(filter(numbers, isOdd))
	fmt.Println(filter(numbers, isEven))
}