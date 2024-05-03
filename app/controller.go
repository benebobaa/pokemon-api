package app

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"pokemon-api/delivery/http/controller"
	"pokemon-api/domain/repository"
	"pokemon-api/domain/usecase"
)

type Controller struct {
	PokemonController    *controller.PokemonController
	FavPokemonController *controller.FavoritePokemonController
}

func NewController(db *gorm.DB, validate *validator.Validate) *Controller {

	// Setup repositories
	pokemonRepository := repository.NewPokemonRepository()
	favPokemonRepository := repository.NewFavoritePokemonRepository()

	// Setup usecases
	pokemonUsecase := usecase.NewPokemonUsecase(db, validate, pokemonRepository)
	favPokemonUsecase := usecase.NewFavoritePokemonUsecase(db, validate, pokemonRepository, favPokemonRepository)

	// Setup controller
	pokemonController := controller.NewPokemonController(pokemonUsecase)
	favPokemonController := controller.NewFavoritePokemonController(favPokemonUsecase)

	// Setup middleware
	return &Controller{
		PokemonController:    pokemonController,
		FavPokemonController: favPokemonController,
	}
}
