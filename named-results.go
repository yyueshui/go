package main

import "fmt"

func spilt(sum int) (x, y int) {
	x = sum * 4
	y = sum + 4

	return
}

func main() {
	fmt.Println(spilt(12))
}