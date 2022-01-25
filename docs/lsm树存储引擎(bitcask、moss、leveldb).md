# 数据存储与检索背景



| 维度   | OLTP(online transaction processing)　面向行 | OLAP(online analysis processing)　面向列 |
|------|-----------------------------------------|--------------------------------------|
| 系统功能 | 日常交易处理/在线事务处理                           | 统计、分析、报表/在线分析处理                      |
| 设计目标 | 面向实时交易类应用                               | 面向统计分析类应用                            |
| 数据处理 | 当前的。最新的                                 | 历史的，集聚的                              |
| 实时性  | 实时性读写要求高                                | 实时性读写要求低                             |
| 事务   | 强事务                                     | 弱事务                                  |
| 分析要求 | 低、简单                                    | 高、复杂                                 |

# 基于结构的存储引擎
![](../assets/store&search.png)

# 面向页(b+ tree)存储引擎

# 基于日志结构(lsm tree)存储引擎

## 理论

## 实践

## 总结

# 参考
- http://kernelmaker.github.io/
- www.cs.umb.edu/~poneil/lsmtree.pdf
- https://kernelmaker.github.io/lsm-tree
