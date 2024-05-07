package middleware

import (
	"net/http"
	interfaces "scalexmedia-assignment/pkg/helper/interface"
	"scalexmedia-assignment/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

var (
	authHelper interfaces.HelperAuth
)

func SetHelper(helper interfaces.HelperAuth) {
	authHelper = helper
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHelper.GetTokenFromHeader(authHeader)
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		adminEmail, adminName, err := authHelper.ExtractAdminFromToken(tokenString)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("adminname", adminName)
		c.Set("adminemail", adminEmail)
		c.Next()
	}
}

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := authHelper.GetTokenFromHeader(authHeader)
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		userEmail, userName, err := authHelper.ExtractUserDetailsFromToken(tokenString)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("username", userName)
		c.Set("useremail", userEmail)
		c.Next()
	}
}
