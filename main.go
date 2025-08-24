package main

import (
	"log"
	"webapi-example/cmd"
	"webapi-example/pkg/config"
	"webapi-example/pkg/database"
	"webapi-example/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化日志
	cfg := config.Cfg
	logger.InitLogger(cfg.Log.Path)

	// 初始化数据库
	_, err = database.InitDatabase(&cfg.MySQL, cfg.Debug)
	if err != nil {
		logger.Log.Fatal("Failed to init database:", zap.Any("err", err))
	}
	logger.Log.Info("Load config success, Database init success", zap.Any("config", cfg))
	cmd.Execute()
}
