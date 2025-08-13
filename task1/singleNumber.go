package main

import "fmt"

func main2() {

	ret := singleNumber([]int{7, 2, 7, 11, 11})
	fmt.Println(ret)

	ret = singleNumber([]int{3, 2, 2})
	fmt.Println(ret)
}

func singleNumber(nums []int) int {

	tempMap := map[int]int{}
	for _, v := range nums {
		tempMap[v]++
	}

	for k, v := range tempMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
