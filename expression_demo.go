package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Expr struct {
	Body        string `json:"body"`
	ContentType string `json:"content_type"`
	Status      string `json:"status"`
}

func splitFun(r rune) bool {
	return r == '.' || r == '='
}

func ExprHandler(expr string) *Expr {
	m := make(map[string]string)
	expSliceStep1 := strings.Split(expr, "&&")
	for _, v := range expSliceStep1 {
		v = strings.TrimSpace(v)
		//fmt.Println("|--------", v)
		expSliceStep2 := strings.FieldsFunc(v, splitFun)
		//for _, vv := range expSliceStep2 {
		//	fmt.Println("   |*****", strings.TrimSpace(vv))
		//}
		// 未使用resp字段
		m[strings.TrimSpace(expSliceStep2[1])] = strings.TrimSpace(expSliceStep2[2])
	}
	// map -> json
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("json.Marshal err:", err)
	}
	// json -> struct
	var exprHandled *Expr
	err = json.Unmarshal(data, &exprHandled)
	if err != nil {
		log.Fatalf("json.Unmarshal err:", err)
	}
	return exprHandled
}

// 模式匹配
func ExprMatcher(resp *Response, expr *Expr) bool {
	fmt.Printf("* resp: %+v %+v\n", resp.Status, resp.Headers)
	fmt.Printf("* expr: %+v\n", expr)
	// 处理的太冗余了...
	// header
	respHeader := fmt.Sprint(resp.Headers)
	exprContentType := strings.Split(expr.ContentType, "'")
	var exprContent string
	if len(exprContentType) > 1 {
		exprContent = exprContentType[1]
	}
	// body
	exprBodyList := strings.Split(expr.Body, "'")
	var exprBody string
	if len(exprBodyList) > 1 {
		exprBody = exprBodyList[1]
	}
	if expr.Status != "" && strconv.Itoa(resp.Status) != expr.Status {
		return false
	} else if !Bcontains([]byte(respHeader), []byte(exprContent)) {
		return false
	} else if !Bcontains(resp.Body, []byte(exprBody)) {
		return false
	}
	return true
}
