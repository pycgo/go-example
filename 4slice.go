package main

import (
	"fmt"
)

func main() {
	//切片初始化 长度最少是2
	slice1 := make([]string, 3)
	slice1[0] = "one"
	slice1 = append(slice1, "two")
	slice1 = append(slice1, "3")
	slice1 = append(slice1, "3")
	fmt.Println(slice1)
	fmt.Println("长度", len(slice1))
	//循环读取切片 _ 的意思是 返回的序号给它 但是可以不需要处理
	for _, value := range slice1 {
		if value != "" {
			fmt.Println(value, "pos")
		}
	}
}
