package main

import "fmt"

func main() {
	//切片定义
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	//初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"a", "b", "c"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	//长度 容量
	fmt.Printf("len(s1)%d,cap(s1)%d\n", len(s1), cap(s1))

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s3 := a1[0:4] //基于数组切割，左闭右开（左包含右不包含）
	s4 := a1[1:4]
	fmt.Println(s3)
	fmt.Printf("len(s3)%d,cap(s1)%d\n", len(s3), cap(s3)) //容量是底层数组的第一个元素到最后到长度
	fmt.Printf("len(s4)%d,cap(s4)%d\n", len(s4), cap(s4)) //容量是底层数组的第一个元素到最后到长度

	//由切片得到切片
	s5 := s3[1:]
	fmt.Printf("len(s5)%d,cap(s5)%d\n", len(s5), cap(s5))

}
