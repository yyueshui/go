package main

import "fmt"

func main() {
	//如果有很多调用defer，那么defer是采用后进先出模式，所以如下代码会输出4 3 2 1 0
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
