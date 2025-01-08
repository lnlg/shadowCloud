# shadowCloud

go项目使用gin封装一个简洁的项目框架

## 安装插件

```go
初始化项目
go mod init serviceCloud

引用项目需要的依赖增加到go.mod文件,去掉go.mod文件中项目不需要的依赖。
go mod tidy

安装cobra命令行工具
go get -u github.com/spf13/cobra@latest
安装gin框架
go get -u github.com/gin-gonic/gin
```

## 目录结构

```shell
$ tree -d -L 3
.
├── app
│   ├── api           //api接口层
│   ├──├── app        //应用层
│   └──└── admin      //后台管理层
│   ├── service       //service层
│   ├── middleware    //中间件层
│   ├── route         //路由
│   ├── task          //定时任务
│   ├── util          //工具层
│   └── models        //数据库操作层
├── internal          //内部包
│   ├── bootstrap     //启动引导
│   ├── config        //配置
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
