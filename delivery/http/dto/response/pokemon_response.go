package response

type PokemonResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Move     string `json:"move"`
	Weight   int    `json:"weight"`
	Height   int    `json:"height"`
	ImageUrl string `json:"image_url"`
}
