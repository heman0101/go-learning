package homework1010

import (
	"testing"
)

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] { //target小于或者等于遍历的数时，索引或者插入的位置就是i
			return i
		}
	}
	return len(nums) //大于所有的数，插入最后
}

func TestSearchInsert(t *testing.T) {


	data := []struct {
		nums   []int
		target int
		answer int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, answer: 2},
		{nums: []int{1, 3, 5, 6}, target: 1, answer: 0},
		{nums: []int{1, 3, 5, 6}, target: 8, answer: 4},
	}

	for _, v := range data {
		n, tar  := v.nums, v.target
		r := searchInsert(n, tar)
		if r != v.answer {
			t.Fatalf("expect:[%v] != result[%v]", r,v.answer)
		}
	}
}