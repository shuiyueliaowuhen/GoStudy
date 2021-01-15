package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum1(2, 2))
	fmt.Println(sum2(1, 2, 3, 4))
	fmt.Println(sum3(2, 2, 3, 4))
	a, b := calc(2, 1) //多返回值
	fmt.Println(a, b)
	fmt.Println(calc2(4, 1)) //返回值命名
	fmt.Println(someFunc())  //返回nil
}

func sum(i int, j int) int {
	return i + j
}

//类型简写
func sum1(i, j int) int {
	return i + j
}

//可变参数
func sum2(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

//固定+可变参数
func sum3(i int, j ...int) int {
	for _, v := range j {
		i += v
	}
	return i
}

//多返回值
func calc(i, j int) (int, int) {
	a := i + j
	b := i - j
	return a, b
}

func calc2(i, j int) (sum, sub int) {
	sum = i + j
	sub = i - j
	return
}

func someFunc() []int {
	return nil
}
