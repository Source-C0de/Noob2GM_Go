// 2270. Number of Ways to Split Array
// https://leetcode.com/problems/number-of-ways-to-split-array/description/?envType=daily-question&envId=2025-01-03


package main

import "fmt"

func waysToSplitArray(nums []int) int {
	var sumLeft int64 = 0
	var sumRight int64 = 0
	var sumRightNew int64
	cnt := 0

	for i := 0; i < len(nums); i++ {
		sumRight += int64(nums[i])
	}
	for i := 0; i < len(nums)-1; i++ {
		sumLeft += int64(nums[i])
		sumRightNew = sumRight - sumLeft
		fmt.Println(sumLeft, sumRightNew)
		if sumLeft >= sumRightNew {
			cnt++
		}
	}
	return cnt
}
