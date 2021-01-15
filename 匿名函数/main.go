package main

import "fmt"

func main() {
	//匿名函数保存到变量
	add := func(i, j int) int {
		return i + j
	}
	fmt.Println(add(1, 2))

	//自己执行匿名函数
	println("========自己执行匿名函数=========")
	func(s string) {
		println(s)
	}("asda")

}
