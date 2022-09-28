# 复合数据类型

## 数组

固定长度，不如slice灵活，go中很少直接使用数组。

数组是值传递，不能通过修改数组类型的参数来修改这个数组的值。

````
package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)
}
````

```
[7 2 3]
[1 2 3]
```



```
q := [...]int{1,2,3}
如果数组长度位置是省略号，则数组长度由实际插入决定
```

## slice

- 变长序列。写法是[]T，T代表类型。
- 前边整理过的s[m:n]  s[m:]  s[:n]   s[:]
- 跟数组一样，切片包括一个执行第一个值的指针，所以复制是就是生成一个别名的指针，指向第一个值。
- 切片反转
- 不能通过== 判断两个切片是否相同，但是可以通过bytes.Equal函数判断两个字节型slice是否相容（[]byte），但是对于其他类型的slice必须自己展开比较。
- 判断为空;
```
len(s)==0
```
- 用内置的make函数创建一个指定元素类型、长度和容量的slice。容量部分可以忽略，容量将等于长度。
```
make([]T,len)
make([]T,len,cap)  	//类似  make([]T,cap)[:len]
```
这段看不懂
在底层，make创建了一个匿名的数组变量，然后返回一个slice；只有通过返回的slice才能用底层匿名的数组变量。在第一中语言中，slice是整个数组的view。在第二种语句中，slice只引用了底层数组的前len个元素，但是容量包含整个的数组。额外的元素是留给未来的增长用的。

### append函数
用于向slice中追加元素
```
var runes []rune
for _, r:=range "hello, 世界"{
    runes = append(runes, r)
}
fmt.Printf("%q\n", runes)     //"['h' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"
```
在循环中使用append构建一个由9个rune字符组成的slice。
- cap（）返回的是数组切片分配的空间大小，可以有预留。
```
my := make([]int, 5, 10)
cap(my)返回的是10，make生成5个初始化值的my切片。
```
-  copy(a,b)复制切片（有返回值，返回的是成功复制的元素个数），从b复制到a
- appendInt函数分配时，容量大于切片的容量会重新分配内存，*2。容量够用就不会。
- append可以追加多个元素甚至可以是一个切片。
### slice内存技巧
旋转，反转，修改元素。
- 返回不含空字符串的列表nonempty函数
```
func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
    	if s != "" {
            strings[i] = s
            i++
    	}
    }
    return strings[:i]
}
```
- 一个slice可以模拟一个stack。
```
stack = append(stack, v)
```
- 要删除slice中的某个元素但是保存原有的元素顺序，可以通过内置的copy函数将后边的子slice向前移动一位完成。
```
func remove(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}
func main() {
    s := []int{5,6,7,8,9}
    fmt.Println(remove(s, 2))   //"[5 6 8 9]"
}
```
- 改写reverse函数，用数组指针代替slice
- 给函数传递数组，传递的只是数组的拷贝，问不是他的指针
```
package main

import "fmt"

func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j ; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	s := [5]int{1,2,3,4,5}
	reverse(&s)
	fmt.Println(s)
}
```
- 编写一个rotate函数，通过一次循环完成旋转。
```
package main

import "fmt"

func rotate(s []int, x int) []int{
	a := make([]int, 5, 10)
	for i := 0; i < len(s); i++ {
		a[i] =  s[(i+x)%5]
	}
	return a
}
func main() {
	my := []int{1,2,3,4,5}
	fmt.Println(rotate(my, 5))
//	fmt.Println(my)
}
```
- 消除[]string中相邻重复字符串
```
func StringNoRep(s []string) []string {
	for i := range s {
		if i == len(s)-1 {
			return s
		}
		if s[i] == s[i+1] {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
```
## map
在go语言中，一个MAP就是一个哈希表的引用，map类型可以写成map[K]V,key和value。map中所有的key可以使同一类型，所有的value可以是同一类型，但是key和value可以是不同类型。
内置的make函数可以创建一个map：

