package main

import "fmt"

func main() {
	i := 100

	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%o\n", i)

	str := "str"

	fmt.Printf("%s\n", str)
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
}
