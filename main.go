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
	nums := []int{1, 2, 3}
	fmt.Println(nums)

	fmt.Println(nums[0])
	fmt.Println(len(nums))
	nums = append(nums, 4)
	fmt.Println(nums)
	for i, v := range nums {
		fmt.Println(i, v)
	}
	ages := map[string]int{"Alice": 30}
	fmt.Println(ages)
	ages["Bob"] = 27
	fmt.Println(ages["Bob"])
}
