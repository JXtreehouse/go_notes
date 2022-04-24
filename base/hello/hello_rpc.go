/*
 * @Author: your name
 * @Date: 2022-04-24 15:47:21
 * @LastEditTime: 2022-04-24 16:03:50
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/base/hello/hello_rpc.go

 package hello
 import (
	 "net"
	 "log"
 )
 */
// 我们先构造一个HelloService类型，其中Hello()方法用于实现打印功能:

type HelloService struct {}

// Hello方法必须满足go语言中的RPC规则：
//  f方法只能有两个可序列化的参数
// 其中第二参数是指针类型，并且返回一个error类型，同时必须是公共的方法
func (p *HelloService) Hello (request string ,  reply *string) error {
	*reply =  "hello: " + request
	 return nil
}


// 将HelloService类型的对象注册为一个RPC服务
func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err !=nil {
		log.Fatal("ListenTCP error:" , err)
	}
	conn, err := listener.Accept()
	
}