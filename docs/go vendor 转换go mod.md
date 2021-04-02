<!--
 * @Author: your name
 * @Date: 2021-04-02 11:37:37
 * @LastEditTime: 2021-04-02 11:38:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/go vendor 转换go mod.md
-->
1. go mod init [git地址] (生成go.mod)
2. Go Modules 勾选 Enable Go Modules integration
3. 删除vendor包
4. 如果有灰色，把require删去 再次 go mod tidy