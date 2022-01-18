package main

import "fmt"

func main() {
	num := 9
	if num < 0 {
		fmt.Println("不是正数")
	} else if num >= 10 {
		fmt.Println("大于10")
	} else {
		fmt.Println("ok")
	}
}
