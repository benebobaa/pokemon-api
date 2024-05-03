package request

import (
	"mime/multipart"
	"pokemon-api/domain/entity"
)

type PokemonRequest struct {
	Name    string `json:"name" validate:"required"`
	Type    string `json:"type" validate:"required"`
	Move    string `json:"move" validate:"required"`
	Weight  int    `json:"weight" validate:"required"`
	Height  int    `json:"height" validate:"required"`
	Image   *multipart.FileHeader
	BaseUrl string
}

func (p *PokemonRequest) ToEntity(imgUrl string) *entity.Pokemon {
	return &entity.Pokemon{
		Name:     p.Name,
		Type:     p.Type,
		Move:     p.Move,
		Weight:   p.Weight,
		Height:   p.Height,
		ImageUrl: imgUrl,
	}
}
