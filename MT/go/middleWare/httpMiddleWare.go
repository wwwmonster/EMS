package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func HttpCookieAuthMiddleware(c *gin.Context) {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication cookie"})
		c.Abort()
		return
	}
	fmt.Println("==11==authCookie value:", cookie)

	if cookie == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		return
	}

	// cookie = cookie[len("Bearer "):]
	// fmt.Println("==11==Extracted Token value:", cookie)

	err_verifyToken := verifyToken(cookie)
	if err_verifyToken != nil {
		fmt.Println("==11==Token verification failed:", err_verifyToken)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})

		return
	}

	// Token validation logic can be added here
	c.Next()
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your_super_secure_and_long_secret_key"), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid token claims")
	}

	fmt.Println("==11==Token Claims:", tokenClaims)

	return nil
}
