package controllers

import (
	"log"
	"net/http"
	"social-media/database"
	"social-media/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetAllUsers(ctx *gin.Context) {
	if users, err := database.GetAllUsersFromDb(); err == nil {
		ctx.IndentedJSON(http.StatusOK, users)
		return
	}
}

func GetAccountDetails(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	var user models.User

	database.GetAccountDetailsFromDb(&user, claimsData.Issuer)

	ctx.IndentedJSON(http.StatusOK, user)
}

func CurrentUsersPosts(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}

	claimsData := claims.(*jwt.RegisteredClaims)

	posts, err := database.GetAUsersPostsFromDb(claimsData.Subject)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong in the server, please try again after a while",
		})
	}

	ctx.IndentedJSON(http.StatusOK, posts)
}

func GetUsersByUsername(ctx *gin.Context) {
	username := ctx.Param("username")

	users, err := database.GetAUserFromDb(username)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusFound, users)
}

func GetAUsersPost(ctx *gin.Context) {
	username := ctx.Param("username")

	posts, err := database.GetAUsersPostsFromDb(username)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong in the server, please try again after a while",
		})
	}

	ctx.IndentedJSON(http.StatusFound, posts)
}

func UpdateAccount(ctx *gin.Context) {
	var toUpdateUser database.ReqBody

	claims, ok := ctx.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}

	if err := ctx.BindJSON(&toUpdateUser); err != nil {
		return
	}

	claimsData := claims.(*jwt.RegisteredClaims)

	updatedUser, err := database.UpdateAccountFromDb(&toUpdateUser, claimsData.Issuer)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedUser)
}

func DeleteAccount(ctx *gin.Context) {
	claims, ok := ctx.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	var user models.User

	if err := database.DeleteAccountFromDb(&user, claimsData.Issuer); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong in the server, please try again after a while",
		})
		return
	}

	ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Account deleted successfully",
	})

}
