package subject2

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestSendWeChat(t *testing.T) {
	//从文件读取sendkey
	path := "E:\\Desktop\\data.txt"
	fi, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	secert, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
	}

	data := []struct {
		title string
		desp  string
		// secert string
		answer1 string
		answer2 string
	}{
		{title: "第一次", desp: "测试", answer1: "SUCCESS", answer2: "ok"},
		// {title: "第二次", desp: "测试", answer1: "SUCCESS", answer2: "ok"},
		// {title: "第三次", desp: "测试", answer1: "SUCCESS", answer2: "ok"},
	}
	for _, v := range data {
		pushid, readkey, result1 := Refer(v.title, v.desp,string(secert))

		if result1 != v.answer1 {
			t.Fatalf("expect:[%v] != result[%v]", result1, v.answer1)
		}

		time.Sleep(time.Second) //等待1s,防止查询不到推送状态

		result2 := Consult(pushid, readkey)
		if result2 != v.answer2 {
			t.Fatalf("expect:[%v] != result[%v]", result2, v.answer2)
		}
	}
}
