/*
 * @Author: your name
 * @Date: 2021-12-22 10:22:03
 * @LastEditTime: 2021-12-22 21:25:32
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /base/regular-expressions.go
 https://juejin.cn/post/6861581078721740807
 */
package bufferChannel

import (
	// "bytes"
	"fmt"
	"regexp"
)

func MatchString () {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
}

func expandTest() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)
	fmt.Println(reg.NumSubexp())
}
func main () {
	　expandTest()
}
