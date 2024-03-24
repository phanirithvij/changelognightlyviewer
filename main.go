package main

import (
	"net/http"

        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"

	"github.com/a-h/templ"
)

//go:generate nix run github:a-h/templ generate
func main() {
        e := echo.New()

        e.Use(middleware.Logger())
        e.Use(middleware.Recover())
        e.Use(middleware.Gzip())

        e.GET("/browse", HomeHandler)
	e.Static("/dist", "dist")
	e.Static("/", "content")

	e.POST("/bod", BodyHandler)
	e.GET("/bod", BodyHandler)

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

// TODO: combine this and homehandler and use /browse?date=...
// check if req sent by htmx and serve partial body
func BodyHandler(c echo.Context) error {
	if (c.Request().Method == "GET") {
	        return Render(c, http.StatusOK, body("2016-02-02"))
	}
	var bod BodyRecv
	err := c.Bind(&bod); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	return Render(c, http.StatusOK, body(bod.PostDate))
}

type BodyRecv struct {
	PostDate string `form:"postdate"`
}

