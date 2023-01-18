package routes

import (
	"bloggify-api/controllers"

	"github.com/gin-gonic/gin"
)

func authRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
}
