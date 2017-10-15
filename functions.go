package main

import "fmt"

func add(x int) int {
	return x + 1
}

func addPoint(x *int) int {
	*x = *x + 1
	return *x
}

func main() {
	x := 3
	fmt.Println(3)
	//普通传值
	y := add(x)
	fmt.Println(y)
	fmt.Println(x)

	//引用传值
	z := addPoint(&x)
	fmt.Println(z)
	fmt.Println(x)
}