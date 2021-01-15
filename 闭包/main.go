package main

import "fmt"

func a(name string) func() {
	return func() {
		fmt.Println("你好，", name)
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	r := a("大帅比") //闭包=函数 + 外层变量饮用
	fmt.Printf("%T\n", r)
	r()
	println("==========闭包升级=======")
	r1, r2 := calc(10)
	println(r1(5))
	println(r2(3))

}
