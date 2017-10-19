package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	//method也是可以继承的。如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method。
	mark := Student{Human{"张三", 20, "1111"}, "一中"}
	sam := Employee{Human{"李四", 33, "2222"}, "百度"}

	mark.SayHi()
	sam.SayHi()
}
