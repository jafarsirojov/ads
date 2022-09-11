package router

import (
	"ads/cmd/app/handler"
	"ads/pkg/config"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

var Module = fx.Invoke(NewRouter)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    *config.Config
	Logger    *zap.Logger
	Handler   handler.Handler
}

func NewRouter(params Params) {
	router := mux.NewRouter()

	router.HandleFunc("/ad", params.Handler.Add).Methods("POST")
	router.HandleFunc("/ads", params.Handler.GetList).Methods("GET")
	router.HandleFunc("/ad", params.Handler.GetByID).Methods("GET")

	server := http.Server{
		Addr:    params.Config.Port,
		Handler: router,
	}

	params.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				fmt.Println("Application started")
				params.Logger.Info("start")
				go server.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				fmt.Println("Application stopped")
				return server.Shutdown(ctx)
			},
		},
	)
}
