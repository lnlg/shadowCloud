package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"shadowCloud/internal/service/route"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Http struct {
	engine *gin.Engine // gin engine
	port   string      // port to listen
}

// 创建Http服务
func New() *Http {
	return &Http{
		engine: gin.New(),
		port:   ":8082",
	}
}

// 注册路由 定义一个Interface，并在外部实现Interface
func (h *Http) RegisterRoutes(r route.RouterGeneratorInterface) {
	r.AddRoute(h.engine)
}

// 启动服务
func (h *Http) Run() {
	srv := &http.Server{
		Addr:    h.port,
		Handler: h.engine,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Println("Server is running on port " + h.port)
	h.ListenSignal(srv)
}

// 监听信号处理
func (h *Http) ListenSignal(server *http.Server) {
	// 等待中断信号来优雅地关闭服务器
	quit := make(chan os.Signal, 1) // 创建一个接受信号的通道
	// kill 默认发送的是 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号
	// kill -9 发送 syscall.SIGKILL 信号
	// signal.Notify把接受到的信号通知给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞等待中断信号
	<-quit
	// 当接收到中断信号时，停止Gin服务器
	fmt.Println("Shutting down the server...")
	// 创建一个超时时间为5秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 调用Shutdown方法停止服务器，等待所有连接关闭
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Error:", err)
	}
}
