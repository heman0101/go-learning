package homework1017

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Statement 获取一言语录
func statement() (string, error) {
	const api = "https://api.ixiaowai.cn/ylapi/index.php"
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

func TestStatement(t *testing.T) {
	result, _ := statement()
	if len(result) == 0 {
		t.Fatalf("获取错误")
	}
}
