package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"os"
	"pokemon-api/delivery/http/dto/request"
	"pokemon-api/delivery/http/dto/response"
	"pokemon-api/domain/repository"
)

type PokemonUsecase struct {
	DB                *gorm.DB
	Validate          *validator.Validate
	PokemonRepository *repository.PokemonRepository
}

func NewPokemonUsecase(DB *gorm.DB, validate *validator.Validate, pokemonRepository *repository.PokemonRepository) *PokemonUsecase {
	return &PokemonUsecase{DB: DB, Validate: validate, PokemonRepository: pokemonRepository}
}

func (p *PokemonUsecase) Create(ctx context.Context, request *request.PokemonRequest) (*response.PokemonResponse, error) {

	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := p.Validate.Struct(request)
	if err != nil {
		log.Printf("Invalid request body : %+v", err)
		return nil, err
	}

	// Save image to storage
	err = writeToFile(request.Image)
	if err != nil {
		return nil, err
	}

	imageUrl := request.BaseUrl + "/public/images/" + request.Image.Filename

	log.Println(imageUrl)

	pokemonEntity := request.ToEntity(imageUrl)
	err = p.PokemonRepository.PokemonCreate(tx, pokemonEntity)

	if err != nil {
		log.Printf("Usecase -> Failed to create pokemon : %+v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return pokemonEntity.ToResponse(), err
}

func (p *PokemonUsecase) FindAll(ctx context.Context) ([]*response.PokemonResponse, error) {
	tx := p.DB.WithContext(ctx)

	pokemons, err := p.PokemonRepository.PokemonAll(tx)

	if err != nil {
		return nil, err
	}

	var pokemonResponses []*response.PokemonResponse
	for _, pokemon := range pokemons {
		pokemonResponses = append(pokemonResponses, pokemon.ToResponse())
	}

	return pokemonResponses, nil
}

// Function to save image to storage
func writeToFile(f *multipart.FileHeader) error {
	// Save image to storage

	file, err := f.Open()

	if err != nil {
		return err
	}
	fileDestination, err := os.Create("./public/images/" + f.Filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileDestination, file)
	return err
}
