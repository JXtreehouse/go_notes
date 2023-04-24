<!--
 * @Author: your name
 * @Date: 2021-11-08 16:07:39
 * @LastEditTime: 2021-11-08 16:09:14
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/docs/go实现一致性hash算法.md
-->

一致性哈希的概念在 Karger 1997年发布的论文 《一致的哈希和随机树：缓解万维网上的热点的分布式缓存协议》 中引入，
简单来说，一致性哈希将整个哈希值空间组织成一个虚拟的圆环（0-2^32-1）。
一致性哈希之后在许多其他分布式系统（如Cassandra，Riak等）中使用，并不断优化和改良。

# 什么是Hash
Hash是将一个数据（通常是任意大小的对象）映射到另一固定大小的数据（通常是整数，称为哈希码或简称hash）的过程。把将某数据映射到哈希码的这个函数hash()，称为哈希函数。

例如，哈希函数可用于将随机大小的字符串映射到0到N之间的某个固定整数数字。

```bash
hello ---> 60
hello world ---> 40
```
可能会有字符串映射到相同的整数的场景，被称为<b>碰撞</b>。

处理碰撞的常见解决方案是:
- <b>链式</b>（常用）
  - 将哈希表的每个单元作为链表的头结点，所有哈希地址为i的元素构成一个同义词链表
  - 即发生冲突时就把该关键字链在以该单元为头结点的链表的尾部。

- <b>开放式寻址</b>。
  - 一旦发生了冲突，就去寻找下一个空的散列地址，只要散列表足够大，空的散列地址总能找到，并将记录存入。


# 场景
目前有一个有状态的缓存服务，服务需要缓存n位员工的email信息到姓名的映射：john@example.com 张三，mark@example.com 李四 … ，在访问时，我们需要对其做增删查

<b>存储方式对比</b>

|  数据算法 |   |   |   |   |
|---|---|---|---|---|
|  数组 |  每个操作的最坏情况时间复杂度将为O（n）。通过存储排序的数据并使用二分查找，可以将搜索优化为O（logn） |   |   |   |
|  链表|  如果我们将使用链接列表存储员工记录，则最坏情况下的插入时间为O（1），搜索和删除时间为O（n） |   |   |   |
| 二叉搜索树（Binary Search Tree） | 每个操作的最坏情况时间将为O（log n）  |   |   |   |
|  哈希函数+数组 |   使用哈希函数可用于将对象key（即email）哈希为固定大小的整数。然后我们可以使用数组下标i作为哈希结果，存储员工详细信息。但考虑到哈希函数的散列范围大，使用数组下标需要创建一个大数组，会非常浪费内存|   |   |   |
|  哈希函数+哈希表 |  使用哈希表可以解决这个问题，哈希表使用固定大小的N数组来映射所有键的哈希码。假设我们是对key哈希再取模获取数组索引：index = hash(key) mod N由于可能有多个key映射到同一索引，因此每个索引都将附加一个列表或存储桶bucket，以存储映射到同一索引的所有对象。 |   |   |   |

所以方案5使用哈希表是一个很好的方案，由于哈希函数为O(1)复杂度，如果哈希表的大小得当，每个存储桶的数量较少，那么访问的速度很快。
# 分布式哈希
考虑在刚才的场景下，如果员工数量大大增长，存储在一台计算机上的哈希表中变得困难。
在这种情况下，我们希望能够将哈希表分发到多个服务器，避免单台服务器的限制。所以我们需要使key尽量均匀的分布在多个服务器。
最常见的方法是hash取模：
```nashorn js
server = hash（key）mod 3  //假设后端有3个服务器
```
三个服务器分别是S1，S2和S3，示例结果如下：
```bash
john@example.com      89 2（S3）
mark@example.com      30 0（S1）
adam@examle.com       47 2（S3）
smith@example.com     52 1（S2）
alex@example.com     75 0（S0）
bob@example.com      22 1（S2）
```

**分布式哈希存在的问题**:

**重新哈希问题**

假设其中一台服务器（S3）崩溃了，它不再能够接受请求。现在我们只剩下两台服务器了。几乎所有密钥的服务器位置都已更改，不仅是S3中的密钥。

对后端真实服务器来说，将大大增加服务器的负载（因为会丢失缓存后需要重新查询建立映射）。

对客户端来说，并且由于映射方式的改变大量mail都被映射到新的服务器上，导致映射在这一刻都不可用。这个问题被称为重新哈希问题。

```bash
john@example.com      89 1（S2）
mark@example.com      30 0（S1）-
adam@examle.com       47 1（S2）
smith@example.com     52 0（S1）
alex@example.com     75 1（S2）
bob@example.com      22 0（S1）
```
所以 我们需要一个一致性哈希算法

一致性哈希是一种分布式哈希方案，使服务器的伸缩不会给整个系统带来大问题

# 经典一致性哈希
1997年，Karger等人发布了论文 《Consistent Hashing and Random Trees: Distributed Caching Protocols for Relieving Hot Spots on the WWW》 ，并提出了一致性哈希

在论文中还对一致性哈希的算法好坏定义给出了4个评判指标:

- 平衡性（Balance）
不同key的哈希结果分布均衡，尽可能的均衡地分布到各节点上。平衡性跟哈希函数关系密切，目前许多哈希算法都有较好的平衡性。
- 单调性（Monotonicity）
当有新的节点上线后，系统中原有的key要么还是映射到原来的节点上，映射到新加入的节点上，不会出现从一个老节点重新映射到另一个老节点。即表示：当可用存储桶的集合发生更改时，只有在必要时才移动项目以保持均匀的分布。
- 分散性（Spread）
由于客户端可能看不到后端的所有服务，这种情况下对于固定的key，在两个客户端上可能被分散到不同的后端服务，从而降低后端存储的效率，所以算法应该尽量降低分散性。
- 服务器负载均衡（Load）
负载主要是从服务器的角度来看，指各服务器的负载应该尽量均衡

# 算法思路

Karger等人在后续的论文 《Web caching with consistent hashing》 中提出了一致性哈希的实现，也就是大家常称的环割法。
<b>一致性哈希思路</b>

我们把节点通过hash后，映射到一个范围是[0，2^32]的环上
# 用go实现一致性hash算法

一致性Hash算法是一种分布式算法，用于在一组节点中，根据某种规则映射键值到某个节点上，并能够在节点变化时尽量减少映射的变化。

在Go语言中，可以通过实现一个Hash环来实现一致性Hash算法。

Hash环是一个环形的数据结构，用于映射键值到节点上。

它的核心部分是一个SortedMap，用于存储映射信息。

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

# Reference
- [Consistent Hashing and Random Trees:
  Distributed Caching Protocols for Relieving Hot Spots on the World Wide Web](https://www.cs.princeton.edu/courses/archive/fall09/cos518/papers/chash.pdf)
- [Web caching with consistent hashing](http://www.cs.columbia.edu/~asherman/papers/cachePaper.pdf)
- [一致性哈希 （Consistent Hashing）的前世今生](https://candicexiao.com/consistenthashing/)
- https://github.com/serialx/hashring
- [ketama算法Golang实现](https://blog.kimq.cn/2017/07/06/ketama/)