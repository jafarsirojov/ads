package main

import (
	"ads/cmd/app/handler"
	"ads/cmd/app/router"
	"ads/internal"
	"ads/pkg/config"
	"ads/pkg/db"
	"ads/pkg/logger"
	"ads/pkg/repository"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		config.Module,
		db.Module,
		repository.Module,
		logger.Module,
		handler.Module,
		internal.Module,
		router.Module,
	).Run()

}
