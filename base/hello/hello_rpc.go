/*
 * @Author: AlexZ33
 * @Date: 2022-04-24 15:47:21
 * @LastEditTime: 2022-04-24 17:02:06
 * @LastEditors: 基于RPC实现一个打印例子
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @Reference : https://weread.qq.com/web/reader/dd63214071cc7fa3dd61bb8kd9d320f022ed9d4f495e456
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

// RPC SERVER
// 将HelloService类型的对象注册为一个RPC服务
func main() {
	// rpc.RegisterName()函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，所有注册的方法会放在HelloService服务的空间之下。
	rpc.RegisterName("HelloService", new(HelloService))
	// 然后建立一个唯一的TCP链接，
	listener, err := net.Listen("tcp", ":1234")
	if err !=nil {
		log.Fatal("ListenTCP error:" , err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	// 并且通过rpc.ServeConn()函数在该TCP链接上为对方提供RPC服务。
	rpc.ServeConn(conn)
}

// RPC CLIENT
// 下面是客户端请求HelloService服务的代码：
func main() {
	// 通过rpc.Dial拨号RPC服务
	client,err := rpc.Dial("tcp", localhost:1234)
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	// 然后通过client.Call()调用具体的RPC方法
	//然后通过client.Call()调用具体的RPC方法
	// 第二个和第三个参数分别是定义RPC方法的两个参数。
	err = client.Call("HelloSerice.hello", "hello", &reply)
	if  err !=nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
