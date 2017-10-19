package main

import "fmt"
//定义
type person struct {
	name string
	age int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}

	return p2, p2.age - p1.age
}

func main() {
//练习一
	//常规赋值
	var P person
	P.name = "张三"
	P.age = 20

	F := person{"李四", 30}

	M := person{name: "王五", age: 40}

	N := new(person)
	N.age = 60
	N.name = "赵六"

	fmt.Println(P)
	fmt.Println(F)
	fmt.Println(M)
	fmt.Println(*N)

//练习二
	var tom person
	tom.name, tom.age = "tom", 18

	bob := person{age: 20, name: "bob"}

	paul := person{"paul", 25}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)

//练习三 匿名字段
	type Human struct {
		name string
		age int
		weight int
	}

	type Student struct {
		Human
		speciality string
	}

	mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)
	// 修改对应的备注信息
	mark.speciality = "AI"
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	// 修改他的年龄信息
	fmt.Println("Mark become old")
	mark.age = 46
	fmt.Println("His age is", mark.age)
	// 修改他的体重信息
	fmt.Println("Mark is not an athlet anymore")
	mark.weight += 60
	fmt.Println("His weight is", mark.weight)

	mark.Human = Human{"Marcus", 55, 220}
	mark.Human.age -= 1
	fmt.Println("His weight is", mark)

//练习四
	type Skills []string
	type H struct {
		name string
		age int
		weight int
	}

	type Student2 struct {
		H
		Skills
		int
		speciality string
	}

	jane := Student2{H:H{name:"jane", age: 30, weight: 20}, speciality: "test"}
	lisi := Student2{H{"lisi", 30, 20}, []string{"1"}, 3, "test"}
	jane.H = H{name:"1", age: 2, weight: 3}
	fmt.Println("His weight is", jane)
	fmt.Println("His weight is", lisi)
}
