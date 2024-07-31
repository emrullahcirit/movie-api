package controllers

import (
    "context"
    "github.com/emrullahcirit/movie-api/configs"
    "github.com/emrullahcirit/movie-api/models"
    "github.com/emrullahcirit/movie-api/responses"
    "net/http"
    "time"
  
    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
		"go.mongodb.org/mongo-driver/bson"
)

var movieCollection *mongo.Collection = configs.GetCollection(configs.DB, "movies")
var validate = validator.New()

func CreateMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movie models.Movie
	defer cancel()

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MovieResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&movie); validationErr != nil {
	    return c.Status(http.StatusBadRequest).JSON(responses.MovieResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newMovie := models.Movie{
		Id:				primitive.NewObjectID(),
		Name:			movie.Name,
		Director:	movie.Director,
		Rating:		movie.Rating,
		Comment:	movie.Comment,
	}

	result, err := movieCollection.InsertOne(ctx, newMovie)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.MovieResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
	
}

func GetAMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	var movie models.Movie
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	err := movieCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&movie)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": movie}})

}

func GetAllMovies(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var movies []models.Movie
	defer cancel()

	results, err := movieCollection.Find(ctx, bson.M{})

	if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	
	defer results.Close(ctx)
	
	for results.Next(ctx) {
		var singleMovie models.Movie
		if err = results.Decode(&singleMovie); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		movies = append(movies, singleMovie)

	}

	return c.Status(http.StatusOK).JSON(
		responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": movies}},
	)

}

func EditAMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	var movie models.Movie
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	if err := c.BodyParser(&movie); err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.MovieResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&movie); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MovieResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"name": movie.Name, "director": movie.Director, "rating": movie.Rating, "comment": movie.Comment}

	result, err := movieCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	var updatedMovie models.Movie
	if result.MatchedCount == 1 {
		err := movieCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedMovie)
		
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": updatedMovie}})

}

func DeleteAMovie(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	movieId := c.Params("movieId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(movieId)

	result, err := movieCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.MovieResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.MovieResponse{Status: http.StatusNotFound, Message: "error", Data: &fiber.Map{"data": "Movie with specified ID not found!"}},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.MovieResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": "Movie successfully deleted!"}},
	)
}