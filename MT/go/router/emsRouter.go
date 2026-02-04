package router

import (
	middleware "ems/mt/golang/middleWare"
	userService "ems/mt/golang/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type user struct {
	EmployeeId int    `json:"employeeId"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	EmailId    string `json:"emailId"`
	MobileNo   string `json:"mobileNo"`
	Img        string `json:"img"`
	Password   string `json:"password"`
}

// var users = []user{
// 	{EmployeeId: 1, Name: "Alex Zhang", Position: "Developer", EmailId: "john@example.com", MobileNo: "1234567890", Img: "https://i.pravatar.cc/150?img=1"},
// 	{EmployeeId: 2, Name: "Mike Lee", Position: "Designer", EmailId: "jane@example.com", MobileNo: "0987654321", Img: "https://i.pravatar.cc/150?img=2"},
// 	{EmployeeId: 3, Name: "John Wang", Position: "Manager", EmailId: "mike@example.com", MobileNo: "1122334455", Img: "https://i.pravatar.cc/150?img=3"},
// }

type authRequest struct {
	EmailId  string `json:"emailId"`
	Password string `json:"password"`
}

type authResponse struct {
	Data token `json:"data"`
}

type token struct {
	Token string `json:"token"`
}

func SetupEMGRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: false,
		AllowOrigins:    []string{"http://localhost:4200", "http://localhost:9999"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization", "set-cookie"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// r.Use(middleware.HttpCookieAuthMiddleware)
	r.POST("/auth", authenticate)

	// r.GET("/user", getUsers)
	r.POST("/user", addUser)

	r.GET("/setcookies", setcookies)

	authRoutes := r.Group("/auth").Use(middleware.HttpCookieAuthMiddleware)
	{
		authRoutes.GET("/user", middleware.HttpCookieAuthMiddleware, getUsers)
	}

	// authRoutes.Use(middleware.HttpCookieAuthMiddleware)

	return r
}
func setcookies(c *gin.Context) {
	c.SetCookie("testcookie", "testcookievalue", 3600, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Test cookie set"})
}

type authUser struct {
	EmailId  string `json:"emailId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func authenticate(c *gin.Context) {

	var au authUser
	// Bind the request body to the 'content' struct
	if err := c.BindJSON(&au); err != nil {
		// If binding fails, return a 400 Bad Request error
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("====Authenticating user with emailId:", au.EmailId, " password:", au.Password)
	user, err := userService.GetUserByEmail(au.EmailId)
	if err != nil {
		fmt.Println("Error getting user by emailId:", err)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	fmt.Println("---------Authenticating user:", user)

	var tokenData = token{
		Token: "token_12345",
	}

	tokenData.Token = CreateJwtToken(user.Name)
	var authResponses = authResponse{
		Data: tokenData,
	}

	c.SetSameSite(http.SameSiteNoneMode)
	fmt.Println("===== Setting cookie with token:", tokenData.Token)
	c.SetCookie("authentication", tokenData.Token, 3600, "/", "localhost", false, true)
	c.SetCookie("testcookie", "testcookievalue======", 3600, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, authResponses)
}

func getUsers(c *gin.Context) {

	fmt.Println("---------Getting users list")

	dbUsers, err := userService.GetAllUsers()
	if err != nil {
		fmt.Println("Error getting all users:", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	var users []user

	for _, dbUser := range dbUsers {
		var u user
		u.EmployeeId = int(dbUser.ID)
		u.Name = dbUser.Name
		u.Position = "Developer"
		u.EmailId = dbUser.Email
		u.MobileNo = "1234567890"
		u.Img = dbUser.Img.String
		users = append(users, u)
	}

	c.IndentedJSON(http.StatusOK, users)
}

// func verifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("your_super_secure_and_long_secret_key"), nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}
// 	tokenClaims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return fmt.Errorf("invalid token claims")
// 	}

// 	fmt.Println("====Token Claims:", tokenClaims)

// 	return nil
// }

func addUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func CreateJwtToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"ttl":      time.Now().Add(time.Hour * 72).Unix(),
	})

	hmacSampleSecret := []byte("your_super_secure_and_long_secret_key")
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return ""
	}

	fmt.Println("Generated Token:", tokenString)
	return tokenString
}
