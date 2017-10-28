package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan bool, 10)
	num := runtime.NumCPU()
	fmt.Println(num)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for j := 0; j < 1; j++ {
		<-c
	}

}

func Go(c chan bool, i int) {
	a := 1
	for j := 0; j < 10000; j++ {
		a ++
	}
	fmt.Println(i, a)
	//if i == 9 {
		c <- true
	//	close(c)
	//}
}