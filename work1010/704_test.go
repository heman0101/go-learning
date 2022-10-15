package work1010

import (
	"fmt"
	"testing"
)

func search(nums []int, target int) int {
	var left, right = 0, len(nums) - 1 //初始化

	for left <= right {
		// mid := (left +right)/2        太大会溢出
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target { //遍历右半边
			left = mid + 1
		} else if nums[mid] > target { //遍历左半边  
			right = mid - 1
		}
	}
	return -1
}

func TestSearch(t *testing.T) {
	d :=[]int{-1,0,3,5,9,12}
	target :=2

	l :=search(d,target)
	fmt.Printf("位置: %v\n", l)
}