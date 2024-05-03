package request

import "pokemon-api/domain/entity"

type FavoritePokemonRequest struct {
	PokemonID int64  `json:"pokemon_id" validate:"required"`
	NickName  string `json:"nick_name" validate:"required"`
}

type FavoriteUpdateRequest struct {
	ID       int64  `json:"id" validate:"required"`
	NickName string `json:"nick_name" validate:"required"`
}

func (f *FavoritePokemonRequest) ToEntity(p *entity.Pokemon) *entity.FavoritePokemon {
	return &entity.FavoritePokemon{
		NickName:    f.NickName,
		CountUpdate: 0,
		Pokemon:     *p,
	}
}
