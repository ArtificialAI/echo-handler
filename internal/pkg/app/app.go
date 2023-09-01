package app

import (
	"fmt"
	"log"

	"github.com/ArtificialAI/echo-handler/internal/app/endpoint"
	"github.com/ArtificialAI/echo-handler/internal/app/middleware"
	"github.com/ArtificialAI/echo-handler/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e          *endpoint.Endpoint
	s          *service.Service
	testServer *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.testServer = echo.New()

	a.testServer.Use(middleware.RoleCheck)

	a.testServer.GET("/n", a.e.Status)

	return a, nil

}

func (a *App) Run() error {
	fmt.Println("Server is running...")

	err := a.testServer.Start(":1323")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
