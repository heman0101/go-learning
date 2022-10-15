// 反转字符串
package work1010

import (
	"fmt"
	"testing"
)

// reverseString 反转字符串 ( 需要考虑到中文用例 )
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
		// s[i] = s[len(s)-i-1]     //这种前半部分已经覆盖了，后面再和前面换就相当于换自己了
	}
	fmt.Printf("s: %v\n", s)
}

func TestReverseString(t *testing.T) {
	// 表格驱动测试
		data := []struct {
			val    string
			answer string
		}{
			{val: "abcd", answer: "dcba"},
			{val: "5123", answer: "3215"},
			{val: "你好世界", answer: "界世好你"},
		}

		for _, v := range data {
	        tmp :=[]byte(v.val)
			reverseString(tmp[:])
			if string(tmp) != v.answer {
				t.Fatalf("expect:[%s] != result[%s]", string(tmp), v.answer)
			}
		}
}