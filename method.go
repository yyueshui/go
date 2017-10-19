package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

//function
func area(r Rectangle) float64  {
	return r.width * r.height
}

//Rectangle 的method
func (r Rectangle) area() float64 {
	return r.width * r.height
}

//Circle的method
func (c Circle) area () float64 {
	return c.radius * c.radius * math.Pi
}

func main()  {
	//普通方法
	r1 := Rectangle{1, 2}

	fmt.Println("Area of r1 is: ", area(r1))

	fmt.Println("Area of r2 is: ", r1.area())

	c1 := Circle{20}
	c2 := Circle{10}
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
