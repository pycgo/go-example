package main

import "fmt"

func main() {
	for n := 2; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
