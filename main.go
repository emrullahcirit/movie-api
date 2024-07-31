package main

import (
	"github.com/emrullahcirit/movie-api/configs"
	"github.com/emrullahcirit/movie-api/routes"
  "github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  //run db
  configs.ConnectDB()

	//routes
	routes.UserRoute(app)

  app.Listen(":6000")
}
