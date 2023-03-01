package middleware

import (
    "fmt"
    "net/http"

	helper "gin-mongo-api/helpers"
	"github.com/rs/cors"



    "github.com/gin-gonic/gin"
)

// Authz validates token and authorizes users
func Authentication() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientToken := c.Request.Header.Get("token")
        if clientToken == "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
            c.Abort()
            return
        }

        claims, err := helper.ValidateToken(clientToken)
        if err != "" {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err})
            c.Abort()
            return
        }

        c.Set("email", claims.Email)
        c.Set("first_name", claims.First_name)
        c.Set("last_name", claims.Last_name)
        c.Set("uid", claims.Uid)

        c.Next()

    }
}

func CORSMiddleware(c *cors.Cors) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
