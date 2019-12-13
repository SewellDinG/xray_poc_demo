package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Response struct {
	Url     string
	Status  int
	Body    []byte
	Headers http.Header
}

func HttpReqer(url, httpMethod, ruleBody string, ruleHeaders Headers) *Response {
	// goroutine数量-1
	//defer wg.Done()
	client := &http.Client{}
	contentBody := strings.NewReader(ruleBody)
	req, err := http.NewRequest(httpMethod, url, contentBody)
	if err != nil {
		log.Fatalf("http.NewRequest err:", err)
	}
	// 设置请求头
	// Headers struct 空判断
	if !reflect.DeepEqual(ruleHeaders, Headers{}) {
		ruleHeader := StructToMap(ruleHeaders)
		for contentTypeKey, contentTypeValue := range ruleHeader {
			//fmt.Println(contentTypeKey, "--------", contentTypeValue)
			if contentTypeKey == "ContentType"{
				contentTypeKey = "Content-Type"
			}
			req.Header.Set(contentTypeKey, contentTypeValue)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("client.Do err:", err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll err:", err)
	}
	var respStruct = &Response{url, resp.StatusCode, body, resp.Header}
	return respStruct
}
