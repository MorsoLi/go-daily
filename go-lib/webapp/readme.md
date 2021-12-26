* 首先调用 Http.HandleFunc, 按顺序做了几件事：

1. 调用了 DefaultServeMux 的 HandleFunc
2. 调用了 DefaultServeMux 的 Handle
3. 往 DefaultServeMux 的 map [string] muxEntry 中增加对应的 handler 和路由规则

* 其次调用 http.ListenAndServe (":9090", nil), 按顺序做了几件事情：
1. 实例化 Server
2. 调用 Server 的 ListenAndServe ()
3. 调用 net.Listen ("tcp", addr) 监听端口
4. 启动一个 for 循环，在循环体中 Accept 请求
5. 对每个请求实例化一个 Conn，并且开启一个 goroutine 为这个请求进行服务 go c.serve ()
6. 读取每个请求的内容 w, err := c.readRequest ()
7. 判断 handler 是否为空，如果没有设置 handler（这个例子就没有设置 handler），handler 就设置为 DefaultServeMux
* 调用 handler 的 ServeHttp
8. 在这个例子中，下面就进入到 DefaultServeMux.ServeHttp
9. 根据 request 选择 handler，并且进入到这个 handler 的 ServeHTTP
`mux.handler(r).ServeHTTP(w, r)`
10. 选择 handler：
    - 判断是否有路由能满足这个 request（循环遍历 ServeMux 的 muxEntry）
    - 如果有路由满足，调用这个路由 handler 的 ServeHTTP
    - 如果没有路由满足，调用 NotFoundHandler 的 ServeHTTP