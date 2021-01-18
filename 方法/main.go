package main

import "fmt"

type Person struct {
	name string
	age  int
}

//构造函数
func newPerson(name string, age int) *Person {
	return &Person{name: name, age: age}
}

func (p Person) dream() {
	fmt.Printf("%s做梦都想发财\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func main() {
	p := newPerson("呵呵", 19)
	fmt.Printf("%v\n", p)

	p.dream()

	p.SetAge(20)
	fmt.Printf("%v\n", p)
}
