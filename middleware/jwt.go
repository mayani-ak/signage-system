package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// JWT returns a JWT middleware function that validates JWT tokens using the secret key.
func JWT() echo.MiddlewareFunc {
	// Check if the JWT secret key is provided
	if len(jwtSecret) == 0 {
		panic("JWT_SECRET_KEY environment variable is missing or empty")
	}

	// Return the JWT middleware configuration
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtSecret,
	})
}
