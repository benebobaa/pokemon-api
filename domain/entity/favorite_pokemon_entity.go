package entity

import "pokemon-api/delivery/http/dto/response"

type FavoritePokemon struct {
	ID          int64  `gorm:"column:id;primaryKey;"`
	NickName    string `gorm:"column:name"`
	CountUpdate int    `gorm:"column:count_update"`
	PokemonID   int64  // Foreign key
	Pokemon     Pokemon
}

func (f *FavoritePokemon) TableName() string {
	return "favorite_pokemons"
}

func (f *FavoritePokemon) ToResponse() *response.FavoritePokemonResponse {
	return &response.FavoritePokemonResponse{
		ID:          f.ID,
		NickName:    f.NickName,
		Pokemon:     *f.Pokemon.ToResponse(),
		CountUpdate: f.CountUpdate,
	}
}
