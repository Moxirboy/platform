package main

import (
	configs "teacher/internal/config"
	"teacher/internal/server"
	logger "teacher/pkg/log"
)

func main() {

	var (
		config = configs.Load()
	)

	log := logger.NewLogger(config.Logger.Level, config.Logger.Encoding)
	log.InitLogger()

	log.InfoF("AppVersion: %s, LogLevel: %s, Mode: %s",
		config.AppVersion,
		config.Logger.Level,
		config.Server.Environment,
	)

	s := server.NewServer(config, log)
	log.Fatal(s.Run())

}
