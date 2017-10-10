package main

import "fmt"

//由于长度也是数组类型的一部分，因此[3]int与[4]int是不同的类型，数组也就不能改变长度。数组之间的赋值是值的赋值，即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针。如果要使用指针，那么就需要用到后面介绍的slice类型了。
func main() {
	//普通类型声明
	var arr [10]int
	arr[0] = 1
	arr[1] = 2

	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回1
	fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0

	fmt.Print("\n")

	//短声明
	a := [3]int{1, 2, 3} // 声明了一个长度为3的int数组
	b := [10]int{4, 5, 6} //声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{7, 8, 9} //可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Println(a[0], b[9], c[2])
	fmt.Print("\n")
	//二维数组
	// 声明了一个二维数组，该数组以两个数组作为元素，其中每个数组中又有4个int类型的元素
	doubleArray := [2][4]int{[4]int{1, 2}, [4]int{1, 2, 3, 4}}

	// 上面的声明可以简化，直接忽略内部的类型
	easyDoubleArray := [2][4]int{{1, 2}, {1, 2, 3, 4}}

	fmt.Println(doubleArray[0][1], easyDoubleArray[1][3])
	fmt.Print("\n")
}
