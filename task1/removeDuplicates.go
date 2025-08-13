package main

import "fmt"

func main3() {

	ret := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(ret)

}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for {
		lenth := len(nums)
		if index == lenth-1 {
			break
		}
		if nums[index] == nums[index+1] {
			nums = append(nums[:index+1], nums[index+2:]...)
		} else {
			index++
		}
	}

	return len(nums)
}
