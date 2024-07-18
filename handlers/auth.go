package handlers

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"signage-system/firestore"
	"signage-system/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY"))

// Signup handles the user signup process.
// It hashes the user's password and stores the user data in Firestore.
func Signup(c echo.Context) error {
	user := new(models.User)
	// Bind the request body to the user struct
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Check if the email and password are provided
	if user.Email == "" || user.Password == "" || user.UserName == "" {
		return c.JSON(http.StatusBadRequest, "Email, password and user name are required")
	}

	// TODO: Check if user is already exist with given email address
	user.ID = uuid.New().String()

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(hashedPassword)

	// TODO: This can be a method in firestore utils as AddUser()
	// Add the user data to the Firestore collection
	_, _, err = firestore.Client.Collection("users").Add(context.Background(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// removing password from response
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

// Login handles the user login process.
// It checks the provided email and password against the stored data and returns a JWT token if successful.
func Login(c echo.Context) error {
	user := new(models.User)
	// Bind the request body to the user struct
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Check if the email and password are provided
	if user.Email == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, "Email and password are required")
	}

	// Query Firestore for the user with the provided email
	query := firestore.Client.Collection("users").Where("Email", "==", user.Email).Limit(1)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil || len(docs) == 0 {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	storedUser := new(models.User)
	// Map the Firestore document data to the storedUser struct
	docs[0].DataTo(storedUser)

	// Check if the provided password matches the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid email or password")
	}

	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = storedUser.ID
	claims["email"] = storedUser.Email

	// Sign the JWT token with the secret key
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Return the signed token
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
