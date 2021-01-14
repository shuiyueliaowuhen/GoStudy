package main

import "fmt"

//常量
const pi = 3.1415926

//批量,如果后面不赋值。默认是上一个值
const (
	n0 = 10
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota
	a2
	a3
)

const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)

const (
	c1 = iota //0
	c2 = 100  //100
	c3 = iota //2
	c4        //3
)

const (
	d1, d2 = 1 + iota, 2 + iota
	d3, d4 = 1 + iota, 2 + iota
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Println(n0)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)
	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
}
