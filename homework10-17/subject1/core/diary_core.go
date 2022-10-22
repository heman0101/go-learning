package core

import (
	"fmt"
	"io"
	"net/http"
)

// Statement 获取一言语录
func Statement() (string, error) {
	const api = "https://api.ixiaowai.cn/ylapi/index.php"
	resp, err := http.Get(api)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(resp.Status)
	}
	b, _ := io.ReadAll(resp.Body)
	return string(b), resp.Body.Close()
}

// Diary 获取舔狗日记
func Diary() (string, error) {
	const api = "https://api.ixiaowai.cn/tgrj/index.php"
	resp, err := http.Get(api)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf(resp.Status)
	}
	b, _ := io.ReadAll(resp.Body)
	return string(b), resp.Body.Close()
}
