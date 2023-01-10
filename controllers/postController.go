package controllers

import (
	"log"
	"net/http"
	"social-media/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"social-media/database"
)

func GetAllPosts(ctx *gin.Context) {
	if posts, err := database.GetAllPostsFromDb(); err == nil {
		ctx.IndentedJSON(http.StatusOK, posts)
		return
	}

	ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
		"error": "Something wen wrong in the server, please try again after a while",
	})

}

func GetPostByTitle(ctx *gin.Context) {
	title := ctx.Param("title")

	posts, err := database.GetPostByTitleFromDb(title)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, posts)
}

func PostAPost(ctx *gin.Context) {
	var newPost models.Post
	var currentUser models.User

	claims, ok := ctx.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	if err := ctx.BindJSON(&newPost); err != nil {
		return
	}
	database.GetAccountDetailsFromDb(&currentUser, claimsData.Issuer)

	newPost.Creator = currentUser.UserName

	if err := database.InsertAPostIntoDB(&newPost); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something wen wrong in the server, please try again after a while",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newPost)
}

func DeleteAPostByTitle(ctx *gin.Context) {
	var post models.Post
	title := ctx.Param("title")

	claims, ok := ctx.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	if err := database.DeleteAPostByTitleFromDB(&post, title, claimsData.Subject); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusFound, post)
}

func UpdateAPostByTitle(ctx *gin.Context) {
	title := ctx.Param("title")
	var toUpdateData models.Post

	if err := ctx.BindJSON(&toUpdateData); err != nil {
		return
	}

	updatedPost, err := database.UpdateAPostByTitleFromDb(&toUpdateData, title)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, updatedPost)
}
