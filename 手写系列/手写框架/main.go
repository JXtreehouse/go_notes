/*
 * @Author: your name
 * @Date: 2021-04-16 17:44:53
 * @LastEditTime: 2021-04-16 18:07:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/Web框架/goturing/main.go
 */

 package main
 
 import (
	 "fmt"
	 "net/http"
	 "turingGo"
 )

 func main() {
	 // 使用New()创建 turingGo的实例
	 r: = turingGo.New()
	 // 使用 GET()方法添加路由
	 // 这里的路由，只是静态路由，不支持/hello/:name这样的动态路由，动态路由我们将在下一次实现。
	 r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		 fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	 })

	 r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		 for k, v := range req.Header{
			 fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		 }
	 })
	
	 // 使用Run()启动Web服务
	 r.Run(:"7001")
 }