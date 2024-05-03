package repository

import (
	"gorm.io/gorm"
	"log"
	"pokemon-api/domain/entity"
)

type PokemonRepository struct{}

func NewPokemonRepository() *PokemonRepository {
	return &PokemonRepository{}
}

func (p *PokemonRepository) PokemonCreate(tx *gorm.DB, value *entity.Pokemon) error {
	result := tx.Create(value)

	if result.Error != nil {
		log.Printf("Repository -> Error when create pokemon : %v", result.Error)
		return result.Error
	}

	return nil
}

func (p *PokemonRepository) PokemonFindById(tx *gorm.DB, id int64) (*entity.Pokemon, error) {
	var pokemon entity.Pokemon

	result := tx.Where("id = ?", id).First(&pokemon)

	if result.Error != nil {
		log.Printf("Repository -> Error when get pokemon by id : %v", result.Error)
		return nil, result.Error
	}

	return &pokemon, nil
}

func (p *PokemonRepository) PokemonAll(tx *gorm.DB) ([]entity.Pokemon, error) {
	var pokemons []entity.Pokemon

	result := tx.Find(&pokemons)

	if result.Error != nil {
		log.Printf("Repository -> Error when get all pokemon : %v", result.Error)
		return nil, result.Error
	}

	return pokemons, nil

}
