//递归  函数一直调用本身不断往前算
package main

import "fmt"

//求 1*2*3*4* ...
func recursion(num int) int {
	if num == 1 {
		return 1
	}
	return num * recursion(num-1)
}

//斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	sum := recursion(4)
	fmt.Println("乘积 ", sum)
	fmt.Print("斐波那契 ", fib(7))
}
