package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true

	router.Use(cors.New(config))

	api := router.Group("/api")

	auth := api.Group("/auth")
	authRoutes(auth)

	post := api.Group("/post")
	postRoutes(post)

	user := api.Group("/user")
	userRoutes(user)

	return router
}
