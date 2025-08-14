package main

import "fmt"

func main7() {
	ret := longestCommonPrefix([]string{"aba", "abb", "abc"})
	fmt.Println(ret)
}

func longestCommonPrefix(strs []string) string {
	ret := ""
	counter := 0
	for {
		temp := ""
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == 0 {
				return ret
			}
			if len(strs[i]) <= counter {
				return ret
			}
			if temp == "" {
				temp = string(strs[i][counter])
			} else if temp != string(strs[i][counter]) {
				return ret
			}

		}
		ret += temp
		counter++
	}
}
