package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// scoreMap := make(map[string]int, 8)
	// scoreMap["张三"] = 8
	// scoreMap["李四"] = 19
	// fmt.Println(scoreMap)
	// fmt.Println(scoreMap["李四"])
	// fmt.Printf("type of a:%T\n", scoreMap)

	// //判断值是否存在
	// fmt.Println("========判断值是否存在==========")
	// v, ok := scoreMap["张三"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("没有张三")
	// }

	// //遍历map
	// fmt.Println("=========遍历map=========")
	// for k, v := range scoreMap {
	// 	fmt.Println(k, ":", v)
	// }

	// fmt.Println("=========遍历map key=========")
	// for k := range scoreMap {
	// 	fmt.Println(k)
	// }

	// //删除
	// fmt.Println("=========map删除=========")
	// delete(scoreMap, "张三")
	// fmt.Println(scoreMap)

	//按照指定顺序遍历map
	fmt.Println("=========按照指定顺序遍历map=========")
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}
