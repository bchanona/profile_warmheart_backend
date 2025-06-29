package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY_USER"))

type Claims struct {
	User_id int `json:"user_id"`
	jwt.StandardClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		// Verificar si el token existe y tiene el formato correcto
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token no proporcionado"})
			ctx.Abort()
			return
		}

		// Extraer el token quitando "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Parsear el token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Validar si el token es válido
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		// Guardar el ID del usuario en el contexto para usarlo en controladores
		ctx.Set("user_id", claims.User_id)
		ctx.Next()
	}
}