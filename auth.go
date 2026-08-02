package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type DBManager struct {
	UserDB   *sql.DB
	CipherDB *sql.DB
}

var DBs DBManager

// Helper function to read environment variables with a fallback default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

var jwtKey = []byte(getEnv("JWT_KEY", "cipher_project"))

// JWT Claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Generates a signed JWT token string for a user
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Reads and verifies the JWT stored in the client's cookie
func getUserIDFromCookie(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "", err
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid or expired token")
	}

	return claims.Username, nil
}

// AuthMiddleware protects routes requiring authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := getUserIDFromCookie(c)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Please Login First!"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}

// Helper functions for Password Hashing
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// API Register
func apiRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "Username and password are required"})
		return
	}

	hashbytes, err := HashPassword(password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Server error processing password"})
		return
	}

	_, err = DBs.UserDB.Exec("INSERT INTO app_user (username, password) VALUES ($1, $2)", username, hashbytes)
	if err != nil {
		fmt.Println("Register DB Error:", err)
		c.HTML(http.StatusConflict, "signup.html", gin.H{"error": "Username already taken"})
		return
	}

	// Successfully registered -> Redirect user directly to Login page
	c.Redirect(http.StatusSeeOther, "/login")
}

// API Login
func apiLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 1. Basic validation
	if username == "" || password == "" {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Please fill in all fields"})
		return
	}

	var userID int
	var userPassword string

	// 2. Query user credentials from app_user database
	err := DBs.UserDB.QueryRow("SELECT id, password FROM app_user WHERE username = $1", username).Scan(&userID, &userPassword)

	// 3. Check if user exists and password hash matches
	if err != nil || !CheckPasswordHash(password, userPassword) {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid username or password"})
		return
	}

	// 4. Generate JWT token
	tokenString, err := GenerateJWT(username)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Failed to generate session token"})
		return
	}

	// 5. Set session cookie & redirect to the main app page
	c.SetCookie("token", tokenString, 86400, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/home")
}

// API Logout
func apiLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/login")
}
