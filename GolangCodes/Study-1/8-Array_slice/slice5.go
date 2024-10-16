package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}

	//截取[0,2)的元素  左闭右开
	s1 := s[0:2] // 1, 2
	s2 := s[:2]  // 1, 2,
	s3 := s[1:]  // 2, 3, 4
	s4 := s[:]   // 1, 2 ,3, 4

	//相当于引用
	s1[0] = 100
	s2[1] = 200
	s3[2] = 300
	s4[3] = 400
	fmt.Printf("s = %v\n", s)

	//copy可以拷贝底层数组
	s5 := make([]int, 3)
	copy(s5, s)
	s[1] = 1
	fmt.Printf("s5 = %v\n", s5)
}
