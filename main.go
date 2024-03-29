package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/a-h/templ"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "Port number to listen on")
	help := flag.Bool("help", false, "Print usage information")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
}

//go:generate nix run github:a-h/templ generate
//go:generate nix run nixpkgs#gofumpt -- -w .
func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	e.GET("/browse", BodyHandler)
	e.GET("/browse/", BodyHandler)
	e.GET("/browsehtml", BodyHandlerNoJs)
	e.GET("/browsehtml/", BodyHandlerNoJs)

	e.Static("/dist", "dist")
	e.Static("/", "content")

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}

// https://github.com/a-h/templ/blob/main/examples/integration-echo/main.go
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func BodyHandler(c echo.Context) error {
	_, isHtmx := c.Request().Header["Hx-Request"]
	jsenabled := true
	postdate := c.QueryParam("postdate")
	if postdate == "" {
		postdate = today()
		return c.Redirect(307, "/browse?postdate="+postdate)
	}
	if !isHtmx {
		return Render(c, http.StatusOK, index(postdate, jsenabled))
	}
	return Render(c, http.StatusOK, body(postdate, jsenabled))
}

func BodyHandlerNoJs(c echo.Context) error {
	postdate := c.QueryParam("postdate")
	if postdate == "" {
		postdate = today()
		return c.Redirect(307, "/browsehtml?postdate="+postdate)
	}
	jsenabled := false
	return Render(c, http.StatusOK, index(postdate, jsenabled))
}
