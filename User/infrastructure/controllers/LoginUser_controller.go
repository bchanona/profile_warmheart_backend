package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bchanona/profile_warmheart_backend/User/application"
	"github.com/bchanona/profile_warmheart_backend/User/domain"
	"github.com/bchanona/profile_warmheart_backend/User/infrastructure/adapters"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	UseCase *application.LoginUserUseCase
	JWTKey  []byte
}

func NewLoginUserController(useCase *application.LoginUserUseCase, jwtKey []byte) *LoginUserController {
	return &LoginUserController{UseCase: useCase, JWTKey: jwtKey}
}

type Claims struct {
	User_id int `json:"user_id"`
	jwt.StandardClaims
}

func (ctrl *LoginUserController) Login(c *gin.Context) {
	var login domain.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.UseCase.Login(login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		User_id: user.User_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Publicar datos por MQTT
	err = adapters.PublishUserData(user.User_id, user.Device_id)
	if err != nil {
		fmt.Println("Error al publicar datos a MQTT:", err)
		// Aquí solo se muestra el error, pero no detenemos el flujo
	}

	response := domain.LoginResponse{
		Token: tokenString,
		User:  user,
	}

	c.JSON(http.StatusOK, response)
}
