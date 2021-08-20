<!--
 * @Author: your name
 * @Date: 2021-04-16 11:35:29
 * @LastEditTime: 2021-04-16 11:50:57
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/Web框架/readme.md
-->

大部分时候，我们需要实现一个 Web 应用，第一反应是应该使用哪个框架。不同的框架设计理念和提供的功能有很大的差别。比如 Python 语言的 django和flask，前者大而全，后者小而美。Go语言/golang 也是如此，新框架层出不穷，比如Beego，Gin，Iris等。那为什么不直接使用标准库，而必须使用框架呢？在设计一个框架之前，我们需要回答框架核心为我们解决了什么问题。只有理解了这一点，才能想明白我们需要在框架中实现什么功能。

我们先看看标准库net/http如何处理一个请求。

```
func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)
    // Fatal is equivalent to Print() followed by a call to os.Exit(1).
    log.Fatal(http.ListenAndServe("", nil))
}

func handler (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}
```

当我们离开框架，使用基础库时，需要频繁手工处理的地方，就是框架的价值所在。但并不是每一个频繁处理的地方都适合在框架中完成。Python有一个很著名的Web框架，名叫bottle，整个框架由bottle.py一个文件构成，共4400行，可以说是一个微框架。那么理解这个微框架提供的特性，可以帮助我们理解框架的核心能力。

路由(Routing)：将请求映射到函数，支持动态路由。例如'/hello/:name。
模板(Templates)：使用内置模板引擎提供模板渲染机制。
工具集(Utilites)：提供对 cookies，headers 等处理机制。
插件(Plugin)：Bottle本身功能有限，但提供了插件机制。可以选择安装到全局，也可以只针对某几个路由生效。
…