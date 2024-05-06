
## API Request

### Currently Published Example Endpoints
### Base URL: `https://pokemon.beneboba.me`

## Endpoints
| Endpoint                    | HTTP Method |           Description           |
|-----------------------------|:-----------:|:-------------------------------:|
| `/api/pokemon`              |    `GET`    |       `Get All Pokemons `       |
| `/api/pokemon/favorite`     |    `GET`    |   `Get All Favorite Pokemons`   |
| `/api/pokemon/favorite`     |   `POST`    |    `Add to Favorite Pokemon`    |
| `/api/pokemon/favorite/{id}` |    `PUT`    | `Update Your Pokemon Nickname`  |
| `/api/pokemon/favorite/{id}` |  `DELETE`   | `Release Your Favorite Pokemon` |

### Example Request
```bash
curl -X GET https://pokemon.beneboba.me/api/pokemon
```

### Example Request POST
```bash
curl -X POST https://pokemon.beneboba.me/api/pokemon/favorite \
-H "Content-Type: application/json" \
-d '{
    "pokemon_id": 3,
    "nick_name":"budi"
}'
```

### For more detail api documentation please import `postman_docs.json` in this repo to your Postman
