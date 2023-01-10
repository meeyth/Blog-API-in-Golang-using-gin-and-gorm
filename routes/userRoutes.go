package routes

import (
	"social-media/controllers"
	"social-media/middleware"

	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.RouterGroup) {

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:username", controllers.GetUsersByUsername)
	r.GET("/users/:username/posts", controllers.GetAUsersPost)

	account := r.Group("/account")

	account.Use(middleware.IsAuthorized())

	account.GET("/", controllers.GetAccountDetails)
	account.PUT("/", controllers.UpdateAccount)
	account.DELETE("/", controllers.DeleteAccount)
	account.GET("/posts", controllers.CurrentUsersPosts)
}
