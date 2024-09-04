package middlewares

import (
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CheckAdmin middleware ahli admin - lere dostup beryar
// gelen request - in admin tarapyndan gelip gelmedigini barlayar
// we admin bolmasa gecirmeyar
func CheckToken(forAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Token is required", "status": false})
			return
		}
		var tokenString string

		splitToken := strings.Split(tokenStr, "Bearer ")
		if len(splitToken) > 1 {
			tokenString = splitToken[1]
		} else {
			c.AbortWithStatusJSON(400, gin.H{"message": "Invalid token", "status": false})
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&helpers.JWTClaimForAdmin{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(helpers.JwtKey), nil
			},
		)
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"message": err.Error(), "status": false})
			return
		}
		claims, ok := token.Claims.(*helpers.JWTClaimForAdmin)
		if !ok {
			c.AbortWithStatusJSON(400, gin.H{"message": "couldn't parse claims", "status": false})
			return
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			c.AbortWithStatusJSON(403, gin.H{"message": "token expired", "status": false})
			return
		}

		db, err := config.ConnDB()
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": err.Error(), "status": false})
			return
		}
		defer db.Close()

		tableName := "admins"
		if !forAdmin {
			tableName = "customers"
		}

		if err := helpers.ValidateRecordByID(tableName, claims.ID, "NULL", db); err != nil {
			c.AbortWithStatusJSON(404, gin.H{"message": "record not found", "status": false})
			return
		}

		c.Set("id", claims.ID)
		c.Next()
	}
}
