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

	// var nums = make([]int, 0, 5)
	// fmt.Println(cap(nums))
	// fmt.Println(nums)
	// fmt.Println(nums == nil)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6)
	// nums = append(nums, 7)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// nums := []int{}
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// fmt.Println(nums)
	// fmt.Println(len(nums))

	// var nums = make([]int, 1, 5)
	// nums[0] = 1
	// fmt.Println(nums)

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 1)
	// var nums2 = make([]int, len(nums))
	// fmt.Println(len(nums2))

	// fmt.Println(nums, nums2)

	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// slice operator
	// var nums = []int{1, 2, 3, 4}
	// fmt.Println(nums[0:1])
	// fmt.Println(nums[0:])
	// fmt.Println(nums[:2])

	// slice
	// var nums1 = []int{1, 2}
	// var nums2 = []int{1, 2}
	// fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{1, 2, 3, 4}, {4, 5, 6}}
	fmt.Println(nums)

}
