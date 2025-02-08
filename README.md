# shadowCloud

go项目使用gin封装一个简洁的项目框架

## 启动项目

```go
go run main.go start

air start
```

## 安装插件

```go
初始化项目
go mod init serviceCloud

引用项目需要的依赖增加到go.mod文件,去掉go.mod文件中项目不需要的依赖。
go mod tidy

go clean -modcache

安装cobra命令行工具
go get -u github.com/spf13/cobra@latest
安装gin框架
go get -u github.com/gin-gonic/gin

安装viper库
go get -u github.com/spf13/viper

安装fresh项目热加载库
go install github.com/pilu/fresh@latest
启动应用： 在项目根目录下运行 fresh 命令代替直接执行 go run main.go 或 go build && ./myapp：
/Users/longlian/go/code/bin/fresh
```

## 打包相关

```go
go mod tidy

# 打包
go build -o shadowCloud main.go

GOOS：目标可执行程序运行操作系统，支持 darwin freebsd， linux ，windows
GOARCH：目标可执行程序操作系统构架，包括  386， amd64 ，arm64
CGO_ENABLED=0: 只对当前一次编译生效，不影响全局设置

# 编译mac m1芯片的可执行文件
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o shadowCloud main.go

# 编译mac平台 x86的可执行文件
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o shadowCloud main.go

# 编译linux平台 x86的可执行文件
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shadowCloud main.go

# 编译windows平台 x86的可执行文件
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o shadowCloud main.go
```

## 目录结构

```shell
$ tree -d -L 3
.
├── app
│   ├── api           //api接口
│   ├──── app         //应用
│   └──── admin       //后台管理
│   ├── event         //事件
│   ├── middleware    //中间件
│   ├── route         //路由
│   ├── service       //service
│   ├── task          //定时任务
│   ├── request       //请求
│   ├── response      //响应
│   ├── util          //工具层
│   └── models        //数据库操作层
├── internal          //内部包
│   ├── bootstrap     //启动引导
│   ├── config        //配置
│   ├── crontab       //定时任务
│   ├── event         //事件
│   ├── global        //全局变量
│   ├── logger        //日志
│   ├── mysql         //数据库连接
│   ├── redis         //redis连接
│   ├── service       //http相关服务
│   ├── tool          //工具包
│   └── validator     //自定义验证器
├── cmd               //自定义命令
├── docs              //项目文档
└── runtime           //运行日志
```
