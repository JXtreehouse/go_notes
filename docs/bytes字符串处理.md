逻辑操作中涉及到频繁拼接字符串的代码，请使用bytes.Buffer替代。使用string进行拼接会导致每次拼接都新增string对象，增加GC负担

<b>反例：</b>
```golang
var result string
for _, name := range userList {
    result += name + ","
}
return result
```

<b>正解：</b>
```golang
var buf bytes.Buffer
for _, name := range userList {
    buf.WriteString(name)
    buf.WriteString(",")
}
return buf.String()
```
