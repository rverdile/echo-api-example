package routes

import (
	"echo-api-example/pkg/db"
	"echo-api-example/pkg/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func InitUserRoutes(e *echo.Echo) {
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUserById)
	e.POST("/users", AddUser)
}

func GetUsers(c echo.Context) error {
	var users []models.User
	// users var is updated by reference
	if err := db.DB.Find(&users).Error; err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, users, "  ")
}

func GetUserById(c echo.Context) error {
	var user models.User
	var userID, err = strconv.Atoi(c.Param("id")) // string to int
	if err != nil {
		return err
	}
	if err := db.DB.First(&user, userID).Error; err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, user, "  ")
}

func AddUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := db.DB.Create(&user).Error; err != nil {
		return err
	}
	return c.String(http.StatusOK, "user created")
}
