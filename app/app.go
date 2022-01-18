package app

import (
	"context"
	"go-testify/api/controllers"
	"go-testify/api/routes"
	"go-testify/infrastructure"
	"go-testify/lib"
	"go-testify/repository"
	"go-testify/services"

	"go.uber.org/fx"
)

//  Module exported for initializing application
var Module = fx.Options(
	// Group Providers...
	controllers.Module,
	repository.Module,
	routes.Module,
	services.Module,
	infrastructure.Module,
	lib.Module,
	fx.Invoke(App),
)

func App(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
	env lib.Env,
	logger lib.Logger,
	database infrastructure.Database,
	migrations infrastructure.Migrations,
) {
	_, cancel := context.WithCancel(context.Background())
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application")
			logger.Info("-----------------------------")
			logger.Info("------- gin - test -api 🚀 -------")
			logger.Info("-----------------------------")

			logger.Info("Migrating database schemas")
			migrations.Migrate()

			go func() {
				routes.Setup()
				if env.ServerPort == "" {
					handler.Run()
				} else {
					handler.Run(":" + env.ServerPort)
				}

			}()
			return nil
		},
		OnStop: func(context.Context) error {

			logger.Info("Stopping Application")
			sqlDB, _ := database.DB.DB()
			err := sqlDB.Close()

			if err != nil {
				logger.Info("Database connection not closed")
			}

			logger.Info("Closing context")
			cancel()
			return nil
		},
	})
}
