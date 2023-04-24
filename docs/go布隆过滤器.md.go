1、引入必要的包
```
go
import (
"net/http"
"github.com/gorilla/websocket"
)
```

2、定义websocket升级器
```
go
// 定义websocket升级器
var upgrader = websocket.Upgrader{
// 允许跨域
CheckOrigin: func(r *http.Request) bool {
return true
},
}

```


3、定义websocket处理函数

```
go
// 定义websocket处理函数
func wsHandler(w http.ResponseWriter, r *http.Request) {
// 升级http协议为websocket协议
var (
wsConn *websocket.Conn
err    error
data   []byte
conn   *impl.Connection
)
if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
return
}
if conn, err = impl.InitConnection(wsConn); err != nil {
goto ERR
}
// 启动读协程
go func() {
var (
err error
)
for {
if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
return
}
time.Sleep(1 * time.Second)
}
}()
for {
if data, err = conn.ReadMessage(); err != nil {
goto ERR
}
if err = conn.WriteMessage(data); err != nil {
goto ERR
}
}
ERR:
conn.Close()
}
```




4、注册路由
```
go
// 注册路由
http.HandleFunc("/ws", wsHandler)

```


5、启动服务
```golang
// 启动服务
http.ListenAndServe("0.0.0.0:7777", nil)
```

