package homework1013

import (
	"sort"     //导入sort 包
	"testing"
)


//containsDuplicate函数判断是否存在重复元素  
// 先对元素进行排序，然后遍历是否存在相邻一样的元素
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)         //sort.Ints函数对切片从小到大排序 
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	data := []struct {
		nums   []int
		answer bool
	}{
		{nums: []int{1, 2, 3, 1}, answer: true},
		{nums: []int{1, 2, 3, 4}, answer: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, answer: true},
	}

	for _, v := range data {
		result := containsDuplicate(v.nums)
		if result !=v.answer {      
			t.Fatalf("expect:[%v] != result[%v]", result, v.answer)
		}
	}
}
