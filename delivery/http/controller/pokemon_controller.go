package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/rand"
	"log"
	"pokemon-api/delivery/http/dto/request"
	"pokemon-api/domain/usecase"
	"time"
)

type PokemonController struct {
	PokemonUsecase         *usecase.PokemonUsecase
	FavoritePokemonUsecase usecase.FavoritePokemonUsecase
}

func NewPokemonController(pokemonUsecase *usecase.PokemonUsecase) *PokemonController {
	return &PokemonController{PokemonUsecase: pokemonUsecase}
}

func (p *PokemonController) Create(ctx *fiber.Ctx) error {
	request := new(request.PokemonRequest)

	err := ctx.BodyParser(request)
	if err != nil {
		log.Printf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}
	log.Println(request)

	fileImage, err := ctx.FormFile("image")
	if err != nil {
		log.Printf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	request.Image = fileImage
	request.BaseUrl = ctx.BaseURL()

	response, err := p.PokemonUsecase.Create(ctx.UserContext(), request)
	if err != nil {
		log.Printf("Controller -> Failed to create pokemon : %+v", err)
		return fiber.ErrBadRequest
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})
}

func (p *PokemonController) FindAll(ctx *fiber.Ctx) error {

	response, err := p.PokemonUsecase.FindAll(ctx.UserContext())
	if err != nil {
		log.Printf("Controller -> Failed to get all pokemon : %+v", err)
		return fiber.ErrBadRequest
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    response,
	})
}

func (p *PokemonController) CatchPokemon(ctx *fiber.Ctx) error {
	rand.Seed(uint64(time.Now().UnixNano()))
	result := rand.Intn(2)

	log.Printf("Random number: %d", result)
	if result == 0 {
		return ctx.Status(200).JSON(fiber.Map{
			"message": "You failed to catch the pokemon",
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "You successfully catch the pokemon",
	})
}

func (p *PokemonController) ReleasePokemon(ctx *fiber.Ctx) error {
	rand.Seed(uint64(time.Now().UnixNano()))
	result := rand.Intn(100)

	log.Printf("Random number: %d", result)

	if isPrime(result) {
		return ctx.Status(201).JSON(fiber.Map{
			"message": "You successfully release the pokemon",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "You failed to release the pokemon",
	})
}

func (p *PokemonController) ServeImage(ctx *fiber.Ctx) error {

	params := ctx.Params("filename")

	if params == "" {
		return fiber.ErrBadRequest
	}

	return ctx.SendFile("./public/images/" + params)
}

func (p *PokemonController) HelloWorld(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"Hello": "World",
	})
}
