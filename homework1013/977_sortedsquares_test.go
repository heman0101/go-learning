package homework1013

import "testing"       

func sortedSquares(nums []int) []int {
	var i, j, k int = 0, len(nums) - 1, len(nums) - 1
	r := make([]int, len(nums))

	for i <= j {
		ln, rn := nums[i]*nums[i], nums[j]*nums[j]
		if ln >= rn {
			r[k] = ln
			i++
		} else {
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
		
		// sortedSquares(nums)
		if sortedSquares(nums) != v.answer {
			t.Fatalf("expect:[%v] != result[%v]", sortedSquares(nums), v.answer)
		}
	}
}
