package main

import (
	"log"
	"net/http"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// 优雅重启：https://www.liwenzhou.com/posts/Go/graceful_shutdown/
// 使用 fvbock/endless 来替换默认的 ListenAndServe启动服务来实现，endless 是通过fork子进程处理新请求，
// 待原进程处理完当前请求后再退出的方式实现优雅重启的。
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin!")
	})

	// 默认endless服务器会监听下列信号：
	// syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
	// 接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
	// 接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
	// 接收到 SIGUSR2 信号将触发HammerTime
	// SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	if err := endless.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
	log.Println("Server exiting")
}
