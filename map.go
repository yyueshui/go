package main

import "fmt"

/**
	使用map过程中需要注意的几点：
		map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取
		map的长度是不固定的，也就是和slice一样，也是一种引用类型
		内置的len函数同样适用于map，返回map拥有的key的数量
		map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典值改为11
		map和其他基本型别不同，它不是thread-safe，在多个go-routine存取时，必须使用mutex lock机制
 */
func main() {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string]int

	numbers := make(map[string]int)

	numbers["a"] = 1
	numbers["b"] = 2;

	fmt.Println(numbers["a"])

	// 初始化一个字典
	rating := map[string]string{"c": "5.1", "g": "3", "a": "1"}
	fmt.Println(rating["c"])

	// 判断map里是否包含该元素 map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	ratingC, existsC := rating["c"]

	if existsC {
		fmt.Println("c is in the map and its rating is ", ratingC)
	} else {
		fmt.Println("We have no rating associated with c in the map")
	}

	ratingB, existsB := rating["b"]
	if existsB {
		fmt.Println("b is in the map and its rating is ", ratingB)
	} else {
		fmt.Println("We have no rating associated with b in the map")
	}

	//map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变，另一个也相应的改变
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	fmt.Println("m", m["Hello"])
	m1 := m
	m1["Hello"] = "Salut"  // 现在m["hello"]的值已经是Salut了

	fmt.Println("m1", m1["Hello"])
	fmt.Println("m", m["Hello"])

}
