package main

import (
	"net/http"
        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"

	"github.com/a-h/templ"
)

func main() {
        e := echo.New()

        e.Use(middleware.Logger())
        e.Use(middleware.Recover())
        e.Use(middleware.Gzip())

        e.File("/browse", "dist/index.html")
	e.Static("/dist", "dist")
	e.Static("/", "content")

	e.POST("/browse", func(c echo.Context) error {
		return c.String(http.StatusOK, "not bad")
	})
	e.GET("/hello", HomeHandler)

        e.Logger.Fatal(e.Start(":8080"))
}

// https://github.com/a-h/templ/blob/main/examples/integration-echo/main.go
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
        ctx.Response().Writer.WriteHeader(statusCode)
        ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
        return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HomeHandler(c echo.Context) error {
        return Render(c, http.StatusOK, hello("pp"))
}