```
ages := make(map[string]int)   // mapping from strings to ints
ages["alice"] = 31
ages["chalie"] = 34

ager := map[string]int{
    "alice": 31,
    "charlie": 34
}
```
- 内置delete函数可以删除元素。
```
delete(ages, "alice")
```
- 因为map中的元素不是变量，所以不能进行取址操作。
- 可以用range风格的for循环遍历map的key/value
- map的迭代顺序是随机的，并且不同的哈希函数实现可能导致不同的遍历顺序。如果要按顺序遍历key/value对，必须显式地对key进行排序，可以使sort包的Strings函数符字符串slice进行排序。
```
for name, age := range ages{
    fmt.Pringf("%s\t%d\n", name, age)
}
```
- 不能向一个nil的map中存数据，在向map中存数据前必须先创建map
- 使用range遍历map数据类型时，两个变量分别对应map的第一个类型和第二个类型
```
package main

import "fmt"

func main()  {
	ages := make(map[string]int)   // mapping from strings to ints
	ages["alice"] = 31
	ages["chalie"] = 34
	for i,j := range ages {
		fmt.Println(i)
		fmt.Println(j)
	}
}
```

```
alice
31
chalie
34
Process finished with exit code 0
```



## 结构体
类似于类中的结构体，需要绑定到一个对象中。利用对象进行访问、复制、取址等工作。
- 参数顺序不同也是不同的结构体。
```
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

var dilbert Employee
```
对成员取址，然后通过指针访问
```
position := &dilbert.Salary
*position = "Senior " + *position // promoted, for outsourcing to Elbonia
```
- S结构体中的成员不能是S类型的，但是可以是*S指针类型的，这样可以创建递归的数据结构，如链表和树结构。下边的代码是用数实现插入。

```
package main

import "fmt"

type tree struct {
    value       int
    left, right *tree
}//每个树的节点都是这个结构体递归出来的。
// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}
// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func main() {
	my := []int{1,2,5,8,4,2,5,8}
	Sort(my)
	fmt.Println(my)
}
```

### 结构体面值
结构体值也可以用结构体面值表示，可以指定每个成员的值
```
type Point struct{x,y int}
p := Point{1, 2}
```
通常用第二种写法,顺序和其他成员的值不用考虑
```
p := Point{X: 1}
```
结构体可以做参数和返回值，可以对结构体中的成员做缩放等处理。
因为函数的参数都是值拷贝，所以要对结构体进行修改就要传入指针。

如果结构体的所有成员都是可比较的，那么这个结构体也是可比较的。可比较的结构体类型可以用于map的key类型。
### 结构体嵌入和匿名成员
两个有相似之处的结构体可以嵌入
比如轮胎和圆和点
```
type Point struct {
    X, Y int
}

type Circle struct {
    Center Point
    Radius int
}

type Wheel struct {
    Circle Circle
    Spokes int
}
```
但是访问会很繁琐比如
var w Wheel
w.Circle.Center.X = 8
所以要有匿名成员
```
type Circle struct {
    Point
    Radius int
}

type Wheel struct {
    Circle
    Spokes int
}
```
各有一个匿名成员。特性是：类型必须是命名的类型或者是指向一个命名的类型的指针。
Point类型嵌入到了Circle结构体，Circle类型嵌入到了Wheel结构体
这样访问就变成了
var w Wheel
x.X = 8
这样它们的字面值并没有简短的表示匿名成员的语法，一下语法不存在。
```
w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
```
可以用这样的语法
```
w = Wheel{Circle{Point{8, 8}, 5}, 20}
w.X = 5
fmt.Printf("%#v\n", w)
```
#副词，代表打印时和go语言一样的语法结构
## JSON
JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象——Unicode文本编码。它可以用有效可读的方式表示第三章的基础数据类型和本章的数组、slice、结构体和map等聚合数据类型。
将GO中的一个结构体的切片转换成JSON的过程叫做编组，通过json.Marsha完成。生成的是很长的字符串，没缩进很难阅读。另一个json.Marshallndent函数是格式化的输出。
```
data, err := json.MarshalIndent(movies, "", "    ")
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}
fmt.Printf("%s\n", data)
```
