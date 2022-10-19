package homework1010

import (
	"testing"
)

func reverseString(s []byte) {
	str := []rune(string(s))     //type rune = int32  处理中文

	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
		// s[i] = s[len(s)-i-1]     //这种前半部分已经覆盖了，后面再和前面换就相当于换自己
	}

	strbyte := []byte(string(str))
	for i := 0; i < len(strbyte); i++ {
		s[i] = strbyte[i]
	}
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
		tmp := []byte(v.val) //字节切片
		reverseString(tmp[:])
		if string(tmp) != v.answer { //字符串
			t.Fatalf("expect:[%s] != result[%s]", string(tmp), v.answer)
		}
	}
}