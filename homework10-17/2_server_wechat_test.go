package homework1017

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

// 在Go中Http请求的返回结果为 *http.Response 类型，response.Body 类型为 io.Reader，把请求结果转化为Map
func Transformation(response *http.Response) map[string]interface{} {
	var result map[string]interface{}
	body, err := io.ReadAll(response.Body)
	if err == nil {
		json.Unmarshal([]byte(string(body)), &result)
	}
	return result
}

// 推送消息
func refer(title, desp string) (string, string, string) {
	api := "https://sctapi.ftqq.com/"

	//从文件读取sendkey
	path := "E:\\Desktop\\data.txt"
	fi, _ := os.Open(path)
	str, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Fatal(err)
	}

	resp, _ := http.Get(api + string(str) + ".send?title=" + title + "&desp=" + desp)
	//调用Transformation函数进行返回值转换
	result := Transformation(resp)

	//读取result中的pushid, readkey, error
	data := result["data"]
	t := data.(map[string]interface{})
	pushid := t["pushid"]
	readkey := t["readkey"]
	error := t["error"]

	return pushid.(string), readkey.(string), error.(string)
}

// 查询推送状态
func consult(pushid, readkey string) string {
	api := "https://sctapi.ftqq.com/"
	resp, _ := http.Get(api + "push?id=" + pushid + "&readkey=" + readkey)

	//读取wxstatus
	result := Transformation(resp)
	data := result["data"]
	t := data.(map[string]interface{})
	wxstatus := t["wxstatus"]

	//读取errmsg
	wxstatus_map := make(map[string]interface{})
	err := json.Unmarshal([]byte(wxstatus.(string)), &wxstatus_map)
	if err != nil {
		panic(err)
	}
	errmsg := wxstatus_map["errmsg"]
	return errmsg.(string)
}

// 测试函数
func TestPushWeChat(t *testing.T) {
	data := []struct {
		title   string
		desp    string
		answer1 string
		answer2 string
	}{
		{title: "第一次", desp: "测试", answer1: "SUCCESS", answer2: "ok"},
		{title: "第二次", desp: "测试", answer1: "SUCCESS", answer2: "ok"},
		{title: "第三次", desp: "测试", answer1: "SUCCESS", answer2: "ok"},
	}
	for _, v := range data {
		pushid, readkey, result1 := refer(v.title, v.desp)

		if result1 != v.answer1 {
			t.Fatalf("expect:[%v] != result[%v]", result1, v.answer1)
		}

		time.Sleep(time.Second) //等待1s,防止查询不到推送状态

		result2 := consult(pushid, readkey)
		if result2 != v.answer2 {
			t.Fatalf("expect:[%v] != result[%v]", result2, v.answer2)
		}
	}
}
