package main

import "fmt"

func args(args ...int)  {
	for k, v := range args {
		fmt.Println(k, v, "\n")
	}
}


func main() {
	args(1, 2, 3, 4)
}
