package main

import (
	"fmt"
	entity "practice/entity"
)

func main() {
	h := entity.Human{
		Person: &entity.Person{
			Name: "songjin",
		},
		Power: 10,
	}
	fmt.Println(h.Person.Name)
	//由于我们没有显式地给它一个字段名，所以我们可以隐式地访问组合类型的字段和函数。然而，Go 编译器确实给了它一个字段，下面这样完全有效：
	fmt.Println(h.Name)
	fmt.Println(h.Power)
}
