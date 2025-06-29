package controllers

import (
    "net/http"
    "time"
    "github.com/bchanona/profile_warmheart_backend/User/application"
    "github.com/bchanona/profile_warmheart_backend/User/domain"
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
    claims := &domain.Claims{
        User_id: user.User_id,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(ctrl.JWTKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    response := domain.LoginResponse{
        Token: tokenString,
        User:  user,
    }

    c.JSON(http.StatusOK, response)
}