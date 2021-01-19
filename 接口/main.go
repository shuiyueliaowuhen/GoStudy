package main

import "fmt"

type Person interface {
	Eat()
}

type boy struct{}
type girl struct{}

func (b boy) Eat() {
	fmt.Println("只会吃")
}

func (g girl) Eat() {
	fmt.Println("就会吃")
}

func main() {
	var p Person

	b := boy{}
	g := girl{}
	p = b
	p.Eat()

	p = g
	g.Eat()

	/**
	type 接口类型名 interface{
	    方法名1( 参数列表1 ) 返回值列表1
	    方法名2( 参数列表2 ) 返回值列表2
	    …
	}
	*/
}
