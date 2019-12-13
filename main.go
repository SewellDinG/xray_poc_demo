package main

import (
	"flag"
	"fmt"
	_ "sync"
)

//var wg sync.WaitGroup

var targetURL = flag.String("url", "https://postman-echo.com", "Input Target URL")
var pocPath = flag.String("poc", "test.yaml", "Input POC file")

func main() {
	// 初始化命令行参数
	flag.Parse()
	target := *targetURL
	var poc = &Poc{}
	poc = poc.PocGetter(*pocPath)
	// 设置goroutine数
	//wg.Add(len(poc.Rules))
	//go func() {
	func() {
		for i, rule := range poc.Rules {
			fmt.Println("------------------", i+1, "------------------")
			// 初始化rule
			url := target + rule.Path
			ruleHeaders := rule.Headers
			ruleBody := rule.Body
			ruleExpr := rule.Expression
			fmt.Printf("* rule:%+v\n", rule)
			// 判断请求类型
			var resp *Response
			switch rule.Method {
			case "GET":
				resp = HttpReqer(url, "GET", "", Headers{})
			case "POST":
				resp = HttpReqer(url, "POST", ruleBody, ruleHeaders)
			case "PUT":
				resp = HttpReqer(url, "PUT", "", Headers{})
			case "MOVE":
				resp = HttpReqer(url, "MOVE", "", Headers{})
			default:
				fmt.Println("No allowed method...")
			}
			// 匹配resp
			ruleExprHandled := ExprHandler(ruleExpr)
			ruleExprMatchResult := ExprMatcher(resp, ruleExprHandled)
			if ruleExprMatchResult {
				fmt.Println("* Vuln:", poc.Name, "\n* Target:", target)
			} else {
				fmt.Println("* Failure...")
			}
		}
	}()
	//wg.Wait()
}
