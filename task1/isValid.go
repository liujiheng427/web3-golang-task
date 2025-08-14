package main

import "fmt"

func main() {
	ret := isValid("([])")
	fmt.Println(ret)

}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	arr := []string{}
	for _, i := range s {
		v := string(i)
		if v == "(" {
			arr = append(arr, "(")
		} else if v == ")" {
			if len(arr) == 0 {
				return false
			}
			if arr[len(arr)-1] != "(" {
				return false
			} else {
				arr = arr[:len(arr)-1]
			}
		} else if v == "[" {
			arr = append(arr, "[")
		} else if v == "]" {
			if len(arr) == 0 {
				return false
			}
			if arr[len(arr)-1] != "[" {
				return false
			} else {
				arr = arr[:len(arr)-1]
			}
		} else if v == "{" {
			arr = append(arr, "{")
		} else if v == "}" {
			if len(arr) == 0 {
				return false
			}
			if arr[len(arr)-1] != "{" {
				return false
			} else {
				arr = arr[:len(arr)-1]
			}
		}
	}
	if len(arr) > 0 {
		return false
	}
	return true
}
