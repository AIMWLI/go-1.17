package main

import "fmt"

func main() {
	var scores [10]int
	scores[0] = 20
	scores2 := [4]int{10, 20, 30, 40}
	for index, value := range scores2 {
		fmt.Println(index, value)
	}
	//在 Go 语言中，我们很少直接使用数组。取而代之的是使用切片。
	// 创建切片集中方式
	scores3 := []int{2, 34, 6, 8}
	ints := make([]int, 10) // 容量是 10 的切片
	fmt.Println(scores3)
	fmt.Println(ints)
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // [0 1 2 3 4 5 6 7 8 9] len=10,cap=10
	s1 := s[0:5]                             // [0 1 2 3 4] len=5,cap=10
	s11 := s[0:6]
	s2 := s[5:] // [5 6 7 8 9] len=5,cap=5
	s22 := s[6:]
	fmt.Println(s1, s11)
	fmt.Println(s2, s22)

}
