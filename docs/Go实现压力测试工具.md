
# 压测是什么

压测，　即压力测试，是确定系统稳定性的一种测试方法。通常在系统正常运行范围之外进行，以考察其功能极限和隐患。

主要检测服务器承载能力，包括用户承受能力（多少用户同时使用不影响体验和功能），流量承受等

# 为什么要压测

压测的目的是通过压测(模拟真实用户的行为)。测算出机器的性能(单台机器的qps).从而推算出系统在承受指定用户(100w)时，需要多少机器

压测是在上线前为了应对未来可能达到的用户量（流量）的一次预估(提前演练),
压测以后通过优化程序的性能或者准备充足机器，来保证用户的体验。

| 压测名词 | 解释  |
|:--: |:-----:|
| 并发(Concurrency) | 指一个处理器同时处理多个任务的能力(逻辑上处理的能力)    |
|  并行(Parallel) |  多个处理器或者是多核的处理器同时处理多个不同的任务(物理上同时执行)    |
|　QPS（每秒钟查询数量　Query Per Second）  |  服务器每秒处理请求数量(req/sec 请求数/秒　一段时间内总请求数/请求时间)    |
| 事务(Transactions) | 用户一次或者多次请求的集合     |
| TPS（每秒钟处理事务数量 Transaction Per Second） | 服务器每秒钟处理事务数量(一个事务可能包括多个请求)     |
| 请求成功数(Request Success Number)  | 在一次压测中，请求成功的数量     |
| 请求失败数(Request Failures Number) |  在一次压测中，请求失败的数量    |
| 错误率(Error Rate) | 在压测中，请求成功的数量与请求失败数量的比率     |
| 最大响应时间(Max Response Time) | 　在一次压测中，　从发出请求或者指令系统做出的反映(响应)的最大时间    |
| 最少响应时间(Minimum Response Time) | 　在一次压测中，　从发出请求或者指令系统做出的反映(响应)的最少时间    |
| 最少响应时间(Average Response Time) | 　在一次压测中，　从发出请求或者指令系统做出的反映(响应)的平均时间    |


# 实现逻辑

在Go语言中，可以使用sync.WaitGroup和go关键字来实现压力测试工具。

首先，定义一个函数，用于模拟要测试的业务逻辑：

```golang
func businessLogic() {
  // 模拟业务逻辑的处理时间
  time.Sleep(time.Millisecond * 100)
}

```
然后，定义一个函数，用于执行压力测试：

```golang
func pressureTest(concurrency int) {
  // 创建一个WaitGroup
  var wg sync.WaitGroup

  // 设置WaitGroup的计数器
  wg.Add(concurrency)

  // 开启多个goroutine执行压力测试
  for i := 0; i < concurrency; i++ {
    go func() {
      // 执行业务逻辑
      businessLogic()
      // 将WaitGroup的计数器减1
      wg.Done()
    }()
  }

  // 等待所有goroutine完成
  wg.Wait()
}

```

最后，可以通过调用pressureTest函数来执行压力测试，并指定并发数：

```golang
pressureTest(1000)

```

通过以上方法，就可以实现一个简单的压力测试工具，用于模拟并发的业务逻辑，并测试程序的性能。