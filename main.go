package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func main() {
	age := 27
	fmt.Println(age)
	age = 28
	fmt.Println(age)
	fmt.Println("我是土木转行的，今天第2天")
	fmt.Println(add(3, 5))
}
