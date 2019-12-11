package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

var wg sync.WaitGroup

var targetURL = flag.String("target", "https://postman-echo.com", "Input Target URL")
var pocPath = flag.String("poc", "test.yaml", "Input POC file")

type request struct {
	url          urlType
	method       string
	body         []byte
	headers      map[string]string
	content_type string
}
type response struct {
	url          urlType
	status       int
	body         []byte
	headers      map[string]string
	content_type string
}
type urlType struct {
	scheme   string
	domain   string
	host     string
	port     string
	path     string
	query    string
	fragment string
}

func HttpReq(url, httpMethod, ruleBody string, ruleHeaders Headers) *http.Response {
	// goroutine数量-1
	defer wg.Done()
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
	fmt.Println(string(body))
	return resp
}

func main() {
	// 初始化命令行参数
	flag.Parse()
	target := *targetURL
	var poc = &Poc{}
	poc = poc.GetPoc(*pocPath)
	// 设置goroutine数
	wg.Add(len(poc.Rules))
	go func() {
		for _, rule := range poc.Rules {
			fmt.Println("**********************************\n", rule)
			// 初始化rule
			url := target + rule.Path
			ruleHeaders := rule.Headers
			ruleBody := rule.Body
			ruleExpression := rule.Expression
			fmt.Println(ruleExpression)
			// 判断请求类型
			switch rule.Method {
			case "GET":
				resp := HttpReq(url, "GET", "", Headers{})
				fmt.Println(resp.Status)
			case "POST":
				resp := HttpReq(url, "POST", ruleBody, ruleHeaders)
				fmt.Println(resp.Status)
			case "PUT":
				resp := HttpReq(url, "PUT", "", Headers{})
				fmt.Println(resp.Status)
			case "MOVE":
				resp := HttpReq(url, "MOVE", "", Headers{})
				fmt.Println(resp.Status)
			default:
				fmt.Println("No allowed method...")
			}
			// 匹配resp

		}
	}()
	wg.Wait()
}
