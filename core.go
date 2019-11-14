package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/koron/go-dproxy"
)

const (
	KEY     = "1945325576"
	KEYFROM = "Youdao-dict-v21"
	URL     = "http://fanyi.youdao.com/openapi.do?keyfrom=%s&key=%s&type=data&doctype=json&version=1.1&q=%s"
	VERSION = "0.1.0"
)
const (
        InfoColor    = "\033[1;34m%s\033[0m"
        NoticeColor  = "\033[1;36m%s\033[0m"
        WarningColor = "\033[1;33m%s\033[0m"
        ErrorColor   = "\033[1;31m%s\033[0m"
        DebugColor   = "\033[0;36m%s\033[0m"
)

// 打印版本号
func ShowVersion() {
	fmt.Printf("youdao-go %s\n", VERSION)
}

// 请求连接并返回响应
func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

// 处理返回的 json 数据
func Parser(words string) {
	// 翻译的内容需要进行 urlencode 编码
	escapeeWords := url.QueryEscape(words)
	queryUrl := fmt.Sprintf(URL, KEYFROM, KEY, escapeeWords)

	resp, err := Request(queryUrl)

	var v interface{}
	err = json.Unmarshal([]byte(resp), &v)
	if err != nil {
		return
	}

	// 使用 go-dproxy 库解析 json 数据
	p := dproxy.New(v)
	errorCode, err := p.M("errorCode").Int64()
	if errorCode != 0 || err != nil {
		fmt.Println(errorCode)
		fmt.Println(">>  Exception: The words can't be found, please check your spelling")
		return
	}
	// 词意翻译
	translation, _ := p.M("translation").A(0).String()
	translation = fmt.Sprintf(ErrorColor, translation)
	fmt.Printf(">>  %s: %s\n\n", words, translation)

	// 发音
	usPhonetic, err := p.M("basic").M("us-phonetic").String()
	if err == nil {
		ukPhonetic, _ := p.M("basic").M("uk-phonetic").String()
        usPhonetic = fmt.Sprintf(DebugColor, usPhonetic)
        ukPhonetic = fmt.Sprintf(DebugColor, ukPhonetic)
		fmt.Printf("    美:[%s]  英:[%s]\n\n", usPhonetic, ukPhonetic)
	}

	// 基本释义
	explains, err := p.M("basic").M("explains").Array()
	if err == nil {
		for _, value := range explains {
			value = fmt.Sprintf(WarningColor,value)
			fmt.Printf("    %s\n", value)
		}
		fmt.Println()
	}

	// 网络释义
	web, err := p.M("web").Array()
	if err == nil {
		for _, value := range web {
			item := value.(map[string]interface{})
			v := fmt.Sprintf(InfoColor,item["value"])
			fmt.Printf("    %s %s\n", item["key"], v)
		}
	}
}
