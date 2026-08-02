package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() { // FIX 1: Rename to main() so Go can execute it
	initDBs()

	r := gin.Default()

	// Load all template files inside assets folder
	r.LoadHTMLGlob("assets/*")

	// -------------------------------------------------------------------------
	// PUBLIC ROUTES (Serving Pages)
	// -------------------------------------------------------------------------

	// Home / Dashboard Page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	// Login & Signup HTML Pages (Uses GET)
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// Auth API Endpoint Handlers (Form Submissions)
	r.POST("/login", apiLogin)
	r.POST("/signup", apiRegister)
	r.GET("/logout", apiLogout)

	// -------------------------------------------------------------------------
	// PROTECTED ROUTES (Requires Authentication)
	// -------------------------------------------------------------------------
	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/home", func(c *gin.Context) {
			c.HTML(http.StatusOK, "home.html", nil)
		})
		protected.GET("/encode", func(c *gin.Context) {
			c.HTML(http.StatusOK, "encode.html", nil)
		})

		protected.GET("/decode", func(c *gin.Context) {
			c.HTML(http.StatusOK, "decode.html", nil)
		})

		// FIX 3: Process Cipher submissions on POST
		protected.POST("/encode", apiEncode)
		protected.POST("/decode", apiDecode)
	}

	// FIX 4: Correct log port to match 7777
	log.Println("Server running on http://localhost:7777")
	r.Run(":7777")
}
