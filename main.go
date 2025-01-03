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
	// client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	// d := client.Database("battle_tracker")
	// cs := campaigns.NewCampaignService(d)
	// c, _ := cs.HandleGetCampaigns()
	// fmt.Println(c)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "Index page") })
	g := e.Group("/api")
	api.NewEchoRouter(g)
	e.Logger.Fatal(e.Start(":3000"))
}
