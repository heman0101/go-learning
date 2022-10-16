package homework1013

import (
	"reflect"
	"testing"
)

// sortedSquares有序数组的平方
// 双指针法 
// 数组平方的最大值就在数组的两端，不是最左边就是最右边，不可能是中间。  
func sortedSquares(nums []int) []int {
	var i, j, k int = 0, len(nums) - 1, len(nums) - 1
	r := make([]int, len(nums))  //定义一个新数组r，和原数组一样的大小

	for i <= j {
		ln, rn := nums[i]*nums[i], nums[j]*nums[j]
		if ln >= rn {
			r[k] = ln
			i++
		} 
		if ln < rn {
			r[k] = rn
			j--
		}
		k--
	}
	return r
}

func TestSortedSquares(t *testing.T) {
	data := []struct {
		nums   []int
		answer []int
	}{
		{nums: []int{-4, -1, 0, 3, 10}, answer: []int{0, 1, 9, 16, 100}},
		{nums: []int{-7, -3, 2, 3, 11}, answer: []int{4, 9, 9, 49, 121}},
		{nums: []int{0, 1, 2, 3, 10}, answer: []int{0, 1, 4, 9, 100}},
	}

	for _, v := range data {
		result := sortedSquares(v.nums)
		if !reflect.DeepEqual(result,v.answer) {      //reflect.DeepEqual函数判断切片是否相等
			t.Fatalf("expect:[%v] != result[%v]", result, v.answer)
		}
	}
}
