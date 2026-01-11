package main

import (
	"github.com/gin-gonic/gin"
)

// ---------------------------------------------------------------- //

func setup_gin() {
	r := gin.Default()

	// Middleware: Secure HTTP headers
	r.Use(func(c *gin.Context) {
		c.Header("Content-Security-Policy",
			"default-src 'self'; "+
				"script-src 'self' https://cdn.termly.io https://app.termly.io; "+
				"style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; "+
				"font-src 'self' https://fonts.gstatic.com; "+
				"object-src 'none';")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "no-referrer")
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		c.Next()
	})

	// Middleware: Redirect HTTP to HTTPS (if behind a proxy)
	r.Use(func(c *gin.Context) {
		if c.Request.Header.Get("X-Forwarded-Proto") == "http" {
			url := c.Request.URL
			url.Scheme = "https"
			url.Host = c.Request.Host
			c.Redirect(301, url.String())
			c.Abort()
			return
		}
		c.Next()
	})

	r.LoadHTMLGlob("templates/*.html")

	// Serve static files safely
	r.Static("/static", "./static")

	// Redirect root URL to /home for a user-friendly homepage
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/home")
	})

	r.GET("/home", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Ridge Software Dynamics",
		})
	})

	r.GET("/static_construct_lucy", func(c *gin.Context) {
		c.HTML(200, "lucy_construct.html", gin.H{
			"title": "Static Construct - Lucy",
		})

	})

	r.GET("/research", func(c *gin.Context) {
		c.HTML(200, "research.html", gin.H{
			"title": "Ridge Software Dynamics Lab",
		})
	})

	r.Run(":8080") // runs on localhost:8080
}

func main() {
	setup_gin()
}
