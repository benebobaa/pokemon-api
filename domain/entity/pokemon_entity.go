package entity

import "pokemon-api/delivery/http/dto/response"

type Pokemon struct {
	ID              int64            `gorm:"column:id;primaryKey;"`
	Name            string           `gorm:"column:name"`
	Type            string           `gorm:"column:type"`
	Move            string           `gorm:"column:move"`
	Weight          int              `gorm:"column:weight"`
	Height          int              `gorm:"column:height"`
	ImageUrl        string           `gorm:"column:image_url"`
	FavoritePokemon *FavoritePokemon `gorm:"foreignKey:PokemonID"`
}

func (p *Pokemon) TableName() string {
	return "pokemons"
}

func (p *Pokemon) ToResponse() *response.PokemonResponse {
	return &response.PokemonResponse{
		ID:       p.ID,
		Name:     p.Name,
		Type:     p.Type,
		Move:     p.Move,
		Weight:   p.Weight,
		Height:   p.Height,
		ImageUrl: p.ImageUrl,
	}
}
