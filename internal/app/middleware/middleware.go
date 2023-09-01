package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		value := ctx.Request().Header.Get("User-Role")
		if value == roleAdmin {
			log.Println("red-button user detected")
		}
		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
