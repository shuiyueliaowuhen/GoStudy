package main

import "fmt"

func main() {
	// age := 10
	// if age > 18 {
	// 	fmt.Println("成年啦")
	// } else if age > 16 {
	// 	fmt.Println("未成年1")
	// } else {
	// 	fmt.Println("还是没成年")
	// }

	if age := 19; age > 18 { //age作用域只存在当前if语句中
		fmt.Println("asd")
	}
}
