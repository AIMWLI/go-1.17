package main

import (
	"fmt"
	"practice/entity"
)

func main() {
	// 1.冗长声明变量
	//var power int
	//power = 9000
	//fmt.Println(power)

	// 2. 简便声明 := 自动推断类型
	power := 1000 // 不可重复声明, 声明多个变量下次可重复声明
	name, power := "peter", 2000
	fmt.Println(name)
	fmt.Println(power)

	stu := entity.Stu{
		Name: "song",
		Age:  18,
	}

	stu1 := entity.Stu{Name: "jin"}
	stu2 := entity.Stu{"songjin", 29}
	fmt.Println(stu)
	fmt.Println(stu1)
	fmt.Println(stu2)
	newStu := plusAge(stu2)
	fmt.Println(newStu)

	//&符号的意思是对变量取地址
	//*符号的意思是对指针取值
	stu3 := &entity.Stu{"asf", 1}
	stu3.ChangeAge()
	fmt.Println(stu3.Name)
	fmt.Println(stu3.Age)
	/*
		func main() {
			var a *int // 存储的是int的指针，目前为空
			var b int = 4 // 存储的是int的值
			a = &b // a 指向 b 的地址
			a = b // a 无法等于 b，会报错，a是指针，b是值，存储的类型不同
			fmt.Println(a) // a:0xc00000a090(返回了地址)
			fmt.Println(*a) // *a:4(返回了值)
			fmt.Println(*&a) // *抵消了&，返回了0xc00000a090本身
			*a = 5 // 改变 a 的地址的值
			fmt.Println(b) // b:5，改变后 b 同样受到改变，因为 a 的地址是指向 b 的
		}

	*/

	//new(X) 的结果与 &X{} 相同。
	m := new(entity.Stu)
	m.Name = "newname"

	// 推荐后一种
	n := &entity.Stu{
		Age: 10,
	}
	fmt.Println(m)
	fmt.Println(n)
}

func plusAge(s entity.Stu) entity.Stu {
	s.Age += 10
	return s
}
