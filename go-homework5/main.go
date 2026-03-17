package main

import (
	"blog-system/config"
	"blog-system/routes"
	"fmt"
	"log"
)

func main() {
	if err := config.InitConfig("config.yaml"); err != nil {
		log.Fatal("初始化配置失败:", err)
	}

	if err := config.InitLogger(); err != nil {
		log.Fatal("初始化日志失败:", err)
	}

	if err := config.InitDB(); err != nil {
		log.Fatal("初始化数据库失败:", err)
	}

	defer func() {
		if config.Logger != nil {
			config.Logger.Sync()
		}
	}()

	cfg := config.GetConfig()
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	if addr == ":0" {
		addr = ":8080"
	}

	r := routes.SetupRouter()
	config.Info("服务器启动，监听地址:", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
}
