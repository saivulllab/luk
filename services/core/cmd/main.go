package main

import (
	initCfg "config"
	"core/config"
	"core/db"
	"core/internal/routes"
	"fmt"
	"go.uber.org/zap"
	"logger"
)

func main() {
	cfg := config.Config{}
	initCfg.Build(&cfg)

	log := logger.NewLogger(cfg.Env)
	log.Info("Starting service",
		zap.String("name", cfg.GetServiceName()),
		zap.String("version", cfg.GetServiceVersion()))

	db.Init(&cfg.DB)
	defer db.Close()

	e := routes.InitRoutes()
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Service.Host, cfg.Service.Port)))
}
