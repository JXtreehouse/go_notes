#! https://zhuanlan.zhihu.com/p/501300715
<!--
 * @Author: your name
 * @Date: 2022-04-19 16:53:47
 * @LastEditTime: 2022-04-19 17:07:23
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /go_notes/docs/如何在go语言中判断值的类型.md
-->

# 在 Golang 中查找变量类型的不同方法

变量是可以在运行时更改的信息的占位符。变量允许检索和操作存储的信息。

Golang中查找变量类型的方法有3种，如下：

- reflect.TypeOf 
-  reflect.ValueOf.Kind()
- %T 和 Printf
- 类型断言 `v, ok := i.(T)`

```

// Golang program to show the different ways
// to find the Type of a Variable
package main

// import the fmt and reflect package
import (
	"fmt"
	"reflect"
)

//main function

func main() {

	// string type
	var1 := "hello world"

	// integer
	var2 := 10

	// float
	var3 := 1.55

	// boolean
	var4 := true

	// shorthand string array declaration
	var5 := []string{"foo", "bar", "baz"}

	// map is reference datatype
	var6 := map[int]string{100: "Ana", 101: "Lisa", 102: "Rob"}

	// complex64 and complex128
	// is basic datatype
	var7 := complex(9, 15)

	// using %T format specifier to
	// determine the datatype of the variables

	fmt.Println("Using Percent T with Printf")
	fmt.Println()

	fmt.Printf("var1 = %T\n", var1)
	fmt.Printf("var2 = %T\n", var2)
	fmt.Printf("var3 = %T\n", var3)
	fmt.Printf("var4 = %T\n", var4)
	fmt.Printf("var5 = %T\n", var5)
	fmt.Printf("var6 = %T\n", var6)
	fmt.Printf("var7 = %T\n", var7)

	// using TypeOf() method of reflect package
	// to determine the datatype of the variables
	fmt.Println()
	fmt.Println("Using reflect.TypeOf Function")
	fmt.Println()

	fmt.Println("var1 = ", reflect.TypeOf(var1))
	fmt.Println("var2 = ", reflect.TypeOf(var2))
	fmt.Println("var3 = ", reflect.TypeOf(var3))
	fmt.Println("var4 = ", reflect.TypeOf(var4))
	fmt.Println("var5 = ", reflect.TypeOf(var5))
	fmt.Println("var6 = ", reflect.TypeOf(var6))
	fmt.Println("var7 = ", reflect.TypeOf(var7))

	// using ValueOf() method of reflect package
	// to determine the value of the variable
	// Kind() method returns the datatype of the
	// value fetched by the ValueOf() method
	fmt.Println()
	fmt.Println("Using reflect.ValueOf.Kind() Function")
	fmt.Println()

	fmt.Println("var1 = ", reflect.ValueOf(var1).Kind())
	fmt.Println("var2 = ", reflect.ValueOf(var2).Kind())
	fmt.Println("var3 = ", reflect.ValueOf(var3).Kind())
	fmt.Println("var4 = ", reflect.ValueOf(var4).Kind())
	fmt.Println("var5 = ", reflect.ValueOf(var5).Kind())
	fmt.Println("var6 = ", reflect.ValueOf(var6).Kind())
	fmt.Println("var7 = ", reflect.ValueOf(var7).Kind())

}

```

**输出**


```
Using Percent T with Printf

var1 = string
var2 = int
var3 = float64
var4 = bool
var5 = []string
var6 = map[int]string
var7 = complex128

Using reflect.TypeOf Function

var1 =  string
var2 =  int
var3 =  float64
var4 =  bool
var5 =  []string
var6 =  map[int]string
var7 =  complex128

Using reflect.ValueOf.Kind() Function

var1 =  string
var2 =  int
var3 =  float64
var4 =  bool
var5 =  slice
var6 =  map
var7 =  complex128

```

##  %T 和 Printf
 Go 中检查值类型的一种快速方法是将 %T 动词与 fmt.Printf 结合使用。如果您想将类型打印到控制台以进行调试，这很有效。

```
package main

import (
	"fmt"
)

func main() {
	a, b, c := "yes", 20.6, false
	d := []int{3, 4, 5, 6}
	e := map[string]string{
		"name":  "jake",
		"email": "jake@example.com",
	}
	fmt.Printf("type of a is %T\n", a)
	fmt.Printf("type of b is %T\n", b)
	fmt.Printf("type of c is %T\n", c)
	fmt.Printf("type of d is %T\n", d)
	fmt.Printf("type of e is %T\n", e)
}
```
```
type of a is string
type of b is float64
type of c is bool
type of d is []int
type of e is map[string]string
```

## reflect.TypeOf 
确定值类型的另一种方法是使用反射包中的 TypeOf 方法：

```
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b, c := "yes", 20.6, false
	d := []int{3, 4, 5, 6}
	e := map[string]string{
		"name":  "jake",
		"email": "jake@example.com",
	}
	fmt.Printf("type of a is %v\n", reflect.TypeOf(a))
	fmt.Printf("type of b is %v\n", reflect.TypeOf(b))
	fmt.Printf("type of c is %v\n", reflect.TypeOf(c))
	fmt.Printf("type of d is %v\n", reflect.TypeOf(d))
	fmt.Printf("type of e is %v\n", reflect.TypeOf(e))
}
```


```
type of a is string
type of b is float64
type of c is bool
type of d is []int
type of e is map[string]string
```

## 类型断言

如果你有一个接口值并且你想检查它的底层类型的值，你需要使用如下形式的类型断言

```
v, ok := i.(T)
```

```
package main

import (
	"fmt"
)

func main() {
	var s interface{} = "string"
	if _, ok := s.(string); ok {
		fmt.Println("s is a string")
	} else {
		fmt.Println("s is not a string")
	}
}
```

```
s is a string
```

如果 s 包含一个字符串值，则 ok 变量将为真。否则，它将是错误的。这是如何在运行时检查接口值的类型。