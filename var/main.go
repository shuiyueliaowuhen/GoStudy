package main

import "fmt"

// //声明变量
// var name string
// var age int
// var isOk bool

//批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {

	name = "sf"
	age = 18
	isOk = true

	//局部变量声明后必须使用，不然编译不通过
	// var name1 string

	fmt.Printf("name:%s", name)
	fmt.Println()

	//声明变量同时赋值
	var str string = "hhh"
	fmt.Println(str)

	//类型推导
	var str1 = "bbb"
	fmt.Println(str1)

	//简短变量声明，只能在函数内使用
	str3 := "ccc"
	fmt.Println(str3)

}
