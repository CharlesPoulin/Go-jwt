package controllers

import (
	"Go-jwt/initializers"
	"Go-jwt/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Signup(c *gin.Context) {
	// get email // pass off req body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})

		return

	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})

		return
	}
	// create user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

	// return user
}

func Login(c *gin.Context) {
	// get email // pass off req body
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})

		return
	}

	// find user
	var user models.User
	initializers.DB.Where(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})

		return
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

	// generate a jwt token

	// send it back
}
