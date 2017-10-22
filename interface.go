package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

//human实现Sayhi方法
func (h *Human) SayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//human实现Sing方法
func (h *Human) Sing(lyrics string)  {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

//human 实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//// Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //此句可以分成多行
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32)  {
	s.loan += amount
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32)  {
	e.money -= amount
}

//嵌入接口
type People interface {
	SayHi()
}
// 定义interface
type Men interface {
	//SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
	People //将People接口嵌入到Men，Men对象此时也拥有了SayHi方法
}

type YoungChap interface {
	SayHi()
	sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

//测试接口是否被继承
func TestMInterface(m Men)  {
	fmt.Println("Men interface")
}

func TestSInterface(s Student)  {
	fmt.Println("student interface")
}

func TestEInterface(e Employee)  {
	if p, exists := e.(Human); exists {
		fmt.Println(p.name)
		return
	}
	fmt.Println("ElderlyGent interface")
}

func main()  {
	//当一个类型实现了接口定义的所有方法，它就默认继承了这个接口
	h := &Human{"张三", 20, "123"}
	s := Student{Human{"李四", 20, "123"}, "test", 123}
	e := Employee{Human{"王五", 20, "123"}, "test", 123}
	e.SpendSalary(15);
	fmt.Println(e)
	TestMInterface(h);
	TestSInterface(s);
	TestEInterface(e);
}















