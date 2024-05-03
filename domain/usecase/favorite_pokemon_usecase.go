package usecase

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"pokemon-api/delivery/http/dto/request"
	"pokemon-api/delivery/http/dto/response"
	"pokemon-api/domain/repository"
	"strconv"
)

type FavoritePokemonUsecase struct {
	DB                        *gorm.DB
	Validate                  *validator.Validate
	PokemonRepository         *repository.PokemonRepository
	FavoritePokemonRepository *repository.FavoritePokemonRepository
}

func NewFavoritePokemonUsecase(DB *gorm.DB, validate *validator.Validate, pokemonRepository *repository.PokemonRepository, favoritePokemonRepository *repository.FavoritePokemonRepository) *FavoritePokemonUsecase {
	return &FavoritePokemonUsecase{DB: DB, Validate: validate, PokemonRepository: pokemonRepository, FavoritePokemonRepository: favoritePokemonRepository}
}

func (p *FavoritePokemonUsecase) AddFavorite(ctx context.Context, request *request.FavoritePokemonRequest) (*response.FavoritePokemonResponse, error) {

	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := p.Validate.Struct(request)
	if err != nil {
		log.Printf("Invalid request body : %+v", err)
		return nil, err
	}

	pokemonEntity, err := p.PokemonRepository.PokemonFindById(tx, request.PokemonID)
	if err != nil {
		log.Printf("Usecase -> Failed to get pokemon by id : %+v", err)
		return nil, err
	}

	request.NickName = request.NickName + "-" + strconv.Itoa(fibonacci(0))

	favoritePokemonEntity := request.ToEntity(pokemonEntity)
	err = p.FavoritePokemonRepository.FavoritePokemonCreate(tx, favoritePokemonEntity)
	if err != nil {
		log.Printf("Usecase -> Failed to create favorite pokemon : %+v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return favoritePokemonEntity.ToResponse(), err
}

func (p *FavoritePokemonUsecase) ReleaseFromFavorite(ctx context.Context, id int64) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := p.FavoritePokemonRepository.DeleteByID(tx, id)
	if err != nil {
		log.Printf("Usecase -> Failed to delete favorite pokemon by id : %+v", err)
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (p *FavoritePokemonUsecase) FindAll(ctx context.Context) ([]*response.FavoritePokemonResponse, error) {
	tx := p.DB.WithContext(ctx)

	favoritePokemons, err := p.FavoritePokemonRepository.FindAll(tx)
	if err != nil {
		return nil, err
	}

	var favoritePokemonResponses []*response.FavoritePokemonResponse
	for _, favoritePokemon := range favoritePokemons {
		favoritePokemonResponses = append(favoritePokemonResponses, favoritePokemon.ToResponse())
	}

	return favoritePokemonResponses, nil
}

func (p *FavoritePokemonUsecase) UpdateNickname(ctx context.Context, request *request.FavoriteUpdateRequest) (*response.FavoritePokemonResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := p.Validate.Struct(request)
	if err != nil {
		log.Printf("Invalid request body : %+v", err)
		return nil, err
	}

	favoritePokemonEntity, err := p.FavoritePokemonRepository.FindByID(tx, request.ID)

	if err != nil {
		log.Printf("Usecase -> Failed to get favorite pokemon by id : %+v", err)
		return nil, err
	}

	favoritePokemonEntity.CountUpdate += 1
	favoritePokemonEntity.NickName = request.NickName + "-" + strconv.Itoa(fibonacci(favoritePokemonEntity.CountUpdate))

	err = p.FavoritePokemonRepository.UpdateFavorite(tx, favoritePokemonEntity)
	if err != nil {
		log.Printf("Usecase -> Failed to update favorite pokemon : %+v", err)
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return favoritePokemonEntity.ToResponse(), err
}

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		fibPrev := 0
		fibCurrent := 1
		for i := 2; i <= n; i++ {
			fibPrev, fibCurrent = fibCurrent, fibPrev+fibCurrent
		}
		return fibCurrent
	}
}
