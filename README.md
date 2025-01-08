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