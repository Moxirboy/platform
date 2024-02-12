package main

import (
	configs "gateway/internal/config"
	"gateway/internal/server"
	logger "gateway/pkg/log"
)


func main(){
	var (
		config = configs.Load()
	)
	log := logger.NewLogger(config.Logger.Level, config.Logger.Encoding)
	log.InitLogger()

	log.Infof(
		"AppVersion: %s, LogLevel: %s, Mode: %s",
		config.AppVersion,
		config.Logger.Level,
		config.Server.Environment,
	)
	s:=server.NewServer(config,log)
	
	log.Fatal(s.Run())
}