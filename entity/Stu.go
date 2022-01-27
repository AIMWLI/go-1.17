package entity

type Stu struct {
	Name string
	Age  int
}

func NewStu(name string, age int) *Stu {
	return &Stu{Name: name, Age: age}
}

// *Stu 类型是 ChangeAge 方法的接受者
func (s *Stu) ChangeAge() {
	s.Age += 10000
}
