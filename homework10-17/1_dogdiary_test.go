package homework1017

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

// DogDiary 获取舔狗日记
func dogdiary() (string, error) {
	const api = "https://api.ixiaowai.cn/tgrj/index.php"
	resp, err := http.Get(api)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(resp.Status)
	}
	b, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return string(b), err
}

// 判断获取的结果是否为空
func TestDogDiary(t *testing.T) {
	result, _ := dogdiary()
	if len(result) == 0 {
		t.Fatalf("获取错误")
	}
}
