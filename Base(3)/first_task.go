package main

import "fmt"

// We have an integer array nums, print true if any value
// occurs at least twice in the array, and false if each element
// different.

func main() {
	nums := []int{1, 2, 3, 4, 5, 2}
	nums2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sliceRepetition(nums))
	fmt.Println(sliceRepetition(nums2))
}

func sliceRepetition(nums []int) bool {
	var nums2 []int
	for k, v := range nums {
		nums2 = nums[k+1:]
		for _, j := range nums2 {
			if v == j {
				return true
			}
		}
	}
	return false
}
