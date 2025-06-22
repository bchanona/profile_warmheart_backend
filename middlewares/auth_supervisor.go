package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Clave secreta para firmar el token del supervisor
var jwtSupervisorKey = []byte(os.Getenv("JWT_SUPERVISOR_SECRET"))

// Claims específicos para el supervisor
type SupervisorClaims struct {
	Supervisor_id int `json:"supervisor_idr"`
	User_id       int `json:"user_id"`
	jwt.StandardClaims
}

// Middleware para autenticar rutas de supervisor
func AuthSupervisorMiddleware() gin.HandlerFunc {
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
		claims := &SupervisorClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSupervisorKey, nil
		})

		// Validar si el token es válido
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		// Guardar el ID del supervisor en el contexto para usarlo en controladores
		ctx.Set("supervisor_id", claims.Supervisor_id)
		ctx.Set("user_id", claims.User_id)
		ctx.Next()
	}
}
