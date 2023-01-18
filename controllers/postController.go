package controllers

import (
	"bloggify-api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"bloggify-api/database"
)

func GetPosts(c *gin.Context) {
	title := c.Query("title")

	if posts, err := database.GetPostsFromDb(title); err == nil {
		if len(posts) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "No posts found with title " + title,
			})
			return
		}
		c.IndentedJSON(http.StatusOK, posts)
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"error": "Something wen wrong in the server, please try again after a while",
	})
}

func GetPostsOfAUser(c *gin.Context) {
	username := c.Param("username")

	if posts, err := database.GetPostsOfAUserFromDb(username); err == nil {
		if len(posts) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "This user has no posts",
			})
			return
		}
		c.IndentedJSON(http.StatusOK, posts)
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"error": "Something wen wrong in the server, please try again after a while",
	})
}

func PostAPost(c *gin.Context) {
	var newPost models.Blog
	var currentUser models.User

	claims, ok := c.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	if err := c.BindJSON(&newPost); err != nil {
		return
	}
	database.GetAccountDetailsFromDb(&currentUser, claimsData.Issuer)

	newPost.Creator = currentUser.UserName

	if err := database.InsertAPostIntoDB(&newPost); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something wen wrong in the server, please try again after a while",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, newPost)
}

func DeleteAPostByTitle(c *gin.Context) {
	var Blog models.Blog
	title := c.Param("title")

	claims, ok := c.Get("claims")
	if !ok {
		log.Fatal("No claims found")
	}
	claimsData := claims.(*jwt.RegisteredClaims)

	if err := database.DeleteAPostByTitleFromDB(&Blog, title, claimsData.Subject); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusFound, Blog)
}

func UpdateAPostByTitle(c *gin.Context) {
	title := c.Param("title")
	var toUpdateData models.Blog

	if err := c.BindJSON(&toUpdateData); err != nil {
		return
	}

	updatedPost, err := database.UpdateAPostByTitleFromDb(&toUpdateData, title)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedPost)
}
