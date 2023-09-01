package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("running...")

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware)

	// Routes
	e.GET("/n", daysLeft)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func daysLeft(ctx echo.Context) error {
	currentTime := time.Date(2025, time.January, 01, 00, 00, 00, 0, time.Local)
	deltaTime := time.Until(currentTime)
	daysLeftMessage := fmt.Sprintf("Days left: %d", int64(deltaTime.Hours()/24))
	return ctx.String(http.StatusOK, daysLeftMessage)
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		if ctx.Request().Header.Get("User-Role") == "admin" {
			log.Println("red-button user detected")
		}

		err := next(ctx)
		if err != nil {
			return nil
		}

		return nil
	}

}
