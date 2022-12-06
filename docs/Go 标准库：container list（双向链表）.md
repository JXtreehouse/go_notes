
# Go标准库: container/list(双向链表)
https://pkg.go.dev/container/list 官方标准库文档整理，讲解示例代码说明如何使用。

## 结构体
### Element

Element is an element of a linked list.
```golang
type Element struct {

	// The value stored with this element.
	Value interface{}
	// contains filtered or unexported fields
}
```

### List
List represents a doubly linked list. The zero value for List is an empty list ready to use.

type List struct {
// contains filtered or unexported fields
}

### 支持的方法

```golang
// Init initializes or clears list l.
func (l *List) Init() *List

// Len returns the number of elements of list l. The complexity is O(1).
func (l *List) Len() int

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element
// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element

// InsertAfter inserts a new element e with value v immediately after mark and returns e. If mark is not an element of l, the list is not modified. The mark must not be nil.
func (l *List) InsertAfter(v interface{}, mark *Element) *Element
func (l *List) InsertBefore(v interface{}, mark *Element) *Element

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v interface{}) *Element
// PushBackList inserts a copy of another list at the back of list l. The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List)
func (l *List) PushFront(v interface{}) *Element
func (l *List) PushFrontList(other *List)

// Remove removes e from l if e is an element of list l. It returns the element value e.Value. The element must not be nil.
func (l *List) Remove(e *Element) interface{}

// MoveAfter moves element e to its new position after mark. If e or mark is not an element of l, or e == mark, the list is not modified. The element and mark must not be nil.
func (l *List) MoveAfter(e, mark *Element)
func (l *List) MoveBefore(e, mark *Element)
// MoveToBack moves element e to the back of list l. If e is not an element of l, the list is not modified. The element must not be nil.
func (l *List) MoveToBack(e *Element)
func (l *List) MoveToFront(e *Element)

```


# 示例代码

```golang
package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

```

输出:

```golang
1
2
3
4

```
说明:

链表中元素的类型可以为任何类型，示例中的代码所有元素类型都为整数
查找某个元素是否在链表中可以使用如下的查找函数：

```golang
func search(l *list.List, v interface{}) *list.Element {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == v {
			return e
		}
	}
	return nil
}


```
- Init 函数可将链表清空






