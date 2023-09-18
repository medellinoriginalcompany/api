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
		FullName  string
		Phone     string
		CPF       string
		Email     string
		Password  string
		BirthDate string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// Verificar se o usuário já existe
	var user models.Customer
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email/CPF já cadastrado",
		})

		return
	}

	// Hash senha
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	// Criar usuário
	user = models.Customer{
		FullName:  body.FullName,
		Phone:     body.Phone,
		CPF:       body.CPF,
		Email:     body.Email,
		Password:  string(hash),
		BirthDate: body.BirthDate,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao criar usuário",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Cadastro realizado com sucesso",
		"user":    &user,
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
			"message": "Failed to read body",
		})

		return
	}

	// Procurar usuário
	var user models.Customer
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email não cadastrado",
		})

		return
	}

	// Comparar hash com senha
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Senha incorreta",
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
			"message": "Failed to create token",
		})

		return
	}

	// send it back as response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("auth_token", tokenString, 3600*24*60, "", "", true, true) // 2 meses
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "", "", true, true) // Apaga o cookie setando para um tempo negativo
	c.JSON(http.StatusOK, gin.H{
		"token": "Cookie deleted.",
	})
}

func AdminLogout(c *gin.Context) {
	cookie := config.Getenv("ADMIN_TOKEN")
	c.SetCookie(cookie, "", -1, "", "", true, true) // Apaga o cookie setando para um tempo negativo
	c.JSON(http.StatusOK, gin.H{
		"token": "Cookie deleted.",
	})
}

func AdminLogin(c *gin.Context) {
	// Get email e senha
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// Procurar usuário
	var user models.Admin
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email ou senha incorretos",
		})

		return
	}

	// Comparar hash com senha
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email ou senha incorretos",
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
			"message": "Failed to create token",
		})

		return
	}

	// send it back as response
	c.SetSameSite(http.SameSiteLaxMode)
	cookie := config.Getenv("ADMIN_TOKEN")
	c.SetCookie(cookie, tokenString, 3600*24*7, "", "", true, true) // 1 semana
	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": tokenString,
	})
}

func AdminSignup(c *gin.Context) {
	// Pegar info do usuário do corpo da req
	var body struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})

		return
	}

	// Verificar se o usuário já existe
	var user models.Admin
	database.DB.First(&user, "email = ?", body.Email)

	if user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email já cadastrado",
		})

		return
	}

	// Hash senha
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash password",
		})
		return
	}

	// Criar usuário
	user = models.Admin{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  string(hash),
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Cadastro realizado com sucesso",
	})
}
