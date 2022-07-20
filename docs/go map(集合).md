#! https://zhuanlan.zhihu.com/p/543496771
<!--
 * @Author: AlexZ33
 * @Date: 2021-04-08 14:26:04
 * @LastEditTime: 2022-07-19 14:42:42
 * @LastEditors: AlexZ33 775136985@qq.com
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/go map(集合).md
-->



# Map
散列表(hash table)是一个拥有键值对元素的无序集合。<b>在这个集合中，键的值是唯一的</b>, 键对应的值可以通过键来获取、更新或移除。无论这个散列表有多大，这些操作基本上是通过常量时间的键比较就可以完成。

在GO语言中，map是散列表(hash table)的引用，map的类型是`map[K]V`, 其中K和V是字典的键和值对应的数据类型。<b>map中所有的键都拥有相同的数据类型，同时所有的值也都拥有相同的数据类型, 但键的类型和值的类型不一定相同</b>
- 键的类型K，必须是可以通过操作符`==`或者 `!= `来进行比较的数据类型，所以map可以检测某一个键是否已经存在
- 大多数内置类型都可以用作键，如 int、float64、rune、string、可比较的数组和结构、指针等
- slice和不可比较的arrays和struct等数据类型或不可比较的自定义数据类型不用作map key。
- 注意：虽然浮点类型是可以比较的，但是比较浮点的相等性不是一个好主意，尤其是在NaN可以是浮点型的时候。
- 值类型V没有任何限制,可以是任何类型，如 int、float64、rune、字符串、指针、引用类型、map类型等。  
- map也被称为hash map, hash table, unordered map, dictionary, or [associative array](https://en.wikipedia.org/wiki/Associative_array)
- 在maps中，只有在map初始化的时候才能加值，如果你试图在未初始化的map中加值，那么编译器会报错。
- 在maps中, map的零值是 nil 并且 nil map不包含任何key,如果您尝试在 nil map中添加key-value对，那么编译器将抛出运行时错误。
优势
- 由于它是引用类型，通过它的资源占用成本很低, 如对于 64 位机器，它需要 8 个字节，而对于 32 位机器，它需要 4 个字节

# 创建 & 初始化
### Creating Map using  map[Key_Type]Value_Type{}


```golang
// you can simply create a map using the given syntax:
// An Empty map
map[Key_Type]Value_Type{}

// Map with key-value pair
map[Key_Type]Value_Type{key1: value1, ..., keyN: valueN}
```
- Example:

```golang
var mymap map[int]string
```



```golang
//Initializing map using map literals

// Go program to illustrate how to
// create and initialize maps
package main

import "fmt"

func main() {

	// Creating and initializing empty map
	// Using var keyword
	var map_1 map[int]int

	// Checking if the map is nil or not
	if map_1 == nil {
	
		fmt.Println("True")
	} else {
	
		fmt.Println("False")
	}

	// Creating and initializing a map
	// Using shorthand declaration and
	// using map literals
	map_2 := map[int]string{
	
			90: "Dog",
			91: "Cat",
			92: "Cow",
			93: "Bird",
			94: "Rabbit",
	}
	
	fmt.Println("Map-2: ", map_2)
}

```
```golang
True
Map-2:  map[90:Dog 91:Cat 92:Cow 93:Bird 94:Rabbit]
```
### create a map using make() function
还可以使用 make() 函数创建map。这个函数是一个内置函数，在这个方法中，你只需要传递map的类型，它就会返回一个初始化的map。

```dtd
make(map[Key_Type]Value_Type, initial_Capacity)
make(map[Key_Type]Value_Type)
```

内置函数可以用来创建一个map:
```golang
ages := make(map[string]int) //创建一个从string到int的map
```
也可以使用map的字面量来新建一个带初始化键值对元素的字典:
```golang
ages:= map[string]int {
    "alice" : 31,
    "charlie": 34
}
```
这个等价于:
```golang
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```
因此，　新的空map的另外一种表达形式是: `map[string]int{}`

<b>map的元素访问</b>:
```golang
ages["alice"] = 32

```

# Go map size

map的大小由 len 函数确定。它返回key-value中的对数。

```golang
package main

import "fmt"

func main() {

    countries := map[string]string{
        "sk": "Slovakia",
        "ru": "Russia",
        "de": "Germany",
        "no": "Norway",
    }

    fmt.Printf("There are %d pairs in the map\n", len(countries))
}
```

```
$ go run length.go
There are 4 pairs in the map
```
# Go map loop
使用 for 和 range 关键字，我们可以遍历map元素。

```golang
package main

import "fmt"

func main() {

    countries := map[string]string{
        "sk": "Slovakia",
        "ru": "Russia",
        "de": "Germany",
        "no": "Norway",
    }

    for country := range countries {
        fmt.Println(country, "=>", countries[country])
    }

    for key, value := range countries {
        fmt.Printf("countries[%s] = %s\n", key, value)
    }
}
```
在代码示例中，我们以两种方式循环国家map。
```
for country := range countries {
    fmt.Println(country, "=>", countries[country])
}
```
在第一种情况下，我们按对对象循环。
```golang
for key, value := range countries {
    fmt.Printf("countries[%s] = %s\n", key, value)
}
```
在第二种情况下，我们按键和值循环。
```golang
$ go run loop.go
de => Germany
no => Norway
sk => Slovakia
ru => Russia
countries[ru] = Russia
countries[de] = Germany
countries[no] = Norway
countries[sk] = Slovakia
```


# 面试题
>1. map不初始化使用会怎么样
>2. map不初始化长度和初始化长度的区别
>3. map承载多大，大了怎么办
>4. map的iterator是否安全？能不能一边delete一边遍历？
>5. map触发扩容的时机，满足什么条件时扩容？
> 
>6. map扩容策略是什么
>7. 线程安全的map怎么实现
>
> 
# Reference
[Associative array](https://en.wikipedia.org/wiki/Associative_array)
[golang-maps](https://www.geeksforgeeks.org/golang-maps/)

[golang-map](https://zetcode.com/golang/map/)