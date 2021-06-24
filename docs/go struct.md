<!--
 * @Author: your name
 * @Date: 2021-06-24 20:42:48
 * @LastEditTime: 2021-06-24 20:45:01
 * @LastEditors: Please set LastEditors
 * @Description:  https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.4.md
 * @FilePath: /docs/go struct.md
-->
struct 作为其它类型的属性或字段的容器

例如，我们可以创建一个自定义类型person代表一个人的实体。这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。如下代码所示:

```
type person struct {
	name string
	age int
}
```
- 一个string类型的字段name，用来保存用户名称这个属性
- 一个int类型的字段age,用来保存用户年龄这个属性