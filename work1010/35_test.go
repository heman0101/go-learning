// 搜索插入位置
package work1010

import (
	"fmt"
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
	d := []int{1, 3, 5, 6}
	target := 5

	l := searchInsert(d, target)
	fmt.Printf("位置: %v\n", l)

/* 	data :=[]struct{
		nums  string
		target string  
	}{
		{nums: "1356",target: "5"},
		{nums: "1356",target: "1"},
		{nums: "1356",target: "8"},
	}

	for _, v := range data {
		
		l :=searchInsert()
		fmt.Printf("l: %v\n", l)
	} */
}