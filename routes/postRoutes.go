package routes

import (
	"social-media/controllers"
	"social-media/middleware"

	"github.com/gin-gonic/gin"
)

func postRoutes(r *gin.RouterGroup) {
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/post/:title", controllers.GetPostByTitle)

	account := r.Group("/account")
	account.Use(middleware.IsAuthorized())

	account.POST("/post", controllers.PostAPost)
	account.PUT("/post/:title", controllers.UpdateAPostByTitle)
	account.DELETE("/post/:title", controllers.DeleteAPostByTitle)
}
