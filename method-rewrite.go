package main

import "fmt"

type Human struct {
	name string
	age int
	phone int
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

//Human方法
func (h Human) SayHello() {
	fmt.Printf("Hi, I am %s you can call me on %d\n", h.name, h.phone)
}

func (e Employee) SayHello()  {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %d\n", e.name, e.company, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, 222}, "MIT"}
	sam := Employee{Human{"Sam", 45, 111}, "Golang Inc"}

	mark.SayHello()
	sam.SayHello()
}