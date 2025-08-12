package main

import "fmt"

func main() {

	ret := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(ret)

	ret = twoSum([]int{3, 2, 4}, 6)
	fmt.Println(ret)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result := []int{i, j}
				return result
			}
		}
	}
	return []int{-1, -1}
}
