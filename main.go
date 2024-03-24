package main

import (
        "net/http"

        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
)

func main() {
        e := echo.New()

        e.Use(middleware.Logger())
        e.Use(middleware.Recover())
        e.Use(middleware.Gzip())

        e.File("/", "dist/index.html")
	e.Static("/dist", "dist")

        e.POST("/clicked", clickedP)

        e.Logger.Fatal(e.Start(":8080"))
}

func clickedP(c echo.Context) error {
        return c.HTML(http.StatusOK, `
        <button id="baton" hx-post="/clicked" hx-swap="outerHTML">
                Click Me Again
        </button>`)
}

