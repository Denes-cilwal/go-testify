package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []IRoute

// Route interface
type IRoute interface {
	Setup()
}

// NewRoutes sets up routes
// NewRoutes sets up routes
func NewRoutes(userRoutes UserRoutes) Routes {
	return Routes{
		userRoutes,
	}
}

// Setup all the route

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
