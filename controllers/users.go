package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/medellinoriginalcompany/api/config"
	"github.com/medellinoriginalcompany/api/database"
	"github.com/medellinoriginalcompany/api/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Pegar info do usuário do corpo da req
	var body struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
		Birth     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Failed to read body",
		})

		return
	}

	// Hash senha
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Failed to hash password",
		})
		return
	}


	// Criar usuário
	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash),
		Birth:     body.Birth,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Failed to create user",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Success": "Account create successfully",
		"user": &user,
	})
}


func Login(c *gin.Context) {
	// Get email e senha
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Failed to read body",
		})

		return
	}

	// Procurar usuário
	var user models.User
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Invalid email or password",
		})

		return
	}

	// Comparar hash com senha
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Invalid email or password",
		})

		return
	}

	// Gerar jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": "Failed to create token",
		})

		return
	}

	// Criar cookie com o token de sessão que expira em 30 dias
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true) 
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	id := user.(models.User).ID // Pegar informação específica do usuário

	c.JSON(http.StatusOK, gin.H{
		"message": user,
		"messag":  id,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "", "", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}