package main

import "fmt"

func main() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d,value:%v\n", index, value)
	}
	fmt.Println("after init")
	//对第一个元素赋值
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["飞飞"] = "大帅比"
	for index, value := range mapSlice {
		fmt.Printf("index:%d,value:%v\n", index, value)
	}
}
