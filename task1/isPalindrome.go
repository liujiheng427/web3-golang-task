package main

import (
	"fmt"
)

func main6() {
	ret := isPalindrome(123454321)
	fmt.Println(ret)

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arr := []int{}
	for {
		if x == 0 {
			break
		}
		temp := x % 10
		arr = append(arr, temp)
		x = x / 10
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}
