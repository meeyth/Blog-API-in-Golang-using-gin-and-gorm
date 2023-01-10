package controllers

import (
	"net/http"
	"os"
	"social-media/database"
	"social-media/models"
	"social-media/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var body database.ReqBody

	if err := ctx.BindJSON(&body); err != nil {
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		panic(err)
	}

	birthDay, err := utils.CreateDate(body.Birthday)
	if err != nil {
		panic(err)
	}

	var newUser = models.User{
		UserName: body.UserName,
		Email:    body.Email,
		Password: hashedPassword,
		Birthday: birthDay,
		JoinedAt: time.Now(),
	}

	if err := database.InsertAUserIntoDb(&newUser); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "This username is already taken",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, newUser)
}

func Login(ctx *gin.Context) {
	var body database.ReqBody

	if err := ctx.BindJSON(&body); err != nil {
		return
	}

	var user models.User

	database.GetLoggedInUserFromDb(&user, body.UserName)

	if user.ID == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(body.Password)); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Password",
		})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		Subject:   user.Email,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Something wen wrong in the server, please try again after a while",
		})
		return
	}

	ctx.SetCookie("jwt", token, 3600*24, "/", "localhost", false, true)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func Logout(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "localhost", false, true)

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
