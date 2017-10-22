package main

import "fmt"

//一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值
func main() {
	// a可以存储任意类型的数值
	var a interface{}
	var i int = 5

	s := "hello world"

	a = i
	fmt.Println(a);

	a = s

	fmt.Println(a);
}
