package app

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, c *Controller) {
	app.Get("/api/hello", c.PokemonController.HelloWorld)

	//Pokemon
	app.Post("/api/pokemon", c.PokemonController.Create)
	app.Get("/api/pokemon", c.PokemonController.FindAll)

	//Fav Pokemon
	app.Post("/api/pokemon/favorite", c.FavPokemonController.Create)
	app.Delete("/api/pokemon/favorite/:id", c.FavPokemonController.Release)
	app.Get("/api/pokemon/favorite", c.FavPokemonController.FindAll)
	app.Put("/api/pokemon/favorite/:id", c.FavPokemonController.UpdateFavorite)
}
