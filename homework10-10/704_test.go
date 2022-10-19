package homework1010

import (
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
	
	data := []struct {
		nums   []int
		target int
		answer int
	}{
		{nums: []int{-1,0,3,5,9,12}, target: 5, answer:3},
		{nums: []int{-1,0,3,5,9,12}, target: 6, answer:-1},
		{nums: []int{-1,0,3,5,9,12}, target: 12, answer: 5},
	}

	for _, v := range data {
		n, tar  := v.nums, v.target
		r := search(n, tar)
		if r != v.answer {
			t.Fatalf("expect:[%v] != result[%v]", r,v.answer)
		}
}
}