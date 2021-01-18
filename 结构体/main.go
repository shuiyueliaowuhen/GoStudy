package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p1 person
	p1.name = "ff"
	p1.age = 18
	fmt.Printf("%v\n", p1)
	fmt.Printf("%#v\n", p1)

	fmt.Println("====匿名结构体===")
	var user struct {
		name string
		age  int
	}
	user.name = "hh"
	user.age = 16
	fmt.Printf("%v\n", user)

	fmt.Println("====创建指针型结构体===")
	var p2 = new(person)
	fmt.Printf("%T\n", p2) //*main.person
	fmt.Printf("p2=%#v\n", p2)
	p2.name = "fh"
	p2.age = 17
	fmt.Printf("%v\n", p2)

	fmt.Println("使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。")
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	fmt.Printf("p3=%#v\n", p3)

	println("===结构体初始化===")
	var p4 person
	fmt.Printf("p4=%#v\n", p4)

	//===========使用键值对初始化
	println("===========使用键值对初始化")
	p5 := person{
		name: "hehe",
		age:  7,
	}
	fmt.Printf("p5%#v", p5)

}
