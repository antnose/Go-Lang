package main

import "fmt"

// slice -> dynamic
// most used construct in go
// + useful method
func main() {
	// it's uninitialized slice is nil. We are use null in other language but in go we are use nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	var nums = make([]int, 0, 5)
	fmt.Println(cap(nums))
	// fmt.Println(nums)
	// fmt.Println(nums == nil)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 7)
	fmt.Println(nums)
	fmt.Println(cap(nums))
}
