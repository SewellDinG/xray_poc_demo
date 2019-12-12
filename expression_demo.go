package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		fmt.Println("|--------", v)
		expSliceStep2 := strings.FieldsFunc(v, splitFun)
		for _, vv := range expSliceStep2 {
			fmt.Println("   |*****", strings.TrimSpace(vv))
		}
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

// 模式匹配???
func ExprMatcher(resp *Response, expr *Expr) bool {
	fmt.Printf("%+v\n", resp)
	fmt.Printf("%+v\n", expr)
	return false
}
