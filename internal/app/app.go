package app

import (
	"github.com/imirjar/poliglotim-api/config"
	srv "github.com/imirjar/poliglotim-api/internal/app/http"
	"github.com/imirjar/poliglotim-api/internal/service"
	"github.com/imirjar/poliglotim-api/internal/storage"
)

func Start() error {
	config := config.New()
	storage := storage.New(config.StorageConf.MongoConn, config.StorageConf.PsqlConn)
	defer storage.Disconnect()
	service := service.New()
	srv := srv.New(config.GtwConf.Port)

	srv.Service = service
	service.Storage = storage

	return srv.Run()

}
