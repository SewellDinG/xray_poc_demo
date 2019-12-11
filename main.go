package main

import (
	"fmt"
)

func main() {
	target := "https://postman-echo.com"
	var poc = &Poc{}
	poc = poc.GetPoc("./test.yaml")
	//fmt.Printf("%#v\n", poc)
	for _, rule := range poc.Rules {
		fmt.Println("**********************************")
		fmt.Println(rule)
		// 初始化
		url := target + rule.Path
		ruleHeaders := rule.Headers
		//fmt.Printf("%#v\n",ruleHeaders)
		ruleBody := rule.Body
		// 判断请求类型
		switch rule.Method {
		case "GET":
			//HttpGET(url)
			HttpGen(url, "GET", "", Headers{})
		case "POST":
			//HttpPOST(url)
			HttpGen(url, "POST", ruleBody, ruleHeaders)
		case "PUT":
			HttpGen(url, "PUT", "", Headers{})
		case "MOVE":
			HttpGen(url, "MOVE", "", Headers{})
		default:
			fmt.Println("Default...")
		}
	}
}
