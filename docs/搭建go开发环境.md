<!--
 * @Author: your name
 * @Date: 2021-04-01 20:55:30
 * @LastEditTime: 2021-04-02 16:09:31
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /go_notes/docs/搭建go开发环境.md
-->
一、Go 语言环境安装
1、下载二进制包

安装包下载地址为：https://golang.org/dl/。

如果打不开可以使用这个地址：https://golang.google.cn/dl/

![](./../assets/download.png)

各个系统对应的包名：

|    操作系统 | 包名  |
|  ----  | ----  |
| Mac  | go1.16.2.darwin-amd64.pkg  |
| Linux  | go1.16.2.linux-amd64.tar.gz  |
| Windows  | go1.16.2.windows-amd64.msi |

2、 解压到/user/local 目录

```

tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz
```

3、将 /usr/local/go/bin 目录添加至PATH环境变量(配置 GOROOT  GOPATH  PATH等环境变量)
- 注意：GOPATH路径需要包含你的项目空间路径，支持多个

> You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
```
export PATH=$PATH:/usr/local/go/bin
```

![](./../assets/gedit_profile.png)

> Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.


```
source /etc/profile
```

4、验证安装完毕
```
go version
```

5、 go env

![](./../assets/go_env.png)

二、vscode的安装
   VS Code 是微软提供的一款轻量级但功能十分强大的编辑器
1、安装地址
https://code.visualstudio.com/download
VS Code的安装比较简单，一直下一步即可。

2、相关配置
在 文件 →  首选项 → 设置 → 用户设置  做如下配置
```
{
   "go.useCodeSnippetsOnFunctionSuggestWithoutType": true,
   "go.inferGopath": false,
   "files.autoSave": "afterDelay",
   "go.gocodeAutoBuild": true,
   "go.gopath":"开发目录1:开发目录2",
   "go.useCodeSnippetsOnFunctionSuggest": true,
   "go.gotoSymbol.includeImports": true,
   "go.gocodePackageLookupMode": "go",
   "go.autocompleteUnimportedPackages":true,
    "window.zoomLevel": 0, // 表示VSCode默认的窗口缩放指数值为0
    "go.formatTool": "goimports",
    "go.testFlags": [
        "-v",
        "-count=1",
    ],
    "gitlens.advanced.messages": {
        "suppressShowKeyBindingsNotice": true
    }
}

```
然后在扩展卡中下载

|    Chinese(Simplified) Language Pack for Visual Studio Code| 
|  ----  | 
| Go| 
| GitLens|
| Git Blame  |

这些工具包可以帮助你更好的办公

![](./../assets/vscode_plugin.png)



# 常见问题
>  1、本机安装beego和bee工具时，可能出现类似 “running dsymutil failed: signal: abort trap” 错误，系bee工具与高版本Go不兼容问题

可通过重新安装Go1.8.7及以下版本解决。