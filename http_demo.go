package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func HttpGET(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func HttpPOST(url string) {
	contentType := "application/x-www-form-urlencoded"
	contentBody := strings.NewReader("a=1")
	resp, err := http.Post(url, contentType, contentBody)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

// General func
func HttpGen(url, httpMethod, ruleBody string, ruleHeaders Headers) {
	client := &http.Client{}
	contentBody := strings.NewReader(ruleBody)
	req, err := http.NewRequest(httpMethod, url, contentBody)
	if err != nil {
		log.Fatal(err)
	}
	// 设置请求头
	// Headers struct 空判断
	if !reflect.DeepEqual(ruleHeaders, Headers{}) {
		ruleHeader := StructToMap(ruleHeaders)
		for contentTypeKey, contentTypeValue := range ruleHeader {
			//fmt.Println(contentTypeKey, "--------", contentTypeValue)
			req.Header.Set(contentTypeKey, contentTypeValue)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
