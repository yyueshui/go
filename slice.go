package main

func main() {

	//声明
	//var fslice []int

	//声明并且初始化
	//slice := []byte{'a', 'b', 'c', 'd'}

	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g

	//var array [10]int
	//slice := array[2:4]
	//这个例子里面slice的容量是8，新版本里面可以指定这个容量
	//slice = array[2:4:7]
	//上面这个的容量就是7-2，即5。这样这个产生的新的slice就没办法访问最后的三个元素。
	//如果slice是这样的形式array[:i:j]，即第一个参数为空，默认值就是0。
}
