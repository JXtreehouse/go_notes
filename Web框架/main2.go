/*
 * @Author: your name
 * @Date: 2021-04-16 17:02:44
 * @LastEditTime: 2021-04-16 17:56:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/Web框架/main2.go
 */

 package main
 
 import (
	 "fmt"
	 "log"
	 "net/http"
 )

 // engine is the uni handler for all requests
 // 定义个空结构体
 type Engine struct{}

 // 实现方法ServeHTTP
 // 这个方法有2个参数，第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息；第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。
 func (engin *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	 switch req.URL.Path {
	 case "/":
		fmt.Fprintf(w, "URL.Path")
	case "/hello":
		for k,v := range req.Header{
			fmt.Fprintf(w, "Header[%q]\n", k,v)
		}	
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)	 
	 }
 }

 func main() {
	 engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9998", engine))
 }
