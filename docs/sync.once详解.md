sync.Once 是 Go 官方的一并发辅助对象，它能够让函数方法只执行一次，达到类似 init 函数的效果。我们来看看它的简单用法：

```golang
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {

	onceFunc := func() {
		fmt.Println("Only once")
	}

	for i := 0; i < 10; i++ {
		once.Do(onceFunc)
	}
}
```
这里执行后我们将只看到一次 Only once 的打印信息，这就是 sync.Once 的一次性效果。

举个例子，nancy在购物时，android端和web端只有一个指向用户nancy的对象
```golang
package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	n    *Woman
)


//nancy的淘宝客户端和web端都会指向nancy这一个人
type Woman struct {
	name string
}

func GetNancy(i int) *Woman {

	getNameOnceFunc := func() {
		n = new(Woman)
		n.name = "Nancy"
		fmt.Println("newNancy")
	}

	fmt.Println("---start---")
	fmt.Println(i)
	fmt.Println("greetingNancy")
	once.Do(getNameOnceFunc)
	fmt.Println("---end----")
	return n
}

func main() {
	for i := 0; i < 5; i++ {
		_ = GetNancy(i)
	}
}
```

```bash
---start---
0
greetingNancy
newNancy
---end----
---start---
1
greetingNancy
---end----
---start---
2
greetingNancy
---end----
---start---
3
greetingNancy
---end----
---start---
4
greetingNancy
---end----


```

## 源码实现
sync.Once是由Once结构体和其Do，doSlow俩个方法实现的


```golang
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"sync/atomic"
)

// Once is an object that will perform exactly one action.
//
// A Once must not be copied after first use.
type Once struct {
	// done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/386),
	// and fewer instructions (to calculate offset) on other architectures.
	done uint32
	m    Mutex
}

// Do calls the function f if and only if Do is being called for the
// first time for this instance of Once. In other words, given
// 	var once Once
// if once.Do(f) is called multiple times, only the first call will invoke f,
// even if f has a different value in each invocation. A new instance of
// Once is required for each function to execute.
//
// Do is intended for initialization that must be run exactly once. Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
// 	config.once.Do(func() { config.init(filename) })
//
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
//
// If f panics, Do considers it to have returned; future calls of Do return
// without calling f.
//
func (o *Once) Do(f func()) {
	// Note: Here is an incorrect implementation of Do:
	//
	//	if atomic.CompareAndSwapUint32(&o.done, 0, 1) {
	//		f()
	//	}
	//
	// Do guarantees that when it returns, f has finished.
	// This implementation would not implement that guarantee:
	// given two simultaneous calls, the winner of the cas would
	// call f, and the second would return immediately, without
	// waiting for the first's call to f to complete.
	// This is why the slow path falls back to a mutex, and why
	// the atomic.StoreUint32 must be delayed until after f returns.

	if atomic.LoadUint32(&o.done) == 0 {
		// Outlined slow-path to allow inlining of the fast-path.
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

```
done是标识位，用来判断方法f是否被执行完，其初始值为0，当f执行结束时，done被设为1。