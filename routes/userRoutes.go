package routes

import (
	"bloggify-api/controllers"
	"bloggify-api/middleware"

	"github.com/gin-gonic/gin"
)

func userRoutes(r *gin.RouterGroup) {

	r.GET("/users", controllers.GetUsers)

	account := r.Group("/account")

	account.Use(middleware.IsAuthorized())

	account.GET("", controllers.GetAccountDetails)
	account.PUT("", controllers.UpdateAccount)
	account.DELETE("", controllers.DeleteAccount)
	account.GET("/posts", controllers.CurrentUsersPosts)
}
