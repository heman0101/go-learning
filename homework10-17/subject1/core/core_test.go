package core

import (
	"testing"
)

func TestStatement(t *testing.T) {
	result, err := Statement()
	if err != nil {
		panic(err)
	}
	if result == "" {
		t.Fatalf("获取失败！")
	}
}

func TestDiary(t *testing.T) {
	result, err := Statement()
	if err != nil {
		panic(err)
	}
	if result == "" {
		t.Fatalf("获取失败！")
	}
}
