CREATE TABLE IF NOT EXISTS favorite_pokemons (
                                                 id SERIAL PRIMARY KEY,
                                                 name VARCHAR,
                                                 count_update INTEGER,
                                                 pokemon_id BIGINT REFERENCES pokemons(id)
    );

CREATE TABLE IF NOT EXISTS pokemons (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR,
                                        type VARCHAR,
                                        move VARCHAR,
                                        weight INTEGER,
                                        height INTEGER,
                                        image_url VARCHAR
);
