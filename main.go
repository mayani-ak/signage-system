package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"signage-system/firestore"
	"signage-system/handlers"
	mdwr "signage-system/middleware"
	"signage-system/views"
)

func main() {
	// Initialize Firestore connection
	firestore.InitFirestore()

	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Log HTTP requests
	e.Use(middleware.Recover()) // Recover from panics gracefully

	// Register the template renderer for HTML views
	e.Renderer = views.NewTemplate()

	// Public routes (no authentication required)
	e.POST("/signup", handlers.Signup) // User signup endpoint
	e.POST("/login", handlers.Login)   // User login endpoint

	// Protected routes (require JWT authentication)
	r := e.Group("/restricted")
	r.Use(mdwr.JWT()) // Middleware to verify JWT tokens

	// Views and handlers
	r.GET("/", views.IndexView)               // Homepage
	r.GET("/content", handlers.GetContent)    // Endpoint to fetch content
	r.POST("/update", handlers.UpdateContent) // Endpoint to update content

	// Start the server
	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
}
