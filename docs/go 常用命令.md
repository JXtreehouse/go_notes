
```commandline
# 最常用的 go command 之一，编译 go 文件
# 跨平台编译：env GOOS=linux GOARCH=amd64 go build
go install
# 也是编译，与 build最大的区别是编译后会将输出文件打包成库放在 pkg 下
go get
# 用于获取 go 的第三方包，通常会默认从 git repo 上 pull 最新的版本
# 常用命令如：go get -u github.com/go-sql-driver/mysql（从 github 上获取 MySQL的 driver 并安装至本地）
go fmt
# 类似于 C 中的 lint，统一代码风格和排版
# 常用命令如：go fmt
go test 
# 运行当前包目录下的 tests
# 常用命令如go test 或 go test -v
/* main.go
/* main_test.go
# Go 的 test 一般以 XXX_test.go 为文件名
# XXX 部分一般为 XXX_test.go 所要测试的代码文件名
# 注：Go没有特别要求 XXX 部分必须 是要测试的文件名
```

# 交叉编译

```commandline
# macOS下编译Linux及Windows 64位可执行程序(-o用于指定路径及文件名，可省略)
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o xxx main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o xxx main.go
 
# Linux下编译Mac及Windows 64位可执行程序
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
 
#Windows下编译Mac及Linux 64位可执行程序
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go
 
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

# 查询出所有的库函数、结构体， 如net/http

```commandline
go doc net/http | grep "^func"
go doc net/http | grep "^type"|grep struct
```

# 使用Docker容器编译

```commandline


docker pull golang:1.15.12-alpine3.13
docker run --rm -it \
-v /path/xxx:/app \
-w /app/src \          # 假设main.go及go.mod放在/path/xxx目录下
-e CGO_ENABLED=0 \     # 如在 alpine 中执行可不指定
-e GOPROXY=https://goproxy.cn \
golang:1.15.12-alpine3.13 \
go build -o ../path/xxx main.go
```

出现/lib/ld-musl-x86_64.so.1: bad ELF interpreter: No such file or directory报错可在编译时使用-e CGO_ENABLED=0，或依然在 alpine 内执行

# Golang操作 Docker API
API文档地址：https://docs.docker.com/engine/api/v1.41/

```commandline
 配置文件/usr/lib/systemd/system/docker.service中的ExecStart最后面添加-H tcp://0.0.0.0:2345
ExecStart=xxxxx -H tcp://0.0.0.0:2345
sudo systemctl daemon-reload
sudo systemctl restart docker
# 验证命令
docker -H tcp://ip.address.xxx:2345 ps
 
go get github.com/docker/docker/client
 
# 示例代码 NewClient第2个参数为API 版本，可通过 docker version 进行查看
cli, err := client.NewClient("tcp://the.ip.addr:2345", "1.41", nil, nil)
if err != nil {
   log.Fatal(err)
}
images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
if err != nil {
   log.Fatal(err)
}
for _, image := range images{
   fmt.Println(image.RepoTags)
}
```
# 单元测试覆盖率
```commandline
go test -race -cover  -coverprofile=./coverage.out -timeout=10m -short -v ./...
go tool cover -func ./coverage.out
```

# 利用工具检测变量遮蔽问题(不能定位所有问题)
```commandline
$go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
$go vet -vettool=$(which shadow) -strict xxx.go
```
