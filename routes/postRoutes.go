package routes

import (
	"bloggify-api/controllers"
	"bloggify-api/middleware"

	"github.com/gin-gonic/gin"
)

func postRoutes(r *gin.RouterGroup) {
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:username", controllers.GetPostsOfAUser)

	account := r.Group("/account")
	account.Use(middleware.IsAuthorized())

	account.POST("/post", controllers.PostAPost)
	account.PUT("/post/:title", controllers.UpdateAPostByTitle)
	account.DELETE("/post/:title", controllers.DeleteAPostByTitle)
}
