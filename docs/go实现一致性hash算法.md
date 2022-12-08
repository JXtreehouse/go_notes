<!--
 * @Author: your name
 * @Date: 2021-11-08 16:07:39
 * @LastEditTime: 2021-11-08 16:09:14
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/docs/go实现一致性hash算法.md
-->


一致性Hash算法是一种分布式算法，用于在一组节点中，根据某种规则映射键值到某个节点上，并能够在节点变化时尽量减少映射的变化。

在Go语言中，可以通过实现一个Hash环来实现一致性Hash算法。Hash环是一个环形的数据结构，用于映射键值到节点上。它的核心部分是一个SortedMap，用于存储映射信息。

首先，定义一个Hash环的结构体，包含一个SortedMap和一个锁：

```golang
type HashRing struct {
  // 一个有序的map，用于存储键值到节点的映射
  nodes *sortedmap.SortedMap
  // 一个互斥锁，用于保证Hash环的安全访问
  lock sync.RWMutex
}

```
接着，实现Hash环的构造函数，用于创建一个新的Hash环：

```golang
func NewHashRing() *HashRing {
  return &HashRing{
    nodes: sortedmap.New(),
  }
}


```
然后，实现Hash环的插入方法，用于向Hash环中插入一个节点

```golang
func (ring *HashRing) Insert(node string) {
  // 上锁
  ring.lock.Lock()
  defer ring.lock.Unlock()

  // 使用节点的Hash值作为键值，节点的名称作为值，插入到SortedMap中
  ring.nodes.Set(Hash(node), node)
}

```
其中，Hash函数用于计算节点的Hash值