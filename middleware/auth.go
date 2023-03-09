package middleware

import (
	"fmt"
	"go-blogrpl/service"
	"go-blogrpl/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.CreateFailResponse("No token found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.CreateFailResponse("No token found", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.CreateFailResponse("Invalid token", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		if !token.Valid {
			response := utils.CreateFailResponse("Invalid token", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// get role from token
		roleRes, err := jwtService.GetRoleByToken(string(authHeader))
		fmt.Println("ROLE", roleRes)
		if err != nil || (roleRes != "admin" && roleRes != role) {
			response := utils.CreateFailResponse("Failed to process request", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusForbidden, response)
			return
		}

		// get userID from token
		idRes, err := jwtService.GetIDByToken(authHeader)
		if err != nil {
			response := utils.CreateFailResponse("Failed to process request", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println("ROLE", roleRes)
		c.Set("ID", idRes)
		c.Next()
	}
}
