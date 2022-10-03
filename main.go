package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/xppcnn/gin-demo/config"
	"github.com/xppcnn/gin-demo/initialize"
	"github.com/xppcnn/gin-demo/middleware"
	"github.com/xppcnn/gin-demo/pkg/setting"
	"go.uber.org/zap"
)

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
	Version    string `mapstructure:"version"`
}

var logConf = new(LogConfig)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	conf := config.GetConfig()
	if err := middleware.InitLogger(conf.LogConfig, conf.RunMode); err != nil {
		fmt.Println(err)
	}
	router := initialize.InitRouter()
	s := endless.NewServer(endPoint, router)
	s.BeforeBegin = func(add string) {
		zap.L().Info("init", zap.Int("actual pid is %d", syscall.Getpid()))
	}

	// logging.Info("conf  is %+v", logConf)
	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal(fmt.Sprintf("server listen err:%s", err))
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		zap.L().Fatal(fmt.Sprintf("Server Shutdown:%s", err))
	}
	zap.L().Info("Server exiting")
}
