package main

import "fmt"

func main() {
	c1, c2, c3 := make(chan int), make(chan string), make(chan bool)
	go Go(c1, c2, c3)
	c1 <- 1
	c1 <- 2
	c2 <- "a"
	c2 <- "c"
	<-c3
	close(c1)
	close(c2)
	close(c3)

}

func Go(c1 chan int, c2 chan string, c3 chan bool)  {
	for  {
		select {
		case v, ok := <-c1 :
			if !ok {
				c3 <- true
				close(c3);
				break;
			}
			fmt.Println(v)
		case v, ok := <- c2 :
			if !ok {
				c3 <- true
				close(c3);
				break;
			}
			fmt.Println(v)
		}
	}
}
