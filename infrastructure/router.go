package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router -> Gin Router
type Router struct {
	*gin.Engine
}

//NewRouter : all the routes are defined here
func NewRouter() Router {
	fmt.Println("Getting user routes...")
	// intialize a router.
	httpRouter := gin.Default()
	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "gin_api 📺  Up and Running"})
	})

	return Router{
		httpRouter,
	}
}
