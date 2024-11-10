package main

import (
	"github.com/KauNdb/bookly/internal/config"
	"github.com/KauNdb/bookly/internal/logger"
	"github.com/KauNdb/bookly/internal/server"
	"github.com/KauNdb/bookly/internal/storage"
)

func main() {
	cfg := config.ReadConfig()
	log := logger.Get(cfg.Debug)
	log.Debug().Any("cfg", cfg).Send()

	stor := storage.New()

	serv := server.New(*cfg, stor)
	err := serv.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("server fatal error")
	}
	log.Info().Msg("server stoped")
}
