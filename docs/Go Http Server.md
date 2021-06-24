### Go Http Server

#### 包里的server.go

ResponseWriter： 生成Response的接口

Handler： 处理请求和生成返回的接口

ServeMux： 路由，后面会说到ServeMux也是一种Handler

Conn : 网络连接

#### 几个接口

- Handler

  ```
  type Handler interface {
      ServeHTTP(ResponseWriter, *Request)  // 具体的逻辑函数
  }
  ```

- ResponseWriter, Flusher, Hijacker

- response          

  实现了上边三个接口

- HandlerFunc

  经常使用到的一个type

  ```
  // 这里将HandlerFunc定义为一个函数类型，因此以后当调用a = HandlerFunc(f)之后, 调用a的ServeHttp实际上就是调用f的对应方法
  type HandlerFunc func(ResponseWriter, *Request)
   
  // ServeHTTP calls f(w, r).
  func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
      f(w, r)
  }
  ```

HandlerFunc定义和ServeHTTP合起来说明HandlerFunc的所有实例是实现了ServeHttp方法的。另，实现了ServeHttp方法就是实现了接口Handler。



- URL 重定向，也称为 URL 转发，是一种当实际资源，如单个页面、表单或者整个 Web 应用被迁移到新的 URL 下的时候，保持（原有）链接可用的技术。HTTP 协议提供了一种特殊形式的响应—— HTTP 重定向（HTTP redirects）。进度条等。