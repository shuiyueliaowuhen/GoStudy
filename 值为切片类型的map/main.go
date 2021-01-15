package main

import "fmt"

func main() {
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	sliceMap[key] = make([]string, 3)
	sliceMap[key] = append(sliceMap[key], "北京", "上海")
	fmt.Println(sliceMap)
}
