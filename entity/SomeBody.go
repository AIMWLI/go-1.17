package entity

import "fmt"

type Person struct {
	Name string
}
type Human struct {
	*Person
	Power int
}

func (h *Human) info() {
	fmt.Println("i am s% \n", h)
}
