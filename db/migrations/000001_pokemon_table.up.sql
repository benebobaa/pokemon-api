CREATE TABLE pokemons (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255),
                          type VARCHAR(255),
                          move VARCHAR(255),
                          weight INT,
                          height INT,
                          image_url VARCHAR(255)
);
