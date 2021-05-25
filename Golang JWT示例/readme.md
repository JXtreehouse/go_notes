<!--
 * @Author: your name
 * @Date: 2021-04-08 14:55:07
 * @LastEditTime: 2021-04-08 19:09:47
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/Golang JWT示例/readme.md
-->
Run

![](../asset/../assets/go_jwt.png)
```
$ go run main.go 
> secret_key
> eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMTIsImV4cCI6MTU3NTYzMzE2MCwiaXNzIjoibG9jYWxob3N0In0.tl0obbcF9me3LALR-3uCYuGSFcxK7kXRoFXGz4dj1fk <nil>%
```
![](../assets/go_jwt_parse.png)

```

$ go run parse.go
> secret_key
> 10012 1575633160

```