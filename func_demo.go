package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"reflect"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 结构体struct转map
func StructToMap(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	// 若传入指针类型，编译异常：panic: reflect: NumField of non-struct type
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).String()
	}
	return data
}

// 判断s1是否包含s2，返回bool类型结果
func Contains(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}

// 判断一个b1是否包含b2，返回bool类型结果。与contains不同的是，bcontains是字节流（bytes）的查找
func Bcontains(s1, s2 []byte) bool {
	return strings.Contains(string(s1), string(s2))
}

// 使用正则表达式s1来匹配s2，返回bool类型匹配结果
func Matches(re, s string) bool {
	return true
}

// 使用正则表达式s1来匹配b1，返回bool类型匹配结果。与matches不同的是，bmatches匹配的是字节流（bytes）
func Bmatches(re, s string) bool {
	return true
}

// 判断s1是否由s2开头
func StartsWith(s1, s2 string) bool {
	return strings.HasPrefix(s1, s2)
}

// 判断s1是否由s2结尾
func EndsWith(s1, s2 string) bool {
	return strings.HasSuffix(s1, s2)
}

// map 中是否包含某个 key，目前只有 headers 是 map 类型
func In(s string, m map[string]string) bool {
	if _, ok := m[s]; ok {
		return true
	} else {
		return false
	}
}

// 字符串的 md5
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// 两个范围内的随机数
func RandomInt(from, to int) int {
	return rand.Intn(to) + from
}

// 指定长度的小写字母组成的随机字符串
func RandomLowercase(n int) string {
	lowercase := []rune("abcdefghijklmnopqrstuvwxyz")
	randomLowercase := make([]rune, n)
	for i := range randomLowercase {
		randomLowercase[i] = lowercase[rand.Intn(len(lowercase))]
	}
	return string(randomLowercase)
}

// 将字符串或 bytes 进行 base64 编码
func Base64(i interface{}) string {
	v := reflect.ValueOf(i).String()
	return base64.StdEncoding.EncodeToString([]byte(v))
	//switch i.(type){
	//case string:
	//	return base64.StdEncoding.EncodeToString([]byte(v))
	//case []byte:
	//	return base64.StdEncoding.EncodeToString(v)
	//}
}

// 将字符串或 bytes 进行 base64 解码
func Base64Decode(i interface{}) string {
	v := reflect.ValueOf(i).String()
	decodeBytes, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		log.Fatalf("base64Decode err:", err)
	}
	return string(decodeBytes)
}

// 将字符串或 bytes 进行 urlencode 编码
func UrlEncode(i interface{}) string {
	v := reflect.ValueOf(i).String()
	return url.QueryEscape(v)
}

// 将字符串或 bytes 进行 urldecode 解码
func UrlDecode(i interface{}) string {
	v := reflect.ValueOf(i).String()
	decodeString, err := url.QueryUnescape(v)
	if err != nil {
		log.Fatalf("urlDecode err:", err)
	}
	return decodeString
}
