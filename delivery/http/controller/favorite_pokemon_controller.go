package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/rand"
	"log"
	"pokemon-api/delivery/http/dto/request"
	"pokemon-api/domain/usecase"
	"time"
)

type FavoritePokemonController struct {
	FavPokemonUsecase *usecase.FavoritePokemonUsecase
}

func NewFavoritePokemonController(favPokemonUsecase *usecase.FavoritePokemonUsecase) *FavoritePokemonController {
	return &FavoritePokemonController{FavPokemonUsecase: favPokemonUsecase}
}

func (f *FavoritePokemonController) Create(ctx *fiber.Ctx) error {
	rand.Seed(uint64(time.Now().UnixNano()))
	result := rand.Intn(2)

	log.Printf("Random number: %d", result)
	if result == 0 {
		return ctx.Status(200).JSON(fiber.Map{
			"message": "You failed to catch the pokemon",
		})
	}

	request := new(request.FavoritePokemonRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		log.Printf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}
	log.Println(request)

	response, err := f.FavPokemonUsecase.AddFavorite(ctx.UserContext(), request)
	if err != nil {
		log.Printf("Controller -> Failed to add new favorite : %+v", err)
		return fiber.ErrBadRequest
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})
}

func (f *FavoritePokemonController) Release(ctx *fiber.Ctx) error {
	rand.Seed(uint64(time.Now().UnixNano()))
	result := rand.Intn(100)

	log.Printf("Random number: %d", result)

	if !isPrime(result) {
		return ctx.Status(200).JSON(fiber.Map{
			"message": "You failed to release the pokemon",
		})
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Printf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	err = f.FavPokemonUsecase.ReleaseFromFavorite(ctx.UserContext(), int64(id))
	if err != nil {
		log.Printf("Controller -> Failed to release favorite : %+v", err)
		return fiber.ErrBadRequest
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "Success",
	})
}

func (f *FavoritePokemonController) FindAll(ctx *fiber.Ctx) error {

	response, err := f.FavPokemonUsecase.FindAll(ctx.UserContext())
	if err != nil {
		log.Printf("Controller -> Failed to get all favorite pokemon : %+v", err)
		return fiber.ErrBadRequest
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})
}

func (f *FavoritePokemonController) UpdateFavorite(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Printf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	favRequest := new(request.FavoriteUpdateRequest)

	err = ctx.BodyParser(favRequest)
	if err != nil {
		log.Printf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	favRequest.ID = int64(id)
	response, err := f.FavPokemonUsecase.UpdateNickname(ctx.UserContext(), favRequest)
	if err != nil {
		log.Printf("Controller -> Failed to get favorite pokemon by id : %+v", err)
		return fiber.ErrBadRequest
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})
}

func isPrime(num int) bool {
	if num <= 1 {
		return false // Numbers less than or equal to 1 are not prime
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false // Found a factor, not prime
		}
	}
	return true // No factors found, prime
}
