package routes

import (
	"go-testify/api/controllers"
	"go-testify/infrastructure"
	"go-testify/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        infrastructure.Router
	userController controllers.UserController
}

func NewUserRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	userController controllers.UserController,

) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Group("/api")
	{
		api.POST("/user")

	}
}
