package main

import (
	"battle_tracker/api"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// apiRouter := api.NewApiRouter()
	// // socketServer := socketio.NewServer(nil)
	//
	// http.Handle("/api/", apiRouter)
	// log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Index page") })
	g := e.Group("/api")
	api.NewEchoRouter(g)
	e.Logger.Fatal(e.Start(":3000"))
}
