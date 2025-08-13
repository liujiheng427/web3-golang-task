package main

import (
	"fmt"
	"sort"
)

func main4() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	ret := merge(intervals)
	fmt.Println(ret)

}

func merge(intervals [][]int) [][]int {
	var ret [][]int
	var hasArr [][]int

	sort.Slice(intervals, func(left, right int) bool {
		return intervals[left][0] < intervals[right][0]
	})
	fmt.Println(intervals)

	lenth := len(intervals)
	for i := 0; i < lenth; i++ {
		// 判断intervals[i] 是否存在hasArr里面
		has := false
		for _, v := range hasArr {
			if v[0] == intervals[i][0] && v[1] == intervals[i][1] {
				has = true
				break
			}
		}
		if has {
			continue
		}
		for j := i + 1; j < lenth; j++ {
			if intervals[i][0] <= intervals[j][0] && intervals[i][1] >= intervals[j][1] {
				fmt.Println(1)
				hasArr = append(hasArr, intervals[j])
				continue
			}

			if intervals[i][0] >= intervals[j][0] && intervals[i][1] <= intervals[j][1] {
				fmt.Println(2)
				hasArr = append(hasArr, intervals[j])
				intervals[i][0] = intervals[j][0]
				intervals[i][1] = intervals[j][1]
				continue
			}

			if intervals[i][1] <= intervals[j][1] && intervals[i][1] >= intervals[j][0] {
				fmt.Println(3)
				hasArr = append(hasArr, intervals[j])
				intervals[i][1] = intervals[j][1]
				continue
			}

			if intervals[j][1] <= intervals[i][1] && intervals[j][1] >= intervals[i][0] {
				fmt.Println(4)
				hasArr = append(hasArr, intervals[j])
				intervals[i][0] = intervals[j][0]
				continue
			}
		}
		tmp := make([]int, 2)
		copy(tmp, intervals[i])
		ret = append(ret, tmp)
		fmt.Println(hasArr, ret)
	}

	return ret
}
