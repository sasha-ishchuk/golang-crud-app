package controllers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang-crud-app/database"
	"golang-crud-app/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"time"
)

func Register(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	isEmailValid, _ := regexp.MatchString(emailPattern, user.Email)
	if !isEmailValid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid email address"})
	}

	if len(user.Password) < 6 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password must be at least 6 characters"})
	}

	var existingUser models.User
	database.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		return c.JSON(http.StatusConflict, map[string]string{"error": "User already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not hash password"})
	}

	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create user"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func Login(c echo.Context) error {
	loginData := new(models.User)
	if err := c.Bind(loginData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	var user models.User
	database.DB.Where("email = ?", loginData.Email).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	token := generateSessionToken()
	user.Token = token

	if err := database.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update token"})
	}

	c.SetCookie(&http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful", "token": token})
}

func Logout(c echo.Context) error {
	cookie, err := c.Cookie("auth_token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "No token found"})
	}

	var user models.User
	database.DB.Where("token = ?", cookie.Value).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
	}

	// Invalidate the token
	user.Token = ""
	if err := database.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not update user"})
	}

	c.SetCookie(&http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth_token")
		if err != nil || cookie.Value == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}

		var user models.User
		database.DB.Where("token = ?", cookie.Value).First(&user)
		if user.ID == 0 {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		}

		c.Set("user", user)
		return next(c)
	}
}

func GetUser(c echo.Context) error {
	user := c.Get("user").(models.User)
	return c.JSON(http.StatusOK, user)
}

func generateSessionToken() string {
	return uuid.New().String()
}
