package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["11"] = 1
	m["12"] = 1
	fmt.Println(m)
	//循环读取元素
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "11")
	fmt.Println("after delete: ", m)
}
