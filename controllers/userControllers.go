package controllers

import (
	"bloggify-api/database"
	"bloggify-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetUsers(c *gin.Context) {
	username := c.Query("username")

	if users, err := database.GetUsersFromDb(username); err == nil {
		if len(users) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "No users found with username " + username,
			})
			return
		}
		c.IndentedJSON(http.StatusOK, users)
		return
	}
	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"error": "Something wen wrong in the server, please try again after a while",
	})
}

func GetAccountDetails(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	var user models.User

	database.GetAccountDetailsFromDb(&user, claimsData.Issuer)

	c.IndentedJSON(http.StatusOK, user)
}

func CurrentUsersPosts(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}

	claimsData := claims.(*jwt.RegisteredClaims)

	posts, err := database.GetAUsersPostsFromDb(claimsData.Subject)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong in the server, please try again after a while",
		})
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func GetAUsersPosts(c *gin.Context) {
	username := c.Param("username")

	posts, err := database.GetAUsersPostsFromDb(username)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong in the server, please try again after a while",
		})
	}

	c.IndentedJSON(http.StatusFound, posts)
}

func UpdateAccount(c *gin.Context) {
	var toUpdateUser database.ReqBody

	claims, ok := c.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}

	if err := c.BindJSON(&toUpdateUser); err != nil {
		return
	}

	claimsData := claims.(*jwt.RegisteredClaims)

	updatedUser, err := database.UpdateAccountFromDb(&toUpdateUser, claimsData.Issuer)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

func DeleteAccount(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	var user models.User

	if err := database.DeleteAccountFromDb(&user, claimsData.Issuer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong in the server, please try again after a while",
		})
		return
	}

	c.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Account deleted successfully",
	})

}
