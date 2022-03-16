package middleware

import (
	"log"
	"net/http"

	"github.com/doffy007/price-comparison-api.git/helper"
	"github.com/doffy007/price-comparison-api.git/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//AuthorizeJWT is an validate token given to user and admin, retun 401 if not valid genetate token user and admin
func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[admin_id]", claims["admin_id"])
			log.Println("Claim[user_id],: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
