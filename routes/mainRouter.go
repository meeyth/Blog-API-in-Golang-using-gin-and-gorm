package routes

import (
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/api")

	auth := api.Group("/auth")
	authRoutes(auth)

	post := api.Group("/post")
	postRoutes(post)

	user := api.Group("/user")
	userRoutes(user)

	return router
}
