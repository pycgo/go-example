//递归  函数一直调用本身不断往前算

//求 1*2*3*4* ...

package main

import "fmt"

func recursion(num int) int {
	if num == 1 {
		return 1
	}
	return num * recursion(num-1)
}

func main() {
	sum := recursion(4)
	fmt.Print(sum)
}
