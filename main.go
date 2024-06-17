package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/students", getStudents)
  e.GET("/students/:id", getStudent)
  e.POST("/students", createStudent)
  e.PUT("/students/:id", updateStudent)
  e.DELETE("/students/:id", deleteStudent)
  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "List of all students")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	getStuden := fmt.Sprintf("Get %s student", id)
	return c.String(http.StatusOK, getStuden)
}

func createStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Create students")
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	UpdateStuden := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, UpdateStuden)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStuden := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, deleteStuden)
}

