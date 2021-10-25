package routes

import (
	"kudos/endpoints/handlers"
	"kudos/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadRoutes() *echo.Echo {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Handlers
	personHandler := handlers.NewPersonHandler()

	// Routes
	e.GET("/kudos/people", personHandler.GetPeople)
	e.POST("/kudos/person/new", personHandler.CreatePerson)
	e.PUT("/kudos/person/:id/points/add", personHandler.AddPoints)
	e.GET("/kudos/person/:id/reais", personHandler.GetReais)
	e.GET("/kudos/person/:id/message", personHandler.GetPointsToMessage)
	e.GET("/kudos/convert/:points", utils.ConvertPointsToMessage)

	// Start server
	return e
}
