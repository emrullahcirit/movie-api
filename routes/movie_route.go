package routes

import (
	"github.com/emrullahcirit/movie-api/controllers"
	"github.com/gofiber/fiber/v2"
)	

func UserRoute(app *fiber.App) {
	app.Post("/movie", controllers.CreateMovie)
	app.Get("/movie/:movieId", controllers.GetAMovie)
	app.Get("/movies", controllers.GetAllMovies)
	app.Put("/movie/:movieId", controllers.EditAMovie)
	app.Delete("/movie/:movieId", controllers.DeleteAMovie)
}
