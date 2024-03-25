package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/a-h/templ"
)

//go:generate nix run github:a-h/templ generate && nix run nixpkgs#gofumpt -- -w *.go
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/browse", BodyHandler)
	e.GET("/browse/:date", BodyHandler)
	e.Static("/dist", "dist")
	e.Static("/", "content")

	e.Logger.Fatal(e.Start(":8080"))
}

// https://github.com/a-h/templ/blob/main/examples/integration-echo/main.go
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, index("2016-02-02"))
}

func BodyHandler(c echo.Context) error {
	_, isHtmx := c.Request().Header["Hx-Request"]
	postdate := c.Param("date")
	if postdate == "" {
		// TODO: default value should be the latest issue
		// redirect to /browse/today?
		postdate = today()
	}
	if !isHtmx {
		return Render(c, http.StatusOK, index(postdate))
	}
	return Render(c, http.StatusOK, body(postdate))
}
