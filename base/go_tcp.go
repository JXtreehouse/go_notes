/*
 * @Author: your name
 * @Date: 2021-07-02 16:58:40
 * @LastEditTime: 2021-07-02 18:38:52
 * @LastEditors: Please set LastEditors
 * @Description: https://laravelacademy.org/post/20953#toc-tcp-%E7%A4%BA%E4%BE%8B%E7%A8%8B%E5%BA%8F
 * @FilePath: /go_notes/base/go_tcp.go
 */

 package main

 import (
	 "bytes"
	 "fmt"
	 "io"
	 "net"
	 "os"
 )

 func main () {
	//  fmt.Println(os.Args[1])
	 if len(os.Args) != 2 {
		 fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		 // Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。
		 // 
		 os.Exit(1)
	 }

	 // 从参数中读取主机信息
	 service := os.Args[1]

	 // 建立网络连接
	 conn, err := net.Dial("tcp", service)
	//  fmt.Println(conn)
	 // 连接出错则打印错误信息并退出程序
	 checkError(err)

	 // 调用返回的连接对象提供的write 方法发送请求
	 _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	 checkError(err)

	 // 通过连接对象提供的Read方法读取所有响应数据
	 result, err := readFully(conn)
	 checkError(err)

	 // 打印响应数据
	 fmt.Println(string(result))

	 os.Exit(0)
 }

 func checkError(err error) {
	 if err != nil {
		 fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		 os.Exit(1)
	 }
 }

 func readFully(conn net.Conn) ([]byte, error) {
	 // 读取所有响应数据后主动关闭连接
	 defer conn.Close() 

	 result:= bytes.NewBuffer(nil)
	 var buf[512]byte
	 for {
		 n, err := conn.Read(buf[0:])
		 result.Write(buf[0:n])
		 if err !=nil {
			 // 它用于表示输入流的结尾. 因为每个文件都有一个结尾, 所以io.EOF很多时候并不能算是一个错误, 它更重要的是一个表示输入流结束了.
			 if err == io.EOF{
				 break
			 }
			 return nil, err
		 }
	 }
	 return result.Bytes(), nil
 }