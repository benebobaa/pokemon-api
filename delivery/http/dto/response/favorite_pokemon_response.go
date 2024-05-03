package response

type FavoritePokemonResponse struct {
	ID          int64           `json:"id"`
	NickName    string          `json:"nick_name"`
	Pokemon     PokemonResponse `json:"pokemon"`
	CountUpdate int             `json:"count_update"`
}
