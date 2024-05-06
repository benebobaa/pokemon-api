package repository

import (
	"gorm.io/gorm"
	"log"
	"pokemon-api/domain/entity"
)

type FavoritePokemonRepository struct{}

func NewFavoritePokemonRepository() *FavoritePokemonRepository {
	return &FavoritePokemonRepository{}
}

func (f *FavoritePokemonRepository) FavoritePokemonCreate(tx *gorm.DB, value *entity.FavoritePokemon) error {
	result := tx.Create(value)

	if result.Error != nil {
		log.Printf("Repository -> Error when create favorite pokemon : %v", result.Error)
		return result.Error
	}

	return nil
}

func (f *FavoritePokemonRepository) DeleteByID(tx *gorm.DB, value int64) error {
	result := tx.Delete(&entity.FavoritePokemon{}, value)

	if result.Error != nil {
		log.Printf("Repository -> Error when delete favorite pokemon by id : %v", result.Error)
		return result.Error
	}

	return nil
}

func (f *FavoritePokemonRepository) FindAll(tx *gorm.DB) ([]entity.FavoritePokemon, error) {
	var favoritePokemons []entity.FavoritePokemon

	result := tx.Preload("Pokemon").Order("id DESC").Find(&favoritePokemons)

	if result.Error != nil {
		log.Printf("Repository -> Error when get all favorite pokemon : %v", result.Error)
		return nil, result.Error
	}

	return favoritePokemons, nil
}

func (f *FavoritePokemonRepository) UpdateFavorite(tx *gorm.DB, value *entity.FavoritePokemon) error {
	result := tx.Save(value)

	if result.Error != nil {
		log.Printf("Repository -> Error when update favorite pokemon : %v", result.Error)
		return result.Error
	}

	return nil
}

func (f *FavoritePokemonRepository) FindByID(tx *gorm.DB, value int64) (*entity.FavoritePokemon, error) {
	var favoritePokemon entity.FavoritePokemon

	result := tx.Preload("Pokemon").Where("id = ?", value).First(&favoritePokemon)

	if result.Error != nil {
		log.Printf("Repository -> Error when get favorite pokemon by pokemon id : %v", result.Error)
		return nil, result.Error
	}

	return &favoritePokemon, nil
}
