package main

import "fmt"

func main5() {

	ret := plusOne([]int{1, 9, 9})
	fmt.Println(ret)

}

func plusOne(digits []int) []int {
	temp := []int{}
	jin := 0
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + jin
		if i == len(digits)-1 {
			num++
		}
		if num > 9 {
			num = num - 10
			jin = 1
		} else {
			jin = 0
		}

		temp = append(temp, num)
	}
	if jin > 0 {
		temp = append(temp, jin)
	}

	ret := []int{}
	for i := len(temp) - 1; i >= 0; i-- {
		ret = append(ret, temp[i])
	}
	return ret

}
